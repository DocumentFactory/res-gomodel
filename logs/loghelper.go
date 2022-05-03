package logs

import (
	"github.com/pnocera/res-gomodel/config"
	"go.uber.org/zap"
)

type LogHelper struct {
	conf   *config.Config
	logger *zap.Logger
	info   bool
	warn   bool
	debug  bool
	errors bool
}

func NewLogHelper(conf *config.Config) *LogHelper {
	c := LogHelper{
		conf:   conf,
		errors: true,
		info:   true,
		warn:   true,
		debug:  true,
	}
	logger, _ := zap.NewProduction()

	level := conf.GetLogLevel()

	if level == "none" {
		c.errors = false
		c.info = false
		c.warn = false
		c.debug = false
	}

	if level == "error" {
		c.info = false
		c.warn = false
		c.debug = false
	}

	if level == "info" {
		c.warn = false
		c.debug = false
	}

	if level == "warn" {
		c.debug = false
	}

	c.logger = logger
	return &c
}

func (lh *LogHelper) Info(msg string, fields ...zap.Field) {
	if lh.info {
		lh.logger.Info(msg, fields...)
	}
}

func (lh *LogHelper) Error(msg string, fields ...zap.Field) {
	if lh.errors {
		lh.logger.Error(msg, fields...)
	}
}

func (lh *LogHelper) Warn(msg string, fields ...zap.Field) {
	if lh.warn {
		lh.logger.Warn(msg, fields...)
	}
}

func (lh *LogHelper) Debug(msg string, fields ...zap.Field) {
	if lh.debug {
		lh.logger.Debug(msg, fields...)
	}
}

func (lh *LogHelper) Fatal(msg string, fields ...zap.Field) {
	if lh.errors {
		lh.logger.Fatal(msg, fields...)
	}
}

func (lh *LogHelper) Panic(msg string, fields ...zap.Field) {
	if lh.errors {
		lh.logger.Panic(msg, fields...)
	}
}
