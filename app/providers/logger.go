package providers

import (
	"L/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Option struct {
	level   string
	service string
}

type HandleOption func(opt *Option)

func WithLevel(level string) HandleOption {
	return func(opt *Option) {
		opt.level = level
	}
}

func WithService(service string) HandleOption {
	return func(opt *Option) {
		opt.service = service
	}
}

func NewLogger(handleOptions ...HandleOption) *zap.Logger {
	opt := Option{
		level:   "info",
		service: "system",
	}

	for _, handleOption := range handleOptions {
		handleOption(&opt)
	}

	cfg := NewConfig()
	return newZap(&cfg.Log, opt.service)
}

func newZap(logCfg *config.Log, serviceName string) *zap.Logger {
	filePath := logCfg.Path + logCfg.FileName
	var level zapcore.Level
	switch logCfg.Level {
	case "info":
		level = zap.InfoLevel
	case "debug":
		level = zap.DebugLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	core := newZapCore(filePath, level, logCfg.MaxSize, logCfg.MaxBackups, logCfg.MaxAge, logCfg.Compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("service", serviceName)))
}

func newZapCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	// 日志文件路径配置
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	// 公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
}
