package domain

import "log/slog"

type Domain struct {
	log slog.Logger
}

func NewDomain(log slog.Logger) *Domain {
	return &Domain{log: log}
}

type Flags struct {
	After      int
	Before     int
	Context    int
	Count      int
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

type Result struct {
	Line string
	Num  int
}
