package common

func Must[T interface{}](output T, err error) T {
	if err != nil {
		panic(err)
	}
	return output
}
