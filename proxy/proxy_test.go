package proxy

import (
	"sync"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestNewProxy(t *testing.T) {

	tt := []struct {
		name          string
		incomeRoutes  map[string]string
		expectedProxy *Proxy
		expectedErr   error
		expectedLen   int
	}{
		{
			name: "success creation",
			incomeRoutes: map[string]string{
				"^/test": "http://test.com",
			},
			expectedErr: nil,
			expectedLen: 1,
		},
		{
			name: "bad url",
			incomeRoutes: map[string]string{
				"^/test": "bad-url",
			},
			expectedErr:   errors.New("could not parse a provided url bad-url: parse bad-url: invalid URI for request"),
			expectedProxy: nil,
		},
		{
			name: "bad regexp",
			incomeRoutes: map[string]string{
				"[0--1]": "http://test.com",
			},
			expectedErr:   errors.New("could not compile a regular expression [0--1]: error parsing regexp: invalid character class range: `0--`"),
			expectedProxy: nil,
		},
	}

	var wg sync.WaitGroup
	for _, tc := range tt {
		wg.Add(1)
		go t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()
			actual, err := NewProxy(tc.incomeRoutes)

			if tc.expectedErr == nil {
				require.NotNil(t, actual)
				require.Equal(t, tc.expectedLen, len(actual.routes))
				require.Equal(t, tc.expectedErr, err)
			}
			if tc.expectedErr != nil {
				require.Equal(t, tc.expectedProxy, actual)
				require.Equal(t, tc.expectedErr.Error(), err.Error())
			}

		})
	}
	wg.Wait()
}
