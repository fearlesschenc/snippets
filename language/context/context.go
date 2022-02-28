package context

import (
	"context"
	"fmt"
)

type key struct{}

func insertContextValue(ctx context.Context, i int) {
	vals := ctx.Value(key{}).(map[string]int)
	vals[fmt.Sprintf("%d", i)] = i
}

func Propagate() {
	vals := make(map[string]int)
	ctx := context.WithValue(context.Background(), key{}, vals)
	for i := 1; i < 10; i++ {
		insertContextValue(ctx, i)
	}

	for k, v := range vals {
		fmt.Println(k, v)
	}
}
