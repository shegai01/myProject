package logger

import "go.uber.org/zap"

func ZapLog() (*zap.SugaredLogger, error) {
	configLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer configLogger.Sync()

	return configLogger.Sugar(), nil
}
