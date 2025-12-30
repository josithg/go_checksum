package main

import (
	"fmt"

	"github.com/codingsince1985/checksum"
	"github.com/mbndr/figlet4go"
)

func main() {
	var file string
	var opt int8
	var oh string
	ascii := figlet4go.NewAsciiRender()

	renderStr, _ := ascii.Render("josithg/checksum")
	fmt.Print(renderStr)

	fmt.Println()
	fmt.Print("enter your file path: ")
	fmt.Scanln(&file)

	fmt.Println("	Press 1 for SHA256		Press 2 for SHA1 	 Press 3 for MD5")
	fmt.Print("enter your choice: ")
	fmt.Scanln(&opt)

	switch opt {
	case 1:
		s256, _ := checksum.SHA256sum(file)
		fmt.Println("the SHA256 sum is: ", s256)
		fmt.Print("enter your hash to compare : ")
		fmt.Scanln(&oh)
		if oh == s256 {
			fmt.Println()
			fmt.Println("Hash matched ")
		} else {
			fmt.Println()
			fmt.Println("hash does not match")
		}
	case 2:
		s1, _ := checksum.SHA1sum(file)
		fmt.Println("the SHA1 sum is: ", s1)
		fmt.Print("enter your hash to compare : ")
		fmt.Scanln(&oh)
		if oh == s1 {
			fmt.Println()
			fmt.Println("Hash matched ")
		} else {
			fmt.Println()
			fmt.Println("hash does not match")
		}
	case 3:
		m5, _ := checksum.MD5sum(file)
		fmt.Println(" the MD5 sum is: ", m5)
		fmt.Print("enter your hash to compare : ")
		fmt.Scanln(&oh)
		if oh == m5 {
			fmt.Println("hash matched ")
			fmt.Println()
		} else {
			fmt.Println()
			fmt.Println("hash does not match")
		}
	default:
		fmt.Println("Rerun the program with correct option")

	}
	fmt.Println("visit github.com/josithg for source code")

}
