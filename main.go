package main

import (
	"bufio"
	"fmt"
	"github.com/tomesm/base85/pkg"
	"os"
	"strings"
)

func main() {

	base := base85.New()

	//fmt.Println(base.Encode("Lorem ipsum dolor sit amet, consectetur adipiscing elit"))

	//	fmt.Println(base.Decode("6uQsS@j#Z#@j#?*Ble-0A0>f2@qboC"))
	//

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}

}
