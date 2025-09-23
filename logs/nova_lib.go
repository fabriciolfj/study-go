package logs

import (
	"log/slog"
)

func Execute3() {
	slog.Info("this is default logging")
	slog.Warn("keep an eye on this, it might be an issue")
	slog.Error("oh no, an error happened here!")
	slog.Debug("this is good while developing ...")
}
