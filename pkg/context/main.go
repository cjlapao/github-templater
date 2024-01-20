package context

import (
	"context"

	logger "github.com/cjlapao/common-go-logger"
	"github.com/cjlapao/github-templater/constants"
)

var globalContext *ProvisionerContext

type ProvisionerContext struct {
	ctx       context.Context
	shouldLog bool
	logger    *logger.LoggerService
}

func NewWithContext(ctx context.Context) *ProvisionerContext {
	result := &ProvisionerContext{
		ctx:       ctx,
		shouldLog: true,
		logger:    logger.Get(),
	}

	result.logger.EnableTimestamp(true)
	return result
}

func New() *ProvisionerContext {
	return NewWithContext(context.Background())
}

func Get() *ProvisionerContext {
	if globalContext == nil {
		globalContext = New()
	}

	return globalContext
}

func (p *ProvisionerContext) DisableLogs() {
	p.shouldLog = false
}

func (p *ProvisionerContext) EnableLogs() {
	p.shouldLog = false
}

func (p *ProvisionerContext) Context() context.Context {
	return p.ctx
}

func (p *ProvisionerContext) SetContext(ctx context.Context) {
	p.ctx = ctx
}

func (p *ProvisionerContext) SetRequestId(id string) {
	p.ctx = context.WithValue(p.ctx, constants.ContextRequestID, id)
}

func (p *ProvisionerContext) RequestId() string {
	requestId := p.ctx.Value(constants.ContextRequestID)
	if requestId == nil {
		return ""
	}
	return requestId.(string)
}

func (p *ProvisionerContext) WithValue(key string, value interface{}) {
	p.ctx = context.WithValue(p.ctx, key, value)
}

func (p *ProvisionerContext) Logger() *logger.LoggerService {
	return p.logger
}

func (p *ProvisionerContext) LogError(format string, words ...interface{}) {
	if p.shouldLog {
		p.logger.Error(p.getLoggerFormattedMessage(format), words...)
	}
}

func (p *ProvisionerContext) LogInfo(format string, words ...interface{}) {
	if p.shouldLog {
		p.logger.Info(p.getLoggerFormattedMessage(format), words...)
	}
}

func (p *ProvisionerContext) LogDebug(format string, words ...interface{}) {
	if p.shouldLog {
		p.logger.Debug(p.getLoggerFormattedMessage(format), words...)
	}
}

func (p *ProvisionerContext) LogTrace(format string, words ...interface{}) {
	if p.shouldLog {
		p.logger.Trace(p.getLoggerFormattedMessage(format), words...)
	}
}

func (p *ProvisionerContext) getLoggerFormattedMessage(format string) string {
	requestID := p.RequestId()
	if requestID != "" {
		format = "[" + requestID + "] " + format
	}
	resourceId := p.ctx.Value(constants.ContextResourceID)
	resourceName := p.ctx.Value(constants.ContextResourceName)
	if resourceId != nil {
		format = "[" + resourceId.(string) + "] " + format
	}
	if resourceId == nil && resourceName != nil {
		format = "[" + resourceName.(string) + "] " + format
	}

	return format
}
