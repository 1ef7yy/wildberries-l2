package view

import (
	"log/slog"
	"wildberries-l2/grep/internal/domain"
)

type View struct {
	log    slog.Logger
	domain *domain.Domain
}

func NewView(log slog.Logger) *View {
	return &View{
		log:    log,
		domain: domain.NewDomain(log),
	}
}