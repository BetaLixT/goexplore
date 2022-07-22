package forcast

import "context"

type ICommandRepository interface {
  Create(context.Context, int, string) (Forcast, error)
}

type IQueryRepository interface {
  Get(context.Context, int) (Forcast, error)
  List(context.Context) ([]Forcast, error)
}
