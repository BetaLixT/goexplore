package financial

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"main.go/pkg/domains/financial/evts"
)

type FinancialDomain struct {

}

func NewFinancialDomain() *FinancialDomain {
  return &FinancialDomain{}
}

func (b *FinancialDomain) RegisterCommandQueryHandlers (
  cb *cqrs.CommandBus,
  eb *cqrs.EventBus,
) []cqrs.CommandHandler {
  return []cqrs.CommandHandler {}
}

func (b *FinancialDomain) RegisterEventHandlers (
  cb *cqrs.CommandBus,
  eb *cqrs.EventBus,
) []cqrs.EventHandler {
  return []cqrs.EventHandler {
    evts.NewBookingsFinancialReport(),
  }
}
