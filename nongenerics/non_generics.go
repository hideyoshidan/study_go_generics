package nongenerics

//---------------------------------------
// Genericsを使わないケース
// int, float型それぞれで同じ関数を用意する必要がある
//---------------------------------------

func SumInt(nums []int) int {
	var total int
	for _, v := range nums {
		total += v
	}
	return total
}

func SumFloat32(nums []float32) float32 {
	var total float32
	for _, v := range nums {
		total += v
	}
	return total
}

func SumFloat64(nums []float64) float64 {
	var total float64
	for _, v := range nums {
		total += v
	}
	return total
}
