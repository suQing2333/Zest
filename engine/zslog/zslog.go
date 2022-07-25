package zslog

//日志级别 时间 年月日时分秒毫秒 输出文件 文件行数 内容

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
	"zest/engine/binutil"
)

var (
	// DebugLevel level
	DebugLevel = zapcore.DebugLevel
	// InfoLevel level
	InfoLevel = zapcore.InfoLevel
	// WarnLevel level
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel level
	ErrorLevel = zapcore.ErrorLevel
	// PanicLevel level
	PanicLevel = zapcore.PanicLevel
	// FatalLevel level
	FatalLevel = zapcore.FatalLevel

	sugar *zap.SugaredLogger
	cfg   zap.Config

	outputDir = "./"
	outPath   = "os"
	panicPath = "./panic.log"
	LogLevel  = zapcore.DebugLevel
	// 设置一些基本日志格式
	encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "caller",
		StacktraceKey: "trace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
)

// func init() {
// 	buildLog()
// }

func SetupLog(dir string, fileName string, level zapcore.Level) {
	outputDir = dir
	outPath = fileName
	LogLevel = level
	buildLog()
	setupPanic(dir, fileName)
}

// 重定向panic输出
func setupPanic(dir string, fileName string) {
	panicPath := filepath.Join(dir, "debug")
	os.MkdirAll(panicPath, os.ModeDir|os.ModePerm)
	f, err := os.OpenFile(filepath.Join(panicPath, fileName+"_panic.log"), os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	binutil.RedirectStderr(f)
}

func buildLog() {
	_, err := os.Stat(outputDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(outputDir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			}
		}
	}
	// 实现两个判断日志等级的interface
	logLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	LogHook := getWriter(outputDir, outPath)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(LogHook), logLevel),
	)

	// zap.AddCaller() 打日志点的文件名和行数
	// zap.AddCallerSkip() 调整打印位置层级
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	defer logger.Sync()
	sugar = logger.Sugar()
}

func getWriter(outputDir string, filename string) io.Writer {
	hook, err := rotatelogs.New(
		// outputDir+filename+"_%Y%m%d"+".log",
		filepath.Join(outputDir, filename)+"_%Y%m%d"+".log",
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func SetOutPutDir(Path string) {
	outputDir = Path
	buildLog()
}

func SetFileName(fileName string) {
	outPath = fileName
	buildLog()
}

func SetLogLevel(Level zapcore.Level) {
	LogLevel = Level
	buildLog()
}

func ParseLevel(s string) zapcore.Level {
	if strings.ToLower(s) == "debug" {
		return DebugLevel
	} else if strings.ToLower(s) == "info" {
		return InfoLevel
	} else if strings.ToLower(s) == "warn" || strings.ToLower(s) == "warning" {
		return WarnLevel
	} else if strings.ToLower(s) == "error" {
		return ErrorLevel
	} else if strings.ToLower(s) == "panic" {
		return PanicLevel
	} else if strings.ToLower(s) == "fatal" {
		return FatalLevel
	}
	LogDebug("ParseLevel: unknown level: %s", s)
	return DebugLevel
}

// 重新封装一下日志输出接口
func LogDebug(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func LogInfo(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

func LogWarn(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

func LogError(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func LogFatal(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

func LogPanic(template string, args ...interface{}) {
	sugar.Panicf(template, args...)
}
