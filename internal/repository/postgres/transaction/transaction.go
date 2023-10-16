package transaction

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type QueryExecutor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func Model(ctx context.Context, db QueryExecutor) QueryExecutor {
	tx := extractTx(ctx)
	if tx != nil {
		return tx
	}
	return db
}

type txKey struct{}

func injectTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func extractTx(ctx context.Context) pgx.Tx {
	if tx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		return tx
	}
	return nil
}

func (r *Repository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	tx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	if err = tFunc(injectTx(ctx, tx)); err != nil {
		slog.Error(err.Error())
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
