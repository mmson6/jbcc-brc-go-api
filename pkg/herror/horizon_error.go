package herror

import (
	"encoding/json"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

////////////////////////////////////////////////////////////////////////////////
// INIT

var includeStackTrace bool

func init() {
	valStr := os.Getenv("DEBUG")
	valBool, err := strconv.ParseBool(valStr)
	if err == nil {
		includeStackTrace = valBool
		if includeStackTrace {
			log.Printf(`including stack traces in errors (DEBUG=%s)`, valStr)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// STRUCT

type HorizonError struct {
	// Required
	ID     string `json:"id"`
	Source Source `json:"source"`
	Title  string `json:"title"`

	// Optional
	Code   *string `json:"code,omitempty"`
	Detail *string `json:"detail,omitempty"`
	Status *string `json:"status,omitempty"`
}

type Source struct {
	// Required
	Timestamp time.Time `json:"timestamp"`

	// Optional
	CausedBy   error    `json:"causedBy,omitempty"`
	EventID    *string  `json:"eventId,omitempty"`
	Message    *string  `json:"message,omitempty"`
	RequestID  *string  `json:"requestId,omitempty"`
	StackTrace []string `json:"stackTrace,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////
// INITIALIZERS

// Creates a HorizonError with the specified title. The title should be a
// short, human-readable summary of the problem that will not change from
// occurrence to occurrence of the problem.
func New(title string, options ...Option) *HorizonError {
	id := uuid.NewV4().String()
	// id := "96d7e6e1-e08e-4c08-8775-9d4d3d60a8b7"
	now := time.Now()

	var stackTrace []string
	if includeStackTrace {
		str := string(debug.Stack())
		stackTrace = strings.Split(str, "\n")
	}

	err := &HorizonError{
		ID: id,
		Source: Source{
			StackTrace: stackTrace,
			Timestamp:  now,
		},
		Title: title,
	}

	for _, opt := range options {
		err = opt(err)
	}

	return err
}

////////////////////////////////////////////////////////////////////////////////
// METHODS

func (he HorizonError) HttpStatusCode() (int, bool) {
	if he.Status == nil {
		return 500, false
	}

	status64, err := strconv.ParseInt(*he.Status, 10, 31)
	if err != nil {
		return 500, false
	}

	statusCode := int(status64)
	return statusCode, true
}

//////////
// error

func (he HorizonError) Error() string {
	str := he.Title
	if he.Detail != nil {
		str += ": " + *he.Detail
	} else if he.Source.Message != nil {
		str += ": " + *he.Source.Message
	}
	return str
}

//////////
// json.Marshaler

func (s Source) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{}, 6)
	m["timestamp"] = s.Timestamp.UTC().Format(time.RFC3339)

	if s.CausedBy != nil {
		m["causedBy"] = s.CausedBy.Error()
	}
	if s.EventID != nil {
		m["eventId"] = *s.EventID
	}
	if s.Message != nil {
		m["message"] = *s.Message
	}
	if s.RequestID != nil {
		m["requestId"] = *s.RequestID
	}
	if len(s.StackTrace) > 0 {
		m["stackTrace"] = s.StackTrace
	}

	return json.Marshal(m)
}
