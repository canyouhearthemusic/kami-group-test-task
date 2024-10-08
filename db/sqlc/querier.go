// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateReservation(ctx context.Context, arg CreateReservationParams) (*Reservation, error)
	DeleteReservation(ctx context.Context, id int32) error
	GetAllReservations(ctx context.Context) ([]*Reservation, error)
	GetAllReservationsByRoomID(ctx context.Context, roomID string) ([]*Reservation, error)
	GetReservation(ctx context.Context, id int32) (*Reservation, error)
	UpdateReservation(ctx context.Context, arg UpdateReservationParams) (*Reservation, error)
}

var _ Querier = (*Queries)(nil)
