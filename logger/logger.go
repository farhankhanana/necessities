// Package logger implements [uber-go/zap] library.
//
// Log level based on uber-go/zap:
// DEBUG < INFO < WARN < ERROR < DPANIC < PANIC < FATAL
// -1		 0		1	    2	    3		4		5
// [uber-go/zap]: https://github.com/uber-go/zap
package logger

import (
	"log"
	"os"
	"strings"

	"github.com/gat/necessities/utils"
	"github.com/showa-93/go-mask"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// environment constants
const (
	environmentDevelopment = "development"
	environmentProduction  = "production"
)

// constants that used as key when using JSONEncoder in zap
const (
	logProjectNameKey   = "project_name"
	logProjectModuleKey = "project_module"
	logModuleVersionKey = "module_version"
	logDataKey          = "log_data"
	logCorrelationIDKey = "correlation_id"
)

var (
	logProjectName   = "ces"
	logProjectModule = "<EMPTY_PROJECT_NAME>"
	logModuleVersion = "<EMPTY_MODULE_VERSION>"
)

var logger *zap.Logger

type Logger struct {
	correlationID string
}

// InitLogger initiates logger configuration.
// This function implement uber-go/zap library.
//
// `logLevel` will be selected to debug level if empty string is given or not registered in enums. Available `logLevel`:
//  1. Debug: set `-1` or pass empty string
//  2. Info: set `0` or `info`
//  3. Warn: set `1` or `warn`
//  4. Error: set `2` or `error`
//  5. DPanic: set `3` or `dpanic`
//  6. Panic: set `4` or `panic`
//  7. Fatal: set `5` or `fatal`
//
// There are 2 modes and each mode has different logging format.
//
//  1. environment = development, Format = `<TIMESTAMP> <LEVEL> <CALLER> <FUNCTION_NAME> <MESSAGE> <LOG_DATA>`
//  2. environment = production, Format = `{"level":"<LEVEL>","timestamp":"<TIMESTAMP","caller":"<CALLER>","function_name":"<FUNCTION_NAME>","message":"<MESSAGE>", "log_data":<LOG_DATA>}`
func InitLogger(logLevel, environment, projectName, projectModule, moduleVersion string) {
	environment = strings.ToLower(environment)
	// check if the environment is empty or not development nor production
	if len(environment) == 0 || (environment != environmentDevelopment && environment != environmentProduction) {
		environment = environmentDevelopment
	}

	// check if projectName is not empty
	if len(projectName) > 0 {
		logProjectName = projectName
	}

	// check if projectModule is not empty
	if len(projectModule) > 0 {
		logProjectModule = projectModule
	}

	// check if projectModule is not empty
	if len(moduleVersion) > 0 {
		logModuleVersion = moduleVersion
	}

	// initiate default zap loglevel to debug
	zapLogLevel := zap.DebugLevel
	switch strings.ToLower(logLevel) {
	case "0", "info":
		zapLogLevel = zap.InfoLevel
	case "1", "warn":
		zapLogLevel = zap.WarnLevel
	case "2", "error":
		zapLogLevel = zap.ErrorLevel
	case "3", "dpanic":
		zapLogLevel = zap.DPanicLevel
	case "4", "panic":
		zapLogLevel = zap.PanicLevel
	case "5", "fatal":
		zapLogLevel = zap.FatalLevel
	}

	// define level enabler with requirements:
	// 1. debug until warn level using `os.stdout`
	// 2. error until fatal level using `os.stderr`
	stdOutLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		// check if the level is current configuration level AND level is below error level
		return level >= zapLogLevel && level < zapcore.ErrorLevel
	})

	stdErrLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		// check if the level is current configuration level AND level is above or same with error level
		return level >= zapLogLevel && level >= zapcore.ErrorLevel
	})

	stdOutSyncer := zapcore.Lock(os.Stdout)
	stdErrSyncer := zapcore.Lock(os.Stderr)

	zapEncoderConfig := zap.NewDevelopmentEncoderConfig()
	zapEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if environment == environmentProduction {
		zapEncoderConfig = zap.NewProductionEncoderConfig()
	}
	zapEncoderConfig.FunctionKey = "function_name"
	zapEncoderConfig.CallerKey = "caller"
	zapEncoderConfig.MessageKey = "message"
	zapEncoderConfig.TimeKey = "timestamp"
	zapEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// register all configuration to zap
	// If the environment is development, using console encoder
	// otherwise if is production, using json encoder
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(zapEncoderConfig), stdOutSyncer, stdOutLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(zapEncoderConfig), stdErrSyncer, stdErrLevel),
	)
	if environment == environmentProduction {
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(zapEncoderConfig), stdOutSyncer, stdOutLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(zapEncoderConfig), stdErrSyncer, stdErrLevel),
		)
	}

	// create zap.Logger instance with caller and skip caller 1
	// and then include constant projectName, projectModule, and also moduleVersion
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).With(
		zap.String(logProjectNameKey, logProjectName),
		zap.String(logProjectModuleKey, logProjectModule),
		zap.String(logModuleVersionKey, logModuleVersion),
	)
}

// NewLogger creates instance `Logger` with correlationID as parameter.
// If empty correlationID passed to this function, UUID will be generated as correlationID.
func NewLogger(correlationID string) *Logger {
	if logger == nil {
		log.Panicln("empty logger")
	}
	if len(correlationID) <= 0 {
		correlationID = utils.UUIDGenerator()
	}
	return &Logger{
		correlationID: correlationID,
	}
}

func (l *Logger) MaskLogData(data interface{}, fields ...string) interface{} {
	masker := mask.NewMasker()
	masker.RegisterMaskStringFunc(mask.MaskTypeFilled, masker.MaskFilledString)
	for _, field := range fields {
		masker.RegisterMaskField(field, "filled5")
	}
	masked, _ := masker.Mask(data)
	return masked
}

// processZapFields maps log metadata and information to zap.Field.
// Returns result of mapped fields.
func (l *Logger) processZapFields(data interface{}) (fields []zap.Field) {
	fields = append(fields, zap.String(logCorrelationIDKey, l.correlationID))
	fields = append(fields, zap.Any(logDataKey, data))
	return fields
}

func (l *Logger) GetCorrelationID() string {
	return l.correlationID
}

// LogDebug logs at `debug` level.
func (l *Logger) LogDebug(message string, data ...interface{}) {
	logger.Debug(message, l.processZapFields(data)...)
}

// LogInfo logs at `info` level.
func (l *Logger) LogInfo(message string, data ...interface{}) {
	logger.Info(message, l.processZapFields(data)...)
}

// LogWarn logs at `warn` level.
func (l *Logger) LogWarn(message string, data ...interface{}) {
	logger.Warn(message, l.processZapFields(data)...)
}

// LogError logs at `error` level.
func (l *Logger) LogError(message string, data ...interface{}) {
	logger.Error(message, l.processZapFields(data)...)
}

// LogDPanic logs at `dpanic` level.
func (l *Logger) LogDPanic(message string, data ...interface{}) {
	logger.DPanic(message, l.processZapFields(data)...)
}

// LogPanic logs at `panic` level.
func (l *Logger) LogPanic(message string, data ...interface{}) {
	logger.Panic(message, l.processZapFields(data)...)
}

// LogFatal logs at `fatal` level.
func (l *Logger) LogFatal(message string, data ...interface{}) {
	logger.Fatal(message, l.processZapFields(data)...)
}
