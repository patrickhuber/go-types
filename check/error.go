package check

func Error(err error) {
	if err == nil {
		return
	}
	panic(err)
}
