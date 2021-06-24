package tool

import (
	"fmt"
	"io/ioutil"
)

func GetFile(name string)(string){
	f, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}
