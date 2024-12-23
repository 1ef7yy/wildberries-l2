package main

import (
	"log"
	"log/slog"
	"os"
	"wildberries-l2/wget/internal/view"
)

func main() {
	logger := slog.New(slog.NewTextHandler(log.Writer(), nil))
	view := view.NewView(*logger)

	view.Wget(os.Args[:])

}
