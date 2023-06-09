package generics

type dataType1 interface {
	int | float32 | float64 | int32 | int64
}

type dataType2 interface {
	string | interface{}
}

func GenericsFunc[X dataType1 | dataType2](param X) X {
	return param
}
