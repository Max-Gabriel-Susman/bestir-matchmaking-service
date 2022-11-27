package match

import (
	"context"

	"github.com/google/uuid"
)

func (api *API) DeleteMatch(ctx context.Context, incomingMatch IncomingMatch) (Match, error) {
	id := uuid.New()

	match := Match{
		ID:   id,
		Name: incomingMatch.Name,
	}

	err := api.Store.DeleteMatch(ctx, match)

	return match, err
}
