package domains

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"main.go/pkg/domains/beer"
	"main.go/pkg/domains/financial"
	"main.go/pkg/domains/room"
)

type Domain struct {
  bd beer.BeerDomain
  fd financial.FinancialDomain
  rd room.RoomDomain
}

func NewDomain(
  bd beer.BeerDomain,
  fd financial.FinancialDomain,
  rd room.RoomDomain,
) *Domain {
  return &Domain{
    bd: bd,
    fd: fd,
    rd: rd,
  }
}

func (d *Domain) RegisterCommandQueryHandlers (
  cb *cqrs.CommandBus,
  eb *cqrs.EventBus,
) []cqrs.CommandHandler {
  cqs :=  []cqrs.CommandHandler {
  }
  cqs = append(cqs, d.bd.RegisterCommandQueryHandlers(cb, eb)...)
  cqs = append(cqs, d.fd.RegisterCommandQueryHandlers(cb, eb)...)
  cqs = append(cqs, d.rd.RegisterCommandQueryHandlers(cb, eb)...)
  return cqs
}

func (d *Domain) RegisterEventHandlers (
  cb *cqrs.CommandBus,
  eb *cqrs.EventBus,
) []cqrs.EventHandler {
  evnts := []cqrs.EventHandler {
  }
  evnts = append(evnts, d.bd.RegisterEventHandlers(cb, eb)...)
  evnts = append(evnts, d.fd.RegisterEventHandlers(cb, eb)...)
  evnts = append(evnts, d.rd.RegisterEventHandlers(cb, eb)...)
  return evnts
}
