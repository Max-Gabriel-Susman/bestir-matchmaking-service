package match

import (
	"context"

	"github.com/Max-Gabriel-Susman/bestir-matchmaking-service/internal/foundation/database"
	"github.com/gocraft/dbr/v2"
)

func NewMySQLStore(conn *dbr.Connection) *MySQLStorage {
	return &MySQLStorage{conn: conn, sess: conn.NewSession(nil)}
}

type MySQLStorage struct {
	conn *dbr.Connection
	sess *dbr.Session
}

var (
	matchTable = database.NewTable("match", Match{})
)

func (s *MySQLStorage) Listmatchs(ctx context.Context) ([]Match, error) {
	query := s.sess.Select(matchTable.Columns...).
		From(matchTable.Name)

	matchs := []Match{}

	if _, err := query.LoadContext(ctx, &matchs); err != nil {
		return matchs, database.ClassifyError(err)
	}

	return matchs, nil
}

func (s *MySQLStorage) getMatchByIdempotencyKey(ctx context.Context, idempotencyKey string) (Match, error) {
	var match Match
	err := s.sess.Select(matchTable.Columns...).
		From(matchTable.Name).
		Where("idempotency_key = ?", idempotencyKey).
		LoadOneContext(ctx, &match)
	return match, database.ClassifyError(err)
}

func (s *MySQLStorage) CreateMatch(ctx context.Context, match Match) error {
	_, err := s.sess.InsertInto(matchTable.Name).
		Columns(matchTable.Columns...).
		Record(match).
		ExecContext(ctx)
	return database.ClassifyError(err)
}

func (s *MySQLStorage) DeleteMatch(ctx context.Context, match Match) error {
	_, err := s.sess.InsertInto(matchTable.Name).
		Columns(matchTable.Columns...).
		Record(match).
		ExecContext(ctx)
	return database.ClassifyError(err)
}

func (s *MySQLStorage) UpdateMatch(ctx context.Context, match Match) error {
	_, err := s.sess.InsertInto(matchTable.Name).
		Columns(matchTable.Columns...).
		Record(match).
		ExecContext(ctx)
	return database.ClassifyError(err)
}
