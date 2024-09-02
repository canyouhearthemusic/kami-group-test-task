package dto

import (
	"errors"
	"time"

	"github.com/canyouhearthemusic/kamigr/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type ReservationDTO struct {
	RoomID    string `json:"room_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (dto *ReservationDTO) Validate() error {
	if dto.RoomID == "" {
		return errors.New("room_id cannot be empty")
	}

	if _, err := time.Parse(time.DateTime, dto.StartTime); err != nil {
		return errors.New("start_time must be a valid datetime in DateTime format")
	}

	if _, err := time.Parse(time.DateTime, dto.EndTime); err != nil {
		return errors.New("end_time must be a valid datetime in DateTime format")
	}

	startTime, _ := time.Parse(time.DateTime, dto.StartTime)
	endTime, _ := time.Parse(time.DateTime, dto.EndTime)
	if !startTime.Before(endTime) {
		return errors.New("start_time must be before end_time")
	}

	return nil
}

func (dto *ReservationDTO) ConvertToSQLC() (sqlc.CreateReservationParams, error) {
	var stp pgtype.Timestamptz
	var etp pgtype.Timestamptz

	st, err := time.Parse(time.DateTime, dto.StartTime)
	if err != nil {
		return sqlc.CreateReservationParams{}, errors.New("invalid start_time format")
	}
	stp.Time = st
	stp.Valid = true

	et, err := time.Parse(time.DateTime, dto.EndTime)
	if err != nil {
		return sqlc.CreateReservationParams{}, errors.New("invalid end_time format")
	}
	etp.Time = et
	etp.Valid = true

	return sqlc.CreateReservationParams{
		RoomID:    dto.RoomID,
		StartTime: stp,
		EndTime:   etp,
	}, nil
}
