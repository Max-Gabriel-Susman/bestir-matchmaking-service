package match

import "github.com/google/uuid"

/*
 applicaiton should capture all the information needed to provision and
 manage an match on the bestir network
*/
type Match struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

type IncomingMatch struct {
	Name string `json:"name" required:"true"`
	// IdempotencyKey null.String `json:"-" db:"idempotency_key"`
}

type matchGroup struct {
}
