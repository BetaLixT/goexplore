package forcast

import "context"

type Queries struct {
  repo IQueryRepository
}

func NewQueries (
  repo IQueryRepository,
) *Queries {
  return &Queries{
    repo: repo,
  }
}

func (qry *Queries) GetForcast (
  ctx context.Context,
  id int,
) (Forcast, error) {
  return qry.repo.Get(
    ctx,
    id,
  )
}

func (qry *Queries) ListForcasts (
  ctx context.Context,
) ([]Forcast, error) {
  return qry.repo.List(
    ctx,
  )
}
