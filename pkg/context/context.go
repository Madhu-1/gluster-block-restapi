package context

import (
	"context"

	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

// ref: github.com/gluster/glusterd2/glusterd2/middleware/request_id.go
type ctxKeyType int

const (
	reqIDKey ctxKeyType = iota
	reqLoggerKey
)

// WithReqID returns a new context with provided request id set as a value in the context.
func WithReqID(ctx context.Context, reqid uuid.UUID) context.Context {
	return context.WithValue(ctx, reqIDKey, reqid)
}

// ReqID returns request ID stored in the context provided.
func ReqID(ctx context.Context) uuid.UUID {
	reqid, ok := ctx.Value(reqIDKey).(uuid.UUID)
	if !ok {
		return nil
	}
	return reqid
}

// WithReqLogger returns a new context with provided logger set as a value in the context.
func WithReqLogger(ctx context.Context, logger log.FieldLogger) context.Context {
	return context.WithValue(ctx, reqLoggerKey, logger)
}

// ReqLogger returns logger stored in the context provided.
func ReqLogger(ctx context.Context) log.FieldLogger {
	reqLogger, ok := ctx.Value(reqLoggerKey).(log.FieldLogger)
	if !ok {
		return log.StandardLogger()
	}
	return reqLogger
}
