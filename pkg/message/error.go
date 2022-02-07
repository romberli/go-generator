package message

import (
	"github.com/romberli/go-util/config"
)

const (
	ErrPrintHelpInfo              = 400001
	ErrEmptyLogFileName           = 400002
	ErrNotValidLogFileName        = 400003
	ErrNotValidLogLevel           = 400004
	ErrNotValidLogFormat          = 400005
	ErrNotValidLogMaxSize         = 400006
	ErrNotValidLogMaxDays         = 400007
	ErrNotValidLogMaxBackups      = 400008
	ErrNotValidServerPort         = 400009
	ErrNotValidServerAddr         = 400010
	ErrNotValidPidFile            = 400011
	ErrNotValidServerReadTimeout  = 400012
	ErrNotValidServerWriteTimeout = 400013
	ErrValidateConfig             = 400014
	ErrInitDefaultConfig          = 400015
	ErrReadConfigFile             = 400016
	ErrOverrideCommandLineArgs    = 400017
	ErrAbsoluteLogFilePath        = 400018
	ErrInitLogger                 = 400019
	ErrBaseDir                    = 400020
	ErrInitConfig                 = 400021
	ErrEmptyStructFile            = 400022
	ErrNotValidStructFile         = 400023
	ErrReadStructFile             = 400024
	ErrEmptyOutputFile            = 400025
	ErrNotValidOutputFile         = 400026
	ErrWriteOutputFile            = 400027
)

func initErrorMessage() {
	Messages[ErrPrintHelpInfo] = config.NewErrMessage(DefaultMessageHeader, ErrPrintHelpInfo, "got message when printing help information")
	Messages[ErrEmptyLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyLogFileName, "log file name could not be an empty string")
	Messages[ErrNotValidLogFileName] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFileName, "log file name must be either unix or windows path format, %s is not valid")
	Messages[ErrNotValidLogLevel] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogLevel, "log level must be one of [debug, info, warn, message, fatal], %s is not valid")
	Messages[ErrNotValidLogFormat] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogFormat, "log level must be either text or json, %s is not valid")
	Messages[ErrNotValidLogMaxSize] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxSize, "log max size must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxDays] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxDays, "log max days must be between %d and %d, %d is not valid")
	Messages[ErrNotValidLogMaxBackups] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidLogMaxBackups, "log max backups must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerPort] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerPort, "server port must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerAddr] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerAddr, "server addr must be formatted as host:port, %s is not valid")
	Messages[ErrNotValidPidFile] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidPidFile, "pid file name must be either unix or windows path format, %s is not valid")
	Messages[ErrNotValidServerReadTimeout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerReadTimeout, "server read timeout must be between %d and %d, %d is not valid")
	Messages[ErrNotValidServerWriteTimeout] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidServerWriteTimeout, "server write timeout must be between %d and %d, %d is not valid")
	Messages[ErrValidateConfig] = config.NewErrMessage(DefaultMessageHeader, ErrValidateConfig, "validate config failed")
	Messages[ErrInitDefaultConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitDefaultConfig, "init default configuration failed")
	Messages[ErrReadConfigFile] = config.NewErrMessage(DefaultMessageHeader, ErrReadConfigFile, "read config file failed")
	Messages[ErrOverrideCommandLineArgs] = config.NewErrMessage(DefaultMessageHeader, ErrOverrideCommandLineArgs, "override command line arguments failed")
	Messages[ErrAbsoluteLogFilePath] = config.NewErrMessage(DefaultMessageHeader, ErrAbsoluteLogFilePath, "get absolute path of log file failed. log file: %s")
	Messages[ErrInitLogger] = config.NewErrMessage(DefaultMessageHeader, ErrInitLogger, "initialize logger failed")
	Messages[ErrBaseDir] = config.NewErrMessage(DefaultMessageHeader, ErrBaseDir, "get base dir of %s failed")
	Messages[ErrInitConfig] = config.NewErrMessage(DefaultMessageHeader, ErrInitConfig, "init config failed")
	Messages[ErrEmptyStructFile] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyStructFile, "struct file could not be an empty string")
	Messages[ErrNotValidStructFile] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidStructFile, "struct file must be either unix or windows path format, %s is not valid")
	Messages[ErrReadStructFile] = config.NewErrMessage(DefaultMessageHeader, ErrReadStructFile, "read struct file failed")
	Messages[ErrEmptyOutputFile] = config.NewErrMessage(DefaultMessageHeader, ErrEmptyOutputFile, "output file could not be an empty string")
	Messages[ErrNotValidOutputFile] = config.NewErrMessage(DefaultMessageHeader, ErrNotValidOutputFile, "output file must be either unix or windows path format, %s is not valid")
	Messages[ErrWriteOutputFile] = config.NewErrMessage(DefaultMessageHeader, ErrWriteOutputFile, "write output file failed")
}
