package handler

import (
	"net/http"

	"github.com/Max-Gabriel-Susman/bestir-matchmaking-service/internal/foundation/database"
	"github.com/Max-Gabriel-Susman/bestir-matchmaking-service/internal/foundation/web"
	"github.com/Max-Gabriel-Susman/bestir-matchmaking-service/internal/match"
)

var _ http.Handler = (*web.App)(nil)

// maybe we'll add gitsha and other params later
func API(d Deps) *web.App {
	app := web.NewApp()
	dbrConn := database.NewDBR(d.DB)
	matchAPI := match.NewAPI(match.NewMySQLStore(dbrConn))
	matchEndpoints(app, matchAPI)
	return app
}
