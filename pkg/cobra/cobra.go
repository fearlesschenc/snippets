package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
)

func Main() {
	var str string
	var fooflag string

	cmd := cobra.Command{
		Use:  "test",
		Long: "test long",
		Run: func(cmd *cobra.Command, args []string) {
			println(str)
			println(fooflag)
			fmt.Println("hello, world")
		},
	}

	bar := pflag.NewFlagSet("bar", pflag.ExitOnError)
	bar.StringVar(&str, "str", "bar", "bar usage")
	foo := pflag.NewFlagSet("foo", pflag.ExitOnError)
	foo.StringVar(&str, "str", "foo", "foo usage")
	foo.StringVar(&fooflag, "foo", "fooflag", "fooflag usage")

	//bar.StringVar(&str, "str", "foobar", "foobar usage")

	flags := cmd.Flags()
	flags.AddFlagSet(bar)
	flags.AddFlagSet(foo)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
