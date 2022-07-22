package repos

import (
	"context"
	"sync"
	"waterforcast/pkg/common"
	"waterforcast/pkg/domain/forcast"
)

type ForcastRepository struct {
	forcasts map[int]string
	mtx      sync.Mutex
}

func NewForcastRepository() *ForcastRepository {
	return &ForcastRepository{
		forcasts: map[int]string{},
	}
}

func (repo *ForcastRepository) Create(
	ctx context.Context,
	id int,
	title string,
) (forcast.Forcast, error) {
  repo.mtx.Lock()
  defer repo.mtx.Unlock()
	if _, exists := repo.forcasts[id]; exists {
		return forcast.Forcast{}, common.NewForcastConflictError()
	}

	repo.forcasts[id] = title
	return forcast.Forcast{
		Id:    id,
		Title: title,
	}, nil
}

func (repo *ForcastRepository) Get(
	ctx context.Context,
	id int,
) (forcast.Forcast, error) {
  repo.mtx.Lock()
  defer repo.mtx.Unlock()
	if val, exists := repo.forcasts[id]; !exists {
		return forcast.Forcast{}, common.NewForcastMissingError()
	} else {
		return forcast.Forcast{
			Id:    id,
			Title: val,
		}, nil
	}
}

func (repo *ForcastRepository) List(
	context.Context,
) ([]forcast.Forcast, error) {
  repo.mtx.Lock()
  defer repo.mtx.Unlock()
  list := make([]forcast.Forcast, len(repo.forcasts))
  idx := 0
  for key, val := range repo.forcasts {
    list[idx] = forcast.Forcast{
      Id: key,
      Title: val,
    }
    idx++
  }
  return list, nil
}
