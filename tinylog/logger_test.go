package tinylog_test

import (
	"testing"

	"github.com/th3oth3rjak3/logger/tinylog"
)

// testWriter is a testing type that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level    tinylog.Level
		expected string
	}

	testCases := map[string]testCase{
		"debug": {
			level:    tinylog.LevelDebug,
			expected: debugMessage + "\n" + infoMessage + "\n" + errorMessage + "\n",
		},
		"info": {
			level:    tinylog.LevelInfo,
			expected: infoMessage + "\n" + errorMessage + "\n",
		},
		"error": {
			level:    tinylog.LevelError,
			expected: errorMessage + "\n",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := tinylog.New(testCase.level, tinylog.WithOutput(tw))
			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			if tw.contents != testCase.expected {
				t.Errorf("invalid contents, got: %s, want: %s", tw.contents, testCase.expected)
			}
		})
	}
}
