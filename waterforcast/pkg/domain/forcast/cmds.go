package forcast

import "context"

type Commands struct {
  repo ICommandRepository
}

func NewCommands(
  repo ICommandRepository,
) *Commands {
  return &Commands{
    repo: repo,
  }
}

func (cmd *Commands) AddForcast (
  ctx context.Context,
  id int,
  title string,
) (Forcast, error) {
  return cmd.repo.Create(
    ctx,
    id,
    title,
  )
}
