package herror

import (
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
)

// A function that is able to apply transformations to a HorizonError.
type Option func(*HorizonError) *HorizonError

// Provide the original error that caused this problem.
func CausedBy(origError error) Option {
	return func(err *HorizonError) *HorizonError {
		if origError == nil {
			stacktrace := string(debug.Stack())
			log.Printf("WARN herror.CausedBy() called with nil parameter; ignoring. stacktrace: %s", stacktrace)
			return err
		}

		msg := origError.Error()

		err.Source.CausedBy = origError
		err.Source.Message = &msg

		return err
	}
}

// Provide an application-specific error code, expressed as a string value.
func Code(code string) Option {
	return func(err *HorizonError) *HorizonError {
		err.Code = &code
		return err
	}
}

// Provide a human-readable explanation specific to this occurrence of the
// problem. You may use standard fmt value substitutions in msgf.
func Detail(msgf string, args ...interface{}) Option {
	return func(err *HorizonError) *HorizonError {
		detail := fmt.Sprintf(msgf, args...)
		err.Detail = &detail
		return err
	}
}

// Provide an event ID to correlate the error to other platform activity.
func EventID(id string) Option {
	return func(err *HorizonError) *HorizonError {
		err.Source.EventID = &id
		return err
	}
}

// Provide a request ID to correlate the error to a user request.
func RequestID(id string) Option {
	return func(err *HorizonError) *HorizonError {
		err.Source.RequestID = &id
		return err
	}
}

// Provide the HTTP status code applicable to this problem.
func Status(httpStatus int) Option {
	status := strconv.Itoa(httpStatus)
	return func(err *HorizonError) *HorizonError {
		err.Status = &status
		return err
	}
}

// Override the default title used when wrapping errors.
func Title(title string) Option {
	return func(err *HorizonError) *HorizonError {
		err.Title = title
		return err
	}
}
