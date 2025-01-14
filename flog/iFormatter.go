package flog

import (
	"encoding/json"
	"fmt"
	"github.com/farseer-go/fs/core/eumLogLevel"
)

// IFormatter 日志格式
type IFormatter interface {
	Formatter(log *logData) string
}

// JsonFormatter json格式输出
type JsonFormatter struct {
}

func (r JsonFormatter) Formatter(log *logData) string {
	marshal, _ := json.Marshal(logData{
		CreateAt:  log.CreateAt,
		LogLevel:  log.LogLevel,
		Component: log.Component,
		Content:   mustCompile.ReplaceAllString(log.Content, ""),
		newLine:   log.newLine,
	})
	return string(marshal)
}

// TextFormatter 文本格式输出
type TextFormatter struct {
	config levelFormat
}

func (r TextFormatter) Formatter(log *logData) string {
	var logLevelString string
	if log.LogLevel != eumLogLevel.NoneLevel {
		logLevelString = Colors[log.LogLevel]("[" + log.LogLevel.ToString() + "]")

	} else if log.Component != "" {
		logLevelString = Colors[0]("[" + log.Component + "]")
	}

	if logLevelString != "" {
		return fmt.Sprintf("%s %s %s", log.CreateAt.ToString(r.config.TimeFormat), logLevelString, log.Content)
	}

	return fmt.Sprintf("%s %s", log.CreateAt.ToString(r.config.TimeFormat), log.Content)
}
