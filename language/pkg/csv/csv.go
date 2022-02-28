package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"gopkg.in/yaml.v2"
)

type byNameAndID []RoleRef

func (s byNameAndID) Len() int {
	return len(s)
}

var roleValues = map[string]int{
	"platform": 1,
	"tenant":   4,
	"project":  7,
	"user":     9,
}

func (s byNameAndID) Less(i, j int) bool {
	if s[i].Name != nil && s[j].Name == nil {
		return true
	}

	if s[i].Name == nil && s[j].Name != nil {
		return false
	}

	if s[i].Name != nil && s[j].Name != nil {
		vi := roleValues[*s[i].Name]
		vj := roleValues[*s[j].Name]
		return vi < vj
	}

	ii, _ := strconv.Atoi(*s[i].ID)
	jj, _ := strconv.Atoi(*s[j].ID)
	return ii < jj
}

func (s byNameAndID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type PermissionTemplateRecord struct {
	Module            string
	ModuleDescription string

	Type            string
	TypeDescription string

	Operation            string
	OperationDescription string

	Scope string

	Role            string
	RoleDescription string
}

type RoleRef struct {
	// +optional
	Name *string `yaml:"name,omitempty"`

	// +optional
	ID *string `yaml:"id,omitempty"`
}

type Operation struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Group struct {
	// Name specify the group of these operations, Name will be displayed
	// in UI when manage user's permissions
	Name string `yaml:"name"`

	Operations []Operation `yaml:"operations"`
}

// RoleBindingSpec defines the desired state of RoleBinding
type RoleBindingSpec struct {
	// +required
	RoleRef []RoleRef `yaml:"roleRef"`

	// +required
	ResourceType string `yaml:"resourceType"`

	// +optional
	Scope string `yaml:"scope,omitempty"`

	// +required
	Groups []Group `yaml:"groups"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type RoleBinding struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`

	Spec RoleBindingSpec `yaml:"spec,omitempty"`
}

func GenPaasRoleBindingForType(completeRecords []PermissionTemplateRecord, typ string) {
	var records []PermissionTemplateRecord
	for _, record := range completeRecords {
		if record.Type == typ {
			records = append(records, record)
		}
	}

	operations := map[string]mapset.Set{}
	for _, r := range records {
		refs, exist := operations[r.Operation]
		if !exist {
			refs = mapset.NewSet()
		}

		refs.Add(r.Role)
		operations[r.Operation] = refs
	}

	var existRoleBindings []*RoleBinding
	for operation, roleRefs := range operations {
		var firstOperationRecord PermissionTemplateRecord
		for _, r := range records {
			if r.Operation == operation {
				firstOperationRecord = r
				break
			}
		}

		found := false
		for _, rb := range existRoleBindings {
			rbRoleRefs := mapset.NewSet()
			for _, ref := range rb.Spec.RoleRef {
				if ref.Name != nil {
					id := ""
					switch *ref.Name {
					case "platform":
						id = "1"
					case "tenant":
						id = "4"
					case "project":
						id = "7"
					case "user":
						id = "9"
					}
					rbRoleRefs.Add(id)

					continue
				}

				if ref.ID != nil {
					rbRoleRefs.Add(*ref.ID)
				}
			}

			if roleRefs.Equal(rbRoleRefs) {
				rb.Spec.Groups[0].Operations = append(rb.Spec.Groups[0].Operations, Operation{
					Name:        firstOperationRecord.Operation,
					Description: firstOperationRecord.OperationDescription,
				})

				found = true
				break
			}
		}

		if !found {
			rb := &RoleBinding{
				ApiVersion: "paas.netease.com/v1alpha1",
				Kind:       "RoleBinding",
				Metadata: Metadata{
					Name: fmt.Sprintf("%s-%d", strings.ToLower(typ), len(existRoleBindings)),
				},
				Spec: RoleBindingSpec{
					ResourceType: typ,
					Scope:        "project",
					Groups: []Group{
						{
							Name: firstOperationRecord.TypeDescription,
							Operations: []Operation{
								{
									Description: firstOperationRecord.OperationDescription,
									Name:        firstOperationRecord.Operation,
								},
							},
						},
					},
				},
			}

			var targetRoles []RoleRef
			for ref := range roleRefs.Iter() {
				id := ref.(string)
				switch id {
				case "1":
					name := "platform"
					targetRoles = append(targetRoles, RoleRef{Name: &name})
				case "4":
					name := "tenant"
					targetRoles = append(targetRoles, RoleRef{Name: &name})
				case "7":
					name := "project"
					targetRoles = append(targetRoles, RoleRef{Name: &name})
				case "9":
					name := "user"
					targetRoles = append(targetRoles, RoleRef{Name: &name})
				default:
					targetRoles = append(targetRoles, RoleRef{ID: &id})
				}
			}
			sort.Sort(byNameAndID(targetRoles))

			rb.Spec.RoleRef = targetRoles
			existRoleBindings = append(existRoleBindings, rb)
		}
	}

	f, err := os.Create(path.Join("/Users/chenchen/Dev/files/rolebindings", fmt.Sprintf("%s.yaml", typ)))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i, rb := range existRoleBindings {
		output, err := yaml.Marshal(rb)
		if err != nil {
			panic(err)
		}

		f.Write(output)

		if i+1 != len(existRoleBindings) {
			f.Write([]byte("\n"))
			f.Write([]byte("---\n"))
		}
	}
}

func GenPaasRoleBindings() {
	f, err := os.Open("/Users/chenchen/Dev/files/paas.csv.bak")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	records, _ := csv.NewReader(f).ReadAll()
	records = records[1:]

	resourceTypes := mapset.NewSet()
	var completeRecords []PermissionTemplateRecord
	for i, _ := range records {
		fields := records[i]
		if len(fields) < 9 {
			continue
		}

		r := &PermissionTemplateRecord{}
		r.Module = fields[0]
		r.ModuleDescription = fields[1]
		r.Type = fields[2]
		r.TypeDescription = fields[3]
		r.Operation = fields[4]
		r.OperationDescription = fields[5]
		r.Scope = fields[6]
		r.Role = fields[7]
		r.RoleDescription = fields[8]

		resourceTypes.Add(r.Type)
		completeRecords = append(completeRecords, *r)
	}

	for typ := range resourceTypes.Iter() {
		resourceType := typ.(string)
		GenPaasRoleBindingForType(completeRecords, resourceType)
	}
}

func Read() {
	f, err := os.Open("/Users/chenchen/Dev/files/paas.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()
	//records = records[1:]
	for _, r := range records {
		fmt.Println(strings.Join(r, ","))
	}

	writer := csv.NewWriter(os.Stdout)
	err = writer.Write([]string{"PaaS", "PaaS", "zookeeper", "zookeeper", "listCluster11111", "查看Zookeeper列表", "project", "1", "平台全部权限"})
	if err != nil {
		panic(err)
	}
	err = writer.Error()
	if err != nil {
		panic(err)
	}

	writer.Flush()
	err = writer.Error()
	if err != nil {
		panic(err)
	}
}
