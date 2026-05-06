package main

import (
	"FundamentalPart8/practice"
	"fmt"
)

func main() {
	//tangTa := practice.TangMucTieu(50)
	//fmt.Println(tangTa(5))
	//fmt.Println(tangTa(10))

	//buoiTap1 := practice.TaoBuoiTap("sontay")
	//ans := buoiTap1(1, 2, 3, 4)
	//fmt.Println(ans)

	//flashSale15 := practice.CreateDiscount(15)
	//clearance50 := practice.CreateDiscount(50)
	//
	//fmt.Println(flashSale15(200))
	//fmt.Println(clearance50(200))
	//price := []float64{150, 100, 50, 20, 200}
	//filterPlant := practice.FilterPlants(price, func(price float64) bool {
	//	return price > 100
	//})
	//fmt.Println(filterPlant)
	//voucher20 := practice.CreateLimitedVoucher(20, 2)
	//voucher20(200)
	//voucher20(100)
	//voucher20(100)
	p1 := 50.0
	p2 := 80.0
	combo := []float64{100.0, 150.0, 200.0}
	combo = append(combo, p1, p2)
	ans := practice.SumCart(combo...)
	fmt.Println(ans)
}
