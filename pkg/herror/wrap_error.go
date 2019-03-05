package herror

// Wraps an error in a HorizonError. When the error is itself a HorizonError,
// the error is returned without being wrapped and without applying the
// options.
func Wrap(err error, options ...Option) HorizonError {
	var herr *HorizonError

	if err == nil {
		herr = New("Invalid Parameter", options...)
	} else if castErr, ok := err.(HorizonError); ok {
		herr = &castErr
	} else if castErr, ok := err.(*HorizonError); ok {
		herr = castErr
	} else {
		herr = New("Unexpected Error", CausedBy(err), Detail(err.Error()))
		for _, opt := range options {
			herr = opt(herr)
		}
	}

	return *herr
}
