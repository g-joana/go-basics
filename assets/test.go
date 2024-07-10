package assets

import "fmt"
// fmt -> format (print funcs)

// func starting with uppercase is a public function (acessible on other files by importing this package)
func Test() {
	fmt.Println("hello!")
}
