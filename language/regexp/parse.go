package regexp

import (
	"fmt"
	"regexp"
)

func Parse() {
	clusterPath := regexp.MustCompile(`{([^/.\s]+)}`)
	//fmt.Println(clusterPath.ReplaceAllString("/apis/nzk.netease.com/v1alpha2/cluster/1", ""))
	fmt.Printf("%q", clusterPath.FindAllStringSubmatch("/foo/{foo}/bar/{bar}", -1))

}
