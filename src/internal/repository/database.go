package repository

import (
	"context"
	"fmt"
	"log/slog"
	"price-tracker-service/src/domain"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type ClickHouseLogRepository struct {
	conn driver.Conn
}

func NewClickHouseLogRepository(cfg *domain.DatabaseConfig) (*ClickHouseLogRepository, error) {
	ctx := context.Background()

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.Database,
			Username: cfg.User,
			Password: cfg.Password,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("open clickhouse: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			slog.Error("clickhouse ping exception",
				"code", exception.Code,
				"message", exception.Message,
			)
		}
		return nil, fmt.Errorf("ping clickhouse: %w", err)
	}

	if err := ensureTable(ctx, conn); err != nil {
		return nil, err
	}

	return &ClickHouseLogRepository{conn: conn}, nil
}

func ensureTable(ctx context.Context, conn driver.Conn) error {
	var exists uint8
	err := conn.QueryRow(ctx, "EXISTS TABLE price_changes_log").Scan(&exists)
	if err != nil {
		return fmt.Errorf("check table exists: %w", err)
	}

	if exists == 1 {
		return nil
	}

	const ddl = `
	CREATE TABLE price_changes_log
	(
		id         UUID,
		client 	   String,
		pairname   String,
		value      Float64,
		createDate DateTime
	)
	ENGINE = MergeTree
	ORDER BY (createDate, id)`

	if err := conn.Exec(ctx, ddl); err != nil {
		return fmt.Errorf("create table price_changes_log: %w", err)
	}

	slog.Info("table price_changes_log created")
	return nil
}

func (c *ClickHouseLogRepository) InsertLog(ctx context.Context, log *domain.PriceChangesLog) error {
	batch, err := c.conn.PrepareBatch(ctx, "INSERT INTO price_changes_log")
	if err != nil {
		return fmt.Errorf("prepare batch: %w", err)
	}

	uuid, _ := uuid.NewUUID()
	if err := batch.Append(uuid, log.PairName, log.Value, log.CreateDate); err != nil {
		return fmt.Errorf("append to batch: %w", err)
	}

	return batch.Send()
}
