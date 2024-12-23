package domain

import "log/slog"

type Domain struct {
	log slog.Logger
}

func NewDomain(log slog.Logger) *Domain {
	return &Domain{log: log}
}
