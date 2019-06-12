package main

import (
	"fmt"

	"github.com/jgoodall/pb-test-one/v2/api"
)

func main() {
	thing := api.TestOneTwo{TestOne: "thing one", TestTwo: false}
	fmt.Printf("%+v\n", thing)
}
