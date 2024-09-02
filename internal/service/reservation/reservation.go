package reservation

import (
	"context"
	"errors"

	"github.com/canyouhearthemusic/kamigr/db/sqlc"
)

func NewService(dao *sqlc.Queries) *Service {
	return &Service{dao: dao}
}

func (s *Service) CreateReservation(ctx context.Context, sqlcreq sqlc.CreateReservationParams) (*sqlc.Reservation, error) {
	existingReservations, err := s.GetAllReservationsByRoomID(ctx, sqlcreq.RoomID)
	if err != nil {
		return nil, err
	}

	if err := reservationExists(sqlcreq, existingReservations); err != nil {
		return nil, err
	}

	reservation, err := s.dao.CreateReservation(ctx, sqlcreq)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (s *Service) GetAllReservationsByRoomID(ctx context.Context, roomID string) ([]*sqlc.Reservation, error) {
	reservations, err := s.dao.GetAllReservationsByRoomID(ctx, roomID)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func reservationExists(curr sqlc.CreateReservationParams, reservations []*sqlc.Reservation) error {
	for _, res := range reservations {
		if (curr.StartTime.Time.Before(res.EndTime.Time) && curr.StartTime.Time.After(res.StartTime.Time)) ||
			(curr.EndTime.Time.After(res.StartTime.Time) && curr.EndTime.Time.Before(res.EndTime.Time)) ||
			(curr.StartTime.Time.Equal(res.StartTime.Time) && curr.EndTime.Time.Equal(res.EndTime.Time)) {
			return errors.New("reservation conflicts with an existing reservation")
		}
	}

	return nil
}
