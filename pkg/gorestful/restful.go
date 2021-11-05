package gorestful

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"log"
	"net/http"
)

type User struct {
	ID   string `json:"id" description:"id of user"`
	Name string `json:"name" description:"name of user"`
	Age  int    `json:"age" description:"age of user"`
}

type UserResource struct {
	users map[string]User
}

func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users/v1").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"users"}
	ws.Route(ws.GET("/").To(u.findAllUsers).
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]User{}).
		Returns(200, "OK", []User{}))

	ws.Route(ws.GET("/{id}").To(u.findUser))

	return ws
}

func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	for _, v := range u.users {
		if v.ID == id {
			response.WriteEntity(v)
		}
	}
}

func (u UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	var list []User
	for _, each := range u.users {
		list = append(list, each)
	}
	response.WriteEntity(list)
}

func Main() {
	u := UserResource{map[string]User{
		"cc": {
			ID:   "1",
			Name: "chenc",
			Age:  30,
		},
	}}

	restful.DefaultContainer.Add(u.WebService())
	//restful.Filter()

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(),
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: func(s *spec.Swagger) {
			s.Info = &spec.Info{
				InfoProps: spec.InfoProps{
					Title:       "UserService",
					Description: "Resource for managing Users",
					Contact: &spec.ContactInfo{
						Name:  "john",
						Email: "john@doe.rp",
						URL:   "http://johndoe.org",
					},
					License: &spec.License{
						Name: "MIT",
						URL:  "http://mit.org",
					},
					Version: "1.0.0",
				},
			}
			s.Tags = []spec.Tag{
				spec.Tag{
					TagProps: spec.TagProps{
						Name:        "users",
						Description: "Managing users",
					},
				},
			}
		},
	}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/chenchen/Dev/tool/swagger-ui/dist"))))

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
