package internal

func Trace(err error, dopanic bool) {

	Logger(err)

	Errors(err, dopanic)

}
