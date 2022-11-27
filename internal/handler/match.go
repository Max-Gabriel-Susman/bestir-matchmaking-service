package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Max-Gabriel-Susman/bestir-matchmaking-service/internal/foundation/web"
	"github.com/Max-Gabriel-Susman/bestir-matchmaking-service/internal/match"
	"github.com/go-chi/chi/v5"
)

type Match struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type matchGroup struct {
	*match.API
}

type ListmatchsResponse struct {
	matchs []match.Match `json:"matchs"`
}

func matchEndpoints(app *web.App, api *match.API) {
	ag := matchGroup{API: api}

	app.Handle("GET", "/match/{id}", ag.GetMatch)
	app.Handle("GET", "/match", ag.ListMatches)
	app.Handle("POST", "/match", ag.CreateMatch)
	app.Handle("DELETE", "/match/{id}", ag.CreateMatch)
	app.Handle("PUT", "/match/{id}", ag.CreateMatch)
}

func (ag matchGroup) ListMatches(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	matchs, err := ag.API.ListMatches(ctx)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, ListmatchsResponse{
		matchs: matchs,
	}, http.StatusOK)
}

// matchs := []match.match{
// 	{
// 		ID:      69,
// 		Balance: 69,
// 	},
// 	{
// 		ID:      420,
// 		Balance: 420,
// 	},
// }

func (ag matchGroup) CreateMatch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Create match a invoked")
	var input match.IncomingMatch
	if err := web.Decode(r.Body, &input); err != nil {
		return err
	}

	match, err := ag.API.CreateMatch(ctx, input)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, match, http.StatusCreated)
}

func (ag matchGroup) GetMatch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	matchID := chi.URLParam(r, "match_id")
	if matchID == "" {
		return nil
		// return handleMissingURLParameter(ctx, matchID, match)
	}

	return nil
}

func (ag matchGroup) UpdateMatch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Create match a invoked")
	var input match.IncomingMatch
	if err := web.Decode(r.Body, &input); err != nil {
		return err
	}

	match, err := ag.API.UpdateMatch(ctx, input)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, match, http.StatusCreated)
}

func (ag matchGroup) DeleteMatch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Create match a invoked")
	var input match.IncomingMatch
	if err := web.Decode(r.Body, &input); err != nil {
		return err
	}

	match, err := ag.API.DeleteMatch(ctx, input)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, match, http.StatusCreated)
}
