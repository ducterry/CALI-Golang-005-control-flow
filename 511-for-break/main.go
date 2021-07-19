package main

import "fmt"

/*
	- Ngày: 19.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("Số chia hết cho cả 3 và 5 là  %d\n", num)
			break
		}
	}
}
