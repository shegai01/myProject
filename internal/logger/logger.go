package zaplog

import "go.uber.org/zap"

func ZapLog() *zap.SugaredLogger {
	configLogger, err := zap.NewProduction()
	if err != nil {
		return nil
	}
	defer configLogger.Sync()

	return configLogger.Sugar()
}
