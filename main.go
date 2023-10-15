package main

import (
	"fmt"
	"strconv"

	"generics.com/generics"
	"generics.com/nongenerics"
	"generics.com/slicegenerics"
	"generics.com/structgenerics"
)

func main() {
	// notUseGenerics()
	// useseGenerics()
	// anyfunc()
	// interfaceFunc()
	// structFunc()
	// greet()
	//
	sliceFunc()
}

// 3つの関数をそれぞれの型に合わせて呼ぶ必要がある
func notUseGenerics() {
	fmt.Printf("==============%s==============\n", "Not using generics")
	intVal := []int{1, 2, 3, 4, 5}
	float32Val := []float32{1.1, 1.2, 1.3, 1.4, 1.5}
	float64Val := []float64{1.1, 1.2, 1.3, 1.4, 1.5}

	// int
	fmt.Printf("%s\n", fmt.Sprintf("%s Total : %d", "Int", nongenerics.SumInt(intVal)))
	// float32
	fmt.Printf("%s\n", fmt.Sprintf("%s Total : %v", "Float32", nongenerics.SumFloat32(float32Val)))
	// float64
	fmt.Printf("%s\n", fmt.Sprintf("%s Total : %v", "Float64", nongenerics.SumFloat64(float64Val)))
}

// 1つの関数でok
func useseGenerics() {
	fmt.Printf("==============%s==============\n", "Using generics")
	intVal := []int{1, 2, 3, 4, 5}
	float32Val := []float32{1.1, 1.2, 1.3, 1.4, 1.5}
	float64Val := []float64{1.1, 1.2, 1.3, 1.4, 1.5}

	// int
	fmt.Printf("%s\n", fmt.Sprintf("%s Total : %d", "Int", generics.Sum(intVal)))
	// float32
	fmt.Printf("%s\n", fmt.Sprintf("%s Total : %v", "Float32", generics.Sum(float32Val)))
	// float64
	fmt.Printf("%s\n", fmt.Sprintf("%s Total : %v", "Float64", generics.Sum(float64Val)))
}

// any型
func anyfunc() {
	type sample struct {
		txt  string
		code string
	}
	sampleSlice := []sample{
		{txt: "RED", code: "X0001"},
		{txt: "BLUE", code: "X0002"},
		{txt: "GREEN", code: "X0003"},
		{txt: "ORANGE", code: "X0004"},
		{txt: "PINK", code: "X0005"},
	}
	intSlice := []int{1, 2, 3, 4, 5}
	float64Slice := []float64{1.1, 1.2, 1.3, 1.4, 1.5}
	stringSclice := []string{"A", "B", "C", "D", "E"}

	// Result
	fmt.Printf("[%s]\n", generics.Any(intSlice))
	fmt.Printf("[%s]\n", generics.Any(float64Slice))
	fmt.Printf("[%s]\n", generics.Any(stringSclice))
	fmt.Printf("[%s]\n", generics.Any(sampleSlice))

}

func interfaceFunc() {
	strval := "1111111111"
	intval := 1111111111

	fmt.Printf("%#v\n", structgenerics.SampleFunc(strval))
	fmt.Printf("%#v\n", structgenerics.SampleFunc(intval))

	// NGになる
	// boolval := true
	// fmt.Printf("%#v\n", structgenerics.SampleFunc(boolval))
}

func structFunc() {
	a := &structgenerics.SampleStructA{Name: "Yunjin"}
	b := &structgenerics.SampleStructB{Name: "Eunchae"}

	fmt.Printf("%#v\n", structgenerics.SampleStructFunc(a))
	fmt.Printf("%#v\n", structgenerics.SampleStructFunc(b))

	// NGになる
	// c := &structgenerics.SampleStructC{Name: "Sakura"}
	// fmt.Printf("%#v\n", structgenerics.SampleStructFunc(c))
}

func greet() {
	var gtEn structgenerics.GreetingEnglish = "Hello"
	gt := structgenerics.GreetingType[structgenerics.GreetingEnglish]{
		Name: "Yunjin",
		Val:  gtEn,
	}
	fmt.Printf("%s\n", gt.Greet())

	var gtJa structgenerics.GreetingEnglish = "こんにちは"
	gt = structgenerics.GreetingType[structgenerics.GreetingEnglish]{
		Name: "ユンジン",
		Val:  gtJa,
	}
	fmt.Printf("%s\n", gt.Greet())
}

func sliceFunc() {
	ints := []int{1, 2, 3, 4, 5}
	result := slicegenerics.ConvertSlices(ints, func(num int) string {
		return strconv.Itoa(num)
	})

	fmt.Printf("result : %s", result)
}
