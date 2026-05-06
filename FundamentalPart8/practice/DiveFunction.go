package practice

import "fmt"

func ChucMung(ten string) {
	fmt.Println("Chuc mung:", ten)
}

func ThucHien(f func(string), name string) {
	f(name)
}

func TangMucTieu(mucTieuHienTai int) func(int) int {
	return func(i int) int {
		mucTieuHienTai = mucTieuHienTai + i
		return mucTieuHienTai
	}
}

func TaoBuoiTap(tenHocVien string) func(...int) int {
	tongKhoiLuong := 0
	return func(hiepTap ...int) int {
		defer fmt.Printf("Hiep tap cua [%s] \n", tenHocVien)
		for _, ta := range hiepTap {
			tongKhoiLuong = tongKhoiLuong + ta
		}
		return tongKhoiLuong
	}
}

func CreateDiscount(percent float64) func(float65 float64) float64 {
	return func(price float64) float64 {
		return price - (price * percent / 100)
	}
}

// recursion
type Category struct {
	Name     string
	Children []Category
}

func PrintCategory(category Category, level int) {
	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}
	fmt.Println(category.Name)
	for _, child := range category.Children {
		PrintCategory(child, level+1)
	}
}

// homework
func FilterPlants(prices []float64, condition func(float64) bool) []float64 {
	ans := []float64{}
	for _, price := range prices {
		if condition(price) {
			ans = append(ans, price)
		}
	}
	return ans
}

// mutation closure
func CreateLimitedVoucher(percent float64, maxUses int) func(float64) float64 {
	return func(price float64) float64 {
		if maxUses <= 0 {
			fmt.Println("Voucher run out of use")
			return price
		}
		maxUses--
		fmt.Printf("Voucher has use, [%d] left \n", maxUses)
		return price - (price * percent / 100)
	}
}

// variadic & spliting
func SumCart(prices ...float64) float64 {
	ans := 0.0
	for _, price := range prices {
		ans = ans + price
	}
	return ans
}

// recursion
func CountCategories(c Category) int {
	ans := 1
	for _, child := range c.Children {
		ans += CountCategories(child)
	}
	return ans
}
