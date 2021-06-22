package server

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenConn(t *testing.T) {
	testCases := []struct {
		name            string
		address         string
		proto           string
		expectedAddr    string
		expectedNetwork string
		expectedError   error
	}{
		{
			name:            "Socket opens",
			address:         ":8887",
			proto:           "tcp",
			expectedAddr:    "[::]:8887",
			expectedNetwork: "tcp",
			expectedError:   nil,
		},
		{
			name:            "Socket opens with default proto",
			address:         ":8887",
			expectedAddr:    "[::]:8887",
			expectedNetwork: "tcp",
			expectedError:   nil,
		},
		{
			name:            "Socket fails to open with udp proto",
			address:         ":8887",
			proto:           "udp",
			expectedAddr:    "[::]:8887",
			expectedNetwork: "udp",
			expectedError:   errors.New("failed to listen on :8887 for protocol udp"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sock, err := OpenConn(tc.address, tc.proto)
			if err != nil {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.Nil(t, err)
			}

			if sock != nil {
				defer sock.Close()
				assert.Equal(t, tc.expectedAddr, sock.Addr().String())
				assert.Equal(t, tc.expectedNetwork, sock.Addr().Network())
			}
		})
	}
}
