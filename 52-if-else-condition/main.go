package main

import "fmt"

/*
	- Ngày: 19.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	var (
		age = 30
	)

	if age > 18 {
		fmt.Println("Bạn có tuổi lớn hơn 18")
	} else {
		fmt.Println("Bạn có tuổi nhỏ hơn hoặc bằng 18")
	}
}
