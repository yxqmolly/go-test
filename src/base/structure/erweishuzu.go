package main

import (
	// "crypto/rand"
	// "fmt"
	// "math/big"
)
// /** 
//  * @Description: 二维数组
//  * @param arr
//  */
// func main() {
// 	var arr [3][5]float64
// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr[i]); j++ {
// 			max := big.NewInt(101)
// 			num, err := rand.Int(rand.Reader, max)
// 			if err != nil {
// 				fmt.Println("生成随机数失败：", err)
// 			} else {
// 				t := float64(num.Int64())
// 				fmt.Println("生成的随机数是：", t)
// 				arr[i][j] = t
// 			}
// 		}
// 	}
// 	for i := 0; i < len(arr); i++ {
// 		var total float64
// 		for j := range arr[i] {
// 			total += arr[i][j]
// 		}
// 		fmt.Printf("第%d班的平均成绩是：%.2f\n", i+1, total/float64(len(arr[i])))
// 	}
// }
