package regexp

import (
	"fmt"
	"regexp"
)

func Main() {
	//clusterPath := regexp.MustCompile(`{([^/.\s]+)}`)
	////fmt.Println(clusterPath.ReplaceAllString("/apis/nzk.netease.com/v1alpha2/cluster/1", ""))
	//fmt.Printf("%q", clusterPath.FindAllStringSubmatch("/foo/{foo}/bar/{bar}", -1))

	re := regexp.MustCompile(`^(\S+): (.*) (https://\S+)$`)
	matches := re.FindStringSubmatch("Subtask-48021: 电科院 memcached-operator 调协逻辑开发 https://overmind-project.netease.com/v2/my_workbench/taskdetail/Task-48021")
	if len(matches) != 0 {
		fmt.Printf("%q\n", matches)
	}

	commitMessageTitlePattern := regexp.MustCompile(`^(!?)(\w+)(?:\((\w+)\))?: (.+)$`)
	matches = commitMessageTitlePattern.FindStringSubmatch(`!feature(style): 将代码扁平化，放在根目录中`)
	//matches = commitMessageTitlePattern.FindStringSubmatch(`!feature: 将代码扁平化，放在根目录中`)
	for _, match := range matches {
		fmt.Println(match)
	}
}
