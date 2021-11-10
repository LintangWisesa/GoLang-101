package main

import (
	"fmt"
	pkg "GoLang-101/myPkg"
)

func main(){
	fmt.Println(pkg.Data)
	fmt.Println(pkg.Data.Nama)
	fmt.Println(pkg.Data.Usia)
	fmt.Println(pkg.Call())
}