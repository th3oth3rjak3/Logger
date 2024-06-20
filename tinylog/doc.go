/*
Package tinylog exposes an API to perform application logging.

First, instantiate a logger with tinylog.New, giving it a threshold level.
Messages with lower severity won't be logged.

Sharing the logger is the responsibility of the caller.

The logger can be called to log messages on three levels:
  - Debug: mostly used to debug code, follow step-by-step processes
  - Info: valuable messages providing insights into the application
  - Error: error messages to understand what went wrong
*/
package tinylog
