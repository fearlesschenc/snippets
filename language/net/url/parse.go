package url

import (
	"fmt"
	"net/url"
)

func Parse() {
	uri, err := url.ParseRequestURI("http://127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println(uri.Host)
}
