// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// func main() {
// 	speed := 100 // #1
// 	// speed := 10 // #2

// 	fast := speed >= 80
// 	slow := speed < 20

// 	fmt.Printf("fast's type is %T\n", fast)

// 	fmt.Printf("going fast? %t\n", fast)
// 	fmt.Printf("going slow? %t\n", slow)

//		fmt.Printf("is it 100 mph? %t\n", speed == 10)
//		fmt.Printf("is it not 100 mph? %t\n", speed != 100)
//	}
func main() {
	var speed int
	fmt.Print("请输入速度: ")
	fmt.Scanln(&speed)

	fast := speed >= 80
	slow := speed < 20

	if fast {
		fmt.Println("fast")
	} else if slow {
		fmt.Println("slow")
	} else {
		fmt.Println("normal")
	}
}
