package generics

type addable interface {
	int | float64 | string
}

func sum[T addable](v1 T, v2 T) T {
	return v1 + v2
}

func main() {
	fmt.Println("int -- 1 + 2: ", min(1, 2))  // 3
	fmt.Println("float64 -- 1.0 + 2.0: " sum(1.0, 2.0)) // 3.0
	fmt.Println("string -- '1' + '2': ", sum("1", "2")) // "12"
}