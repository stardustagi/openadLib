package logger

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var (
	logger hclog.Logger
	once   sync.Once
)

type Config struct {
	LogName       string `json:"log_name" toml:"log_name"`
	LogPath       string `json:"log_patch" toml:"log_path"`
	LogLevel      string `json:"log_level" toml:"log_level"`
	LogRotate     string `json:"log_rotate" toml:"log_rotate"`
	LogRotateSize int64  `json:"log_rotate_size" toml:"log_rotate_size"`
}

func SetupLogger(conf Config) {
	once.Do(func() {
		var fileOption []FileOption
		switch RotatePolicy(conf.LogRotate) {
		case ROTATE_POLICY_DAY:
			fileOption = append(fileOption, WithRotatePolicy(ROTATE_POLICY_DAY))
		case ROTATE_POLICY_SIZE:
			fileOption = append(fileOption, WithRotatePolicy(ROTATE_POLICY_SIZE))
		case ROTATE_POLICY_HOUR:
			fileOption = append(fileOption, WithRotatePolicy(ROTATE_POLICY_HOUR))
		default:
			break
		}
		fileOption = append(fileOption, WithRotateSizeMB(conf.LogRotateSize))
		var fw *FileWriter
		var err error
		if fw, err = NewFileWriter(filepath.Join(conf.LogPath, conf.LogName), fileOption...); err != nil {
			fmt.Printf("Error creating file writer: %s", err.Error())
		}

		multiWriter := io.MultiWriter(os.Stdout, fw) //输控制台和文件
		logger = hclog.New(&hclog.LoggerOptions{
			Name:                     conf.LogName,
			Level:                    hclog.LevelFromString(conf.LogLevel),
			Output:                   multiWriter,
			IncludeLocation:          true,
			IndependentLevels:        true,
			AdditionalLocationOffset: 1,
			JSONFormat:               true,
			JSONEscapeDisabled:       true,
		})
		logger.SetLevel(hclog.LevelFromString(conf.LogLevel))
	})
}

func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	logger.Warn(msg, args...)
}

func Trace(msg string, args ...interface{}) {
	logger.Trace(msg, args...)
}

func Fmt(msg string, args ...interface{}) interface{} {
	return hclog.Fmt(msg, args...)
}

func GetLogger() hclog.Logger {
	return logger
}
