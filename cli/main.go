package main

import (
	"flag"
	"fmt"

	"github.com/bblfsh/go-client"
	"github.com/bblfsh/sdk/uast"
)

var endpoint = flag.String("e", "localhost:9432", "endpoint of the babelfish server")
var filename = flag.String("f", "", "file to parse")
var query = flag.String("q", "", "xpath expression")

func main() {
	flag.Parse()
	if *filename == "" {
		fmt.Println("filename was not provided. Use the -f flag\n")
		return
	}

	client, err := bblfsh.NewBblfshClient(*endpoint)
	if err != nil {
		panic(err)
	}
	res, err := client.Parse().ReadFile(*filename).Do()
	if err != nil {
		panic(err)
	}
	if *query == "" {
		fmt.Println(res.UAST)
	} else {
		results, _ := bblfsh.Find(res.UAST, *query)
		for i, r := range results {
			fmt.Println("-", i+1, "----------------------")
			node := r.(uast.Node)
			fmt.Println(node.String())
		}
	}
}
