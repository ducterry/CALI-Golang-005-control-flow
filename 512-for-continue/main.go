package main

import "fmt"

/*
	- Ngày: 19.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}
}
