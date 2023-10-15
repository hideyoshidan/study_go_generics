package slicegenerics

// ConvertSlices []Tを[]Uへ変換
func ConvertSlices[T, U any](srcList []T, convertFunc func(T) U) []U {
	var result []U
	for _, v := range srcList {
		result = append(result, convertFunc(v))
	}
	return result
}
