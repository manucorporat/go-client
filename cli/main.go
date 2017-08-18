package main

import (
	"flag"
	"fmt"

	"github.com/bblfsh/sdk/uast"

	"github.com/bblfsh/go-client"
)

var endpoint = flag.String("e", "localhost:9432", "endpoint of the babelfish server")
var filename = flag.String("f", "", "file to parse")

func main() {
	flag.Parse()

	client, err := bblfsh.NewBblfshClient(*endpoint)
	if err != nil {
		panic(err)
	}
	res, err := client.Parse().ReadFile(*filename).Do()
	if err != nil {
		panic(err)
	}
	results, _ := bblfsh.Find(res.UAST, "//NumLiteral")
	for i, r := range results {
		fmt.Println("- ", i+1, " ----------------------")
		node := r.(uast.Node)
		fmt.Println(node.String())
	}
}
