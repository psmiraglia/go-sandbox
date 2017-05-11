package main

import (
	"fmt"
	"github.com/psmiraglia/sandbox/common"
)

func main() {
	fmt.Println(common.Version())
	fmt.Println(common.NamedVersion())
}
