/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger/aries-framework-go/pkg/controller/command"
)

const (
	sampleErr1 = iota + command.UnknownStatus
	sampleErr2
	sampleErr3
	sampleErr4
)

func TestSendError(t *testing.T) {
	t.Run("Test sending HTTP status codes", func(t *testing.T) {
		const errMsg = "here is the sample which I want to write to response"
		var errors = []struct {
			err        error
			errCode    command.Code
			statusCode int
			response   genericError
		}{
			{fmt.Errorf(errMsg), sampleErr1, http.StatusOK,
				genericError{Code: sampleErr1, Message: errMsg}},
			{fmt.Errorf(errMsg), sampleErr2, http.StatusForbidden,
				genericError{Code: sampleErr2, Message: errMsg}},
			{fmt.Errorf(errMsg), sampleErr3, http.StatusNotAcceptable,
				genericError{Code: sampleErr3, Message: errMsg}},
			{fmt.Errorf(errMsg), sampleErr4, http.StatusNoContent,
				genericError{Code: sampleErr4, Message: errMsg}},
		}

		for _, data := range errors {
			rr := httptest.NewRecorder()

			SendHTTPStatusError(rr, data.statusCode, data.errCode, data.err)
			require.NotEmpty(t, rr.Body.Bytes())

			response := genericError{}
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			require.NoError(t, err)

			require.Equal(t, data.statusCode, rr.Code)
			require.Equal(t, data.response, response)
		}
	})

	t.Run("Test sending command errors", func(t *testing.T) {
		const errMsg = "here is the sample which I want to write to response"
		var errors = []struct {
			err        command.Error
			statusCode int
			response   genericError
		}{
			{command.NewValidationError(sampleErr1, fmt.Errorf(errMsg)), http.StatusBadRequest,
				genericError{Code: sampleErr1, Message: errMsg}},
			{command.NewExecuteError(sampleErr2, fmt.Errorf(errMsg)), http.StatusInternalServerError,
				genericError{Code: sampleErr2, Message: errMsg}},
			{command.NewValidationError(sampleErr3, fmt.Errorf(errMsg)), http.StatusBadRequest,
				genericError{Code: sampleErr3, Message: errMsg}},
			{command.NewExecuteError(sampleErr4, fmt.Errorf(errMsg)), http.StatusInternalServerError,
				genericError{Code: sampleErr4, Message: errMsg}},
		}

		for _, data := range errors {
			rr := httptest.NewRecorder()

			SendError(rr, data.err)
			require.NotEmpty(t, rr.Body.Bytes())

			response := genericError{}
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			require.NoError(t, err)

			require.Equal(t, data.statusCode, rr.Code)
			require.Equal(t, data.response, response)
		}
	})
}

func TestSendErrorFailures(t *testing.T) {
	rw := &mockRWriter{}
	SendHTTPStatusError(rw, http.StatusBadRequest, command.UnknownStatus, fmt.Errorf("sample error"))
}

// mockRWriter to recreate response writer error scenario
type mockRWriter struct {
}

func (m *mockRWriter) Header() http.Header {
	return make(map[string][]string)
}

func (m *mockRWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("failed to write body")
}

func (m *mockRWriter) WriteHeader(statusCode int) {}
