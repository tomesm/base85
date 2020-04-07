package main

import (
	"fmt"
	"github.com/tomesm/base85/pkg"
)

// Example of usage base85
func main() {
	base := base85.NewCoder()
	fmt.Println(base.Encode([]rune("Lorem ipsum dolor sit amet, consectetur adipiscing elit")))
	fmt.Println(base.Decode([]rune("6uQsS@j#Z#@j#?*Ble-0A0>f2@qboC")))
}
