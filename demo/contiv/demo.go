package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/qquunn/go-package-plantuml/codeanalysis"
	"io/ioutil"
)

func main() {

	log.SetLevel(log.InfoLevel)

	config := codeanalysis.Config{
		CodeDir:   "/appdev/gopath/src/github.com/contiv/netplugin",
		GopathDir: "/appdev/gopath",
	}

	config.VendorDir = config.CodeDir + "/vendor"

	config.IgnoreDirs = []string{config.CodeDir + "/vendor"}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile("/tmp/uml.txt")

	bytes, _ := ioutil.ReadFile("/tmp/uml.txt")

	fmt.Println(string(bytes))

	// java -jar /app/plantuml.jar  /tmp/uml.txt -tsvg && open2 /tmp/uml.svg

}
