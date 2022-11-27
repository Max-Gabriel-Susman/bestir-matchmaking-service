package match

import (
	"context"

	"github.com/google/uuid"
)

func (api *API) CreateMatch(ctx context.Context, incomingmatch IncomingMatch) (Match, error) {
	id := uuid.New()

	match := Match{
		ID:   id,
		Name: incomingmatch.Name,
	}

	err := api.Store.CreateMatch(ctx, match)

	return match, err
}
