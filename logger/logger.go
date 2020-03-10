package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Logger *zap.Logger

/**
 * debug 是否debug模式 debug模式使用zap的DevelopmentEncoderConfig
 * filePath 文件保存路径
 * serviceName 服务名称
 * stdOut 是否控制台输出
 * withCaller 是否开启堆栈跟踪
 */
func InitLogger(debug bool, filePath, serviceName string, stdOut bool, withCaller bool) {
	filename := filePath + "/" + serviceName + ".log"
	hook := lumberjack.Logger{
		Filename:   filename, // 日志文件路径
		MaxSize:    1024,     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,       // 最多保存多少个日志文件
		MaxAge:     30,       // 文件最多保存多少天
		Compress:   true,     // 是否压缩
	}

	var encoderConfig zapcore.EncoderConfig
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	if debug {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	var ws zapcore.WriteSyncer
	if stdOut {
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		ws = zapcore.AddSync(&hook)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		ws,                                    // 打印到控制台和文件
		atomicLevel,                           // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	callerSkip := zap.AddCallerSkip(1)
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", serviceName))
	// 构造日志
	if withCaller {
		Logger = zap.New(core, caller, callerSkip, development, filed)
	} else {
		Logger = zap.New(core, development, filed)
	}

	Logger.Info("log 初始化成功")
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Debugf(template string, args... interface{}) {
	Logger.Debug(fmt.Sprintf(template, args...))
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Infof(template string, args... interface{}) {
	Logger.Info(fmt.Sprintf(template, args...))
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}
func Warnf(template string, args... interface{}) {
	Logger.Warn(fmt.Sprintf(template, args...))
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Errorf(template string, args... interface{}) {
	Logger.Error(fmt.Sprintf(template, args...))
}

func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}
func Panicf(template string, args... interface{}) {
	Logger.Panic(fmt.Sprintf(template, args...))
}
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}
func Fatalf(template string, args... interface{}) {
	Logger.Fatal(fmt.Sprintf(template, args...))
}
