package main

import "fmt"

/*
	- Ngày: 19.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	/*
		Một số dạng câu lệnh có điều kiện
	*/

	var (
		number1 = 25
		number2 = -1
		number3 = 21
		number4 = 10
	)

	if number1%5 == 0 {
		fmt.Printf("%d chia hết cho 5", number1)
	}

	if number2 < 0 {
		fmt.Printf("\n%d là số nhỏ hơn 0", number2)
	}

	if number3 >= 15 && number3 <= 30 {
		fmt.Printf("\n15<= %d <= 30", number3)
	}
	if number4 == 10 && number4%2 == 0 {
		fmt.Printf("\n%d là số tự nhiên bằng 10, và chia hết cho 2", number4)

	}

}
