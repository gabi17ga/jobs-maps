package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	content, err := request("https://www.ejobs.ro/company/ejobs-group/266860")
	if err != nil {
		panic(err)
	}
	i := extract(*content)
	spew.Dump(i)
}
