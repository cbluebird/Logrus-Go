package logs

import (
	"fmt"
	"logs/app/apiException"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Writewarn(c *gin.Context, e *apiException.Error) {
	Log.WithFields(logrus.Fields{
		"code":    e.Code,
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
		"address": c.Request.RemoteAddr,
	}).Warn(e.Msg)
}

func WriteDebug(c *gin.Context, e *apiException.Error) {
	Log.WithFields(logrus.Fields{
		"code":    e.Code,
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
		"address": c.Request.RemoteAddr,
	}).Debug(e.Msg)
}

func WriteError(c *gin.Context, e error) {
	if stackErr, ok := e.(stackTracer); ok {
		Log.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"address":    c.Request.RemoteAddr,
			"stacktrace": fmt.Sprintf("%+v", stackErr.StackTrace()),
		}).Error("an error can be stacktrace")
	} else {
		Log.WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"address": c.Request.RemoteAddr,
		}).Error("unknow error cannot stacktrace")
	}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
