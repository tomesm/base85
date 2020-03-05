package main

import (
	"fmt"
	"github.com/tomesm/base85/pkg"
)

func main() {

	base := base85.New()

	fmt.Println(base.Encode("Donec nec euismod orci"))

}
