package structgenerics

import "fmt"

// InterfaceでTEST
type SampleInterface interface {
	string | int
}

func SampleFunc[T SampleInterface](arg T) T {
	return arg
}

// ----------------------------------------------------------------------------

// StructでTEST
type SampleStructA struct {
	Name string
}

type SampleStructB struct {
	Name string
}

type SampleStructC struct {
	Name string
}

func SampleStructFunc[ST *SampleStructA | *SampleStructB](param ST) string {
	return fmt.Sprintf("%v", param)
}

// ----------------------------------------------------------------------------
type GreetingJapanese string
type GreetingEnglish string

// メソッドには、型パラメーターは持てない。
// type AllowGenerics struct {
// 	Name string
// }
//
// func NewAllowGenerics(name string) *AllowGenerics {
// 	return &AllowGenerics{
// 		Name: name,
// 	}
// }
//
// method must have no type parametersというエラーが出る。
// func (a *AllowGenerics)Greeting[GS GreetingJapanese | GreetingEnglish](greeting GS) string {
// 	return fmt.Sprintf("%v", greeting)
// }

// Genericsの構造体は、インスタンス化することができない
type GreetingType[G GreetingJapanese | GreetingEnglish] struct {
	Name string
	Val  G
}

func (g *GreetingType[G]) Greet() string {
	return fmt.Sprintf("%s %s !!!!!", g.Val, g.Name)
}

// ----------------------------------------------------------------------------
type LastTest struct {
	// 以下のようには使えない。
	// Geet GreetingType
	GreetEn GreetingType[GreetingEnglish]
	GreetJa GreetingType[GreetingEnglish]
}
