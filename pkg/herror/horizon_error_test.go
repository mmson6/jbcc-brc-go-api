package herror

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

////////////////////////////////////////////////////////////////////////////////
// New()

func TestUnit_New_PopulatesID(t *testing.T) {
	// SETUP

	// ACTIONS

	result := New("Test Title")

	// ASSERTIONS

	assert.NotEmpty(t, result.ID)
}

func TestUnit_New_PopulatesSourceTimestamp(t *testing.T) {
	// SETUP

	start := time.Now()

	// ACTIONS

	result := New("Test Title")

	// ASSERTIONS

	ts := result.Source.Timestamp
	assert.True(t, ts.After(start))
}

func TestUnit_New_PopulatesTitle(t *testing.T) {
	// SETUP

	// ACTIONS

	result := New("Test Title")

	// ASSERTIONS

	assert.Equal(t, "Test Title", result.Title)
}

////////////////////////////////////////////////////////////////////////////////
// Error()

func TestUnit_Error_WhenOnlyTitleSet(t *testing.T) {
	// SETUP

	herr := New("Test Title")

	// ACTIONS

	msg := herr.Error()

	// ASSERTIONS

	assert.Equal(t, "Test Title", msg)
}

func TestUnit_Error_WhenTitleAndDetailSet(t *testing.T) {
	// SETUP

	detail := "More information about the error"

	herr := New("Test Title")
	herr.Detail = &detail

	// ACTIONS

	msg := herr.Error()

	// ASSERTIONS

	assert.Contains(t, msg, "Test Title")
	assert.Contains(t, msg, "More information about the error")
}

func TestUnit_Error_WhenTitleAndSourceMessageSet(t *testing.T) {
	// SETUP

	srcMsg := "Message from source"

	herr := New("Test Title")
	herr.Source.Message = &srcMsg

	// ACTIONS

	msg := herr.Error()

	// ASSERTIONS

	assert.Contains(t, msg, "Test Title")
	assert.Contains(t, msg, "Message from source")
}
