package log

func Trace(msg string) {
	logger.Trace().Msg(msg)
}

func Tracef(format string, args ...interface{}) {
	logger.Trace().Msgf(format, args...)
}

func Debug(msg string) {
	logger.Debug().Msg(msg)
}

func Debugf(format string, args ...interface{}) {
	logger.Debug().Msgf(format, args...)
}

func Info(msg string) {
	logger.Info().Msg(msg)
}

func Infof(format string, args ...interface{}) {
	logger.Info().Msgf(format, args...)
}

func Warning(msg string) {
	logger.Warn().Msg(msg)
}

func Warningf(format string, args ...interface{}) {
	logger.Warn().Msgf(format, args...)
}

func Error(msg string, err error) {
	logger.Error().Err(err).Msg(msg)
}

func Errorf(format string, args ...interface{}) {
	logger.Error().Msgf(format, args...)
}

func Fatal(msg string, err error) {
	logger.Fatal().Err(err).Msg(msg)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatal().Msgf(format, args...)
}

func Panic(msg string, err error) {
	logger.Panic().Err(err).Msg(msg)
}

func Panicf(format string, args ...interface{}) {
	logger.Panic().Msgf(format, args...)
}
