package main

import (
	"fmt"
	"github.com/joeyghb/go-training/example02-golang-package/controller"
)

func main() {
	fmt.Println("一天就學會 Go 語言")

	hi := controller.HelloWorld("pgluffy")
	fmt.Println(hi)
}
