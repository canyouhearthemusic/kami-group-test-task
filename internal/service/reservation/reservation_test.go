package reservation

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/canyouhearthemusic/kamigr/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateReservation_NoOverlap(t *testing.T) {
	ctx := context.Background()

	s := setupTestService(ctx, t)

	roomID := "123A"
	startTime1 := time.Now().Add(1 * time.Hour)
	endTime1 := startTime1.Add(2 * time.Hour)
	startTime2 := endTime1.Add(1 * time.Hour)
	endTime2 := startTime2.Add(2 * time.Hour)

	rsrv, _ := s.CreateReservation(ctx, sqlc.CreateReservationParams{
		RoomID:    roomID,
		StartTime: pgtype.Timestamptz{Time: startTime1},
		EndTime:   pgtype.Timestamptz{Time: endTime1},
	})
	require.NotNil(t, rsrv)

	rsrv2, _ := s.CreateReservation(ctx, sqlc.CreateReservationParams{
		RoomID:    roomID,
		StartTime: pgtype.Timestamptz{Time: startTime2},
		EndTime:   pgtype.Timestamptz{Time: endTime2},
	})
	require.NotNil(t, rsrv2)
}

func TestCreateReservation_Overlap(t *testing.T) {
	ctx := context.Background()

	s := setupTestService(ctx, t)

	roomID := "test-room-2"
	startTime := time.Now().Add(1 * time.Hour)
	endTime := startTime.Add(2 * time.Hour)

	_, err := s.CreateReservation(ctx, sqlc.CreateReservationParams{
		RoomID:    roomID,
		StartTime: pgtype.Timestamptz{Time: startTime},
		EndTime:   pgtype.Timestamptz{Time: endTime},
	})
	require.NoError(t, err)

	_, err = s.CreateReservation(ctx, sqlc.CreateReservationParams{
		RoomID:    roomID,
		StartTime: pgtype.Timestamptz{Time: startTime.Add(30 * time.Minute)},
		EndTime:   pgtype.Timestamptz{Time: endTime.Add(30 * time.Minute)},
	})
	require.Error(t, err)
}

func TestConcurrentReservations(t *testing.T) {
	ctx := context.Background()

	s := setupTestService(ctx, t)

	roomID := "test-room-3"
	startTime := time.Now().Add(1 * time.Hour)
	endTime := startTime.Add(2 * time.Hour)

	concurrentRequests := 5
	errs := make(chan error, concurrentRequests)

	for i := 0; i < concurrentRequests; i++ {
		go func() {
			_, err := s.CreateReservation(ctx, sqlc.CreateReservationParams{
				RoomID:    roomID,
				StartTime: pgtype.Timestamptz{Time: startTime},
				EndTime:   pgtype.Timestamptz{Time: endTime},
			})
			errs <- err
		}()
	}

	for i := 0; i < concurrentRequests; i++ {
		err := <-errs
		if i == 0 {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}

func setupTestService(ctx context.Context, t *testing.T) *Service {
	// conf, err := config.MustLoad(ctx)
	// require.NoError(t, err)

	// dbconf := conf.Database

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"al1bek", "", "localhost", "5432", "kami-db")
	// dbconf.User, dbconf.Password, dbconf.Host, dbconf.Port, dbconf.Name)

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to establish connection to db: %s\n", err)
	}
	require.NoError(t, err)

	dao := sqlc.New(conn)

	t.Cleanup(func() {
		conn.Close(ctx)
	})

	s := New(dao)

	return s
}
