package main

import (
	"flag"
	"fmt"

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
	for _, r := range results {
		fmt.Println(r)
	}
}
