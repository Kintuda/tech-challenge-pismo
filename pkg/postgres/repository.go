package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repositories interface {
	AccountRepositoryPg | TransactionRepositoryPg
}

type PostgresRepository struct {
	Connection *Pool
	Trx        *pgx.Tx
}

type PgRepository struct {
	Executor QueryExecutor
}

type QueryExecutor interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

func (p *PgRepository) GetExecutor(ctx context.Context) QueryExecutor {
	tx := ctx.Value("tx")

	if tx != nil {
		return tx.(pgx.Tx)
	}

	return p.Executor
}

func NewPostgresRepository(conn *Pool) *PostgresRepository {
	return &PostgresRepository{
		Connection: conn,
		Trx:        nil,
	}
}

func NewRepository[K Repositories](p *PostgresRepository) *K {
	if p.Trx == nil {
		return &K{
			PgRepository: PgRepository{
				Executor: p.Connection.Conn,
			},
		}
	}

	return &K{
		PgRepository: PgRepository{
			Executor: *p.Trx,
		},
	}
}
