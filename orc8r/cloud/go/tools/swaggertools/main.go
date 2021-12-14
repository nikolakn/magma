package main

import (
	"flag"
	combine_swagger "magma/orc8r/cloud/go/tools/swaggertools/combine_swagger"
	swaggergen "magma/orc8r/cloud/go/tools/swaggertools/swaggergen"

	"github.com/golang/glog"
)

func main() {
	cmdCombine := flag.Bool("combine", false, "calls combine_swagger command")
	cmdGen := flag.Bool("gen", false, "calls swaggergen command")
	flag.Parse()

	if *cmdCombine {
		combine_swagger.Run()
	} else if *cmdGen {
		swaggergen.Run()
	} else {
		glog.Fatal("command combine or gen must be specified")
	}
}
