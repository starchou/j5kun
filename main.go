package main

import (
	"flag"
	"fmt"

	j4kun "github.com/starchou/j5kun/pkg/id"
)

var ID string

func init() {
	flag.StringVar(&ID, "ID", "", "please input your ID number!")
	flag.Parse()

}
func main() {
	if ID != "" {
		fmt.Println(j4kun.NewIDCard(ID).Analysis())
	} else {
		fmt.Println("please input your ID number!")
	}
}
