package main

import (
	"fmt"
	"github.com/tomesm/base85/pkg"
)

func main() {

	base := base85.NewCoder()

	fmt.Println(base.Encode([]rune("Lorem ipsum dolor sit amet, consectetur adipiscing elit")))

	fmt.Println(base.Decode([]rune("6uQsS@j#Z#@j#?*Ble-0A0>f2@qboC")))

	//reader := bufio.NewReader(os.Stdin)

	//fmt.Println("Simple Shell")
	//fmt.Println("---------------------")

	//for {
	//fmt.Print("-> ")
	//text, _ := reader.ReadString('\n')
	//// convert CRLF to LF
	//text = strings.Replace(text, "\n", "", -1)

	//if strings.Compare("hi", text) == 0 {
	//fmt.Println("hello, Yourself")
	//}

	//}

}
