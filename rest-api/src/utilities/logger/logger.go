package logger

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	log "github.com/yfuruyama/stackdriver-request-context-log"
)

var Log *log.ContextLogger

func Init(e *echo.Echo) http.Handler {
	config := log.NewConfig(os.Getenv("PROJECT_ID"))
	config.RequestLogOut = os.Stderr   // request log to stderr
	config.ContextLogOut = os.Stdout   // context log to stdout
	config.Severity = log.SeverityInfo // only over INFO logs are logged

	return log.RequestLogging(config)(e)
}

func SetupLogger(request *http.Request) {
	Log = log.RequestContextLogger(request)
}
