package log

import (
	"net/http"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	errorutils "github.com/oapi-validator-echo-sample/utils/errors"
	"go.uber.org/zap"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func TraceLogger(traceID string) (logr.Logger, error) {
	devLog, err := zap.NewDevelopment()
	if err != nil {
		return logr.New(logr.LogSink(log.NullLogSink{})), errorutils.New(http.StatusServiceUnavailable, err.Error())
	}

	logger := zapr.NewLogger(devLog)
	logger = logger.WithValues("TraceID", traceID)
	return logger, nil
}
