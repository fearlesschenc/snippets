package pflag

import (
	"github.com/spf13/pflag"
	"os"
	//ctrl "sigs.k8s.io/controller-runtime"
)

type Flag struct {
	flag int
}

func Main() {
	f := &Flag{flag: 1234}
	fs := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
	fs.IntVar(&f.flag, "flagname", f.flag, "help message for flagname")

	fs.Parse(os.Args[1:])

	println(f.flag)
}
