package exception

func PanicLog(err interface{}) {
	if err != nil {
		panic(err)
	}
}
