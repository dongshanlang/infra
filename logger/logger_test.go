package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestMain(m *testing.M) {
	InitLogger(true, "./", "test", true, true)
	Debug("debug test", zap.Int32("uid", 1))
	Debugf("this is %s, age %d", "mxl", 18)
	Info("info test", zap.Int32("uid", 1))
	Infof("info test, uid:%d", 1)
	Warn("warn test", zap.Int32("uid", 1))
	Warnf("warn format test, uid:%d", 1)
	Error("error test", zap.Int32("uid", 1))
	Errorf("error format test, uid:%d", 1)
	Panic("panic test", zap.Int32("uid", 1))
	Panicf("pannic test, uid:%d", 1)
	Fatal("fatal test", zap.Int32("uid", 1))

}