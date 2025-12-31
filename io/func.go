package io

import (
	"fmt"
	"io"
	"log/slog"
)

// CloseWithLog to log error when close io.Closer, if error occurs,
// tag will be used as log tag
//
// 关闭 io.Closer 时，如果发生错误则记录日志，将使用 tag 参数作为日志标签
func CloseWithLog(closer io.Closer, tag string) {
	err := closer.Close()
	if err != nil {
		slog.Error(fmt.Sprintf("[%s] close failed", tag), "error", err)
	}
}
