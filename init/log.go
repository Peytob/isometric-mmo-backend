package init

import (
	"io"
	"isonetric-mmo-backend/internal/model"
	log "log/slog"
	"os"
)

func Logging(config *model.Config) error {
	if config.Logging.Disabled {
		log.SetDefault(log.New(log.NewTextHandler(io.Discard, nil)))
		return nil
	}

	log.SetDefault(log.New(log.NewJSONHandler(os.Stdout, nil)))

	if level, err := parseLoggingLevel(config.Logging.Level); err != nil {
		return err
	} else {
		log.SetLogLoggerLevel(level)
	}

	return nil
}

func parseLoggingLevel(s string) (log.Level, error) {
	var level log.Level
	var err = level.UnmarshalText([]byte(s))
	return level, err
}
