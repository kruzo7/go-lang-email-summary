package internal

func Errors(err error, dopanic bool) {

	if !dopanic {
		return
	}

	if err != nil {
		panic(err)
	}
}
