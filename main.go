package main

import (
	"os"

	"github.com/th3oth3rjak3/logger/tinylog"
)

func main() {
	lgr := tinylog.New(
		tinylog.LevelDebug,
		tinylog.WithOutput(os.Stdout),
		tinylog.WithTimestamp(),
		tinylog.WithLogLevel(),
	)

	lgr.Infof("A little copying is better than a little dependency.")
	lgr.Errorf("Errors are values. Documentation is for %s", "users")
	lgr.Debugf("Make the zero (%d) value useful.", 0)

	lgr.Infof("Hello, %s", "world")
}
