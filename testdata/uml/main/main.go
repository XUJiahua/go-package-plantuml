package main

import (
	"fmt"
	uml "github.com/qquunn/go-package-plantuml/testdata/uml"
	sub "github.com/qquunn/go-package-plantuml/testdata/uml/sub"
	sub2 "github.com/qquunn/go-package-plantuml/testdata/uml/sub2"
)

func main() {
	fmt.Println(uml.SA{})
	fmt.Println(sub.SA{})
	fmt.Println(sub2.Sub2A{})
}
