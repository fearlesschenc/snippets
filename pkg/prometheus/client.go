package prometheus

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"time"
)

func ClientMain() {
	client, err := api.NewClient(api.Config{
		Address: "http://127.0.0.1:9090",
	})
	if err != nil {
		panic(err)
	}

	promAPI := v1.NewAPI(client)
	value, _, err := promAPI.Query(context.TODO(), "up", time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(value.Type())
	fmt.Println(value.String())
	fmt.Printf("%+v\n", value)
}
