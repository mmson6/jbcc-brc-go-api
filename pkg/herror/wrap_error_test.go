package herror

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

////////////////////////////////////////////////////////////////////////////////
// Wrap()

func TestUnit_Wrap_WithHorizonError_ReturnsTheError(t *testing.T) {
	// SETUP

	origErr := New("Test Horizon Error")

	// ACTIONS

	herr := Wrap(origErr)

	// ASSERTIONS

	assert.Equal(t, *origErr, herr)
	assert.Equal(t, origErr.ID, herr.ID)
}

func TestUnit_Wrap_WithNil_ReturnsTheError(t *testing.T) {
	// SETUP

	// ACTIONS

	herr := Wrap(nil)

	// ASSERTIONS

	assert.Equal(t, "Invalid Parameter", herr.Title)
}

func TestUnit_Wrap_WithNormalError_WrapsTheError(t *testing.T) {
	// SETUP

	origErr := errors.New("test error")

	// ACTIONS

	herr := Wrap(origErr)

	// ASSERTIONS

	require.NotNil(t, herr.Source.CausedBy)
	assert.Equal(t, origErr, herr.Source.CausedBy)
}

func TestUnit_Wrap_WithNormalError_WithOptions_AppliesOptions(t *testing.T) {
	// SETUP

	origErr := errors.New("test error")

	// ACTIONS

	herr := Wrap(origErr,
		Title("Title from Option"),
		Code("test-error"),
		Status(http.StatusConflict),
	)

	// ASSERTIONS

	assert.Equal(t, "Title from Option", herr.Title)
	require.NotNil(t, herr.Code)
	assert.Equal(t, "test-error", *herr.Code)
	require.NotNil(t, herr.Status)
	assert.Equal(t, "409", *herr.Status)
}
