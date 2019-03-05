package herror

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

////////////////////////////////////////////////////////////////////////////////
// CausedBy

func TestUnit_CausedBy_PopulatesSourceCausedBy(t *testing.T) {
	// SETUP

	origErr := errors.New("test error")

	// ACTIONS

	result := New("Test Title", CausedBy(origErr))

	// ASSERTIONS

	require.NotNil(t, result.Source.CausedBy)
	assert.Equal(t, origErr, result.Source.CausedBy)
}

func TestUnit_CausedBy_PopulatesSourceMessage(t *testing.T) {
	// SETUP

	origErr := errors.New("test error")

	// ACTIONS

	result := New("Test Title", CausedBy(origErr))

	// ASSERTIONS

	require.NotNil(t, result.Source.Message)
	assert.Equal(t, "test error", *result.Source.Message)
}

////////////////////////////////////////////////////////////////////////////////
// Code

func TestUnit_Code_PopulatesCode(t *testing.T) {
	// SETUP

	// ACTIONS

	result := New("Test Title", Code("test-error"))

	// ASSERTIONS

	require.NotNil(t, result.Code)
	assert.Equal(t, "test-error", *result.Code)
}

////////////////////////////////////////////////////////////////////////////////
// Detail

func TestUnit_Detail_WithSubstitutions_PopulatesDetail(t *testing.T) {
	// SETUP

	// ACTIONS

	result := New("Test Title", Detail("I am %d %s", 1, "teapot"))

	// ASSERTIONS

	require.NotNil(t, result.Detail)
	assert.Equal(t, "I am 1 teapot", *result.Detail)
}

func TestUnit_Detail_WithoutSubstitutions_PopulatesDetail(t *testing.T) {
	// SETUP

	// ACTIONS

	result := New("Test Title", Detail("I am a detail message"))

	// ASSERTIONS

	require.NotNil(t, result.Detail)
	assert.Equal(t, "I am a detail message", *result.Detail)
}

////////////////////////////////////////////////////////////////////////////////
// Status

func TestUnit_Status_PopulatesStatus(t *testing.T) {
	// SETUP

	// ACTIONS

	result := New("Test Title", Status(http.StatusNoContent))

	// ASSERTIONS

	require.NotNil(t, result.Status)
	assert.Equal(t, "204", *result.Status)
}

////////////////////////////////////////////////////////////////////////////////
// Title

func TestUnit_Title_PopulatesTitle(t *testing.T) {
	// SETUP

	// ACTIONS

	result := New("First Title", Title("Second Title"))

	// ASSERTIONS

	assert.Equal(t, "Second Title", result.Title)
}
