package view

import (
	"log/slog"
	"wildberries-l2/wget/internal/domain"
)

type View struct {
	domain *domain.Domain
	log    slog.Logger
}

func NewView(log slog.Logger) *View {
	return &View{
		log:    log,
		domain: domain.NewDomain(log),
	}
}
