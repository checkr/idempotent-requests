package captures_mongo

import (
	"context"
)

func (repo *RepositoryImpl) Ready(ctx context.Context) bool {
	return repo.ready
}
