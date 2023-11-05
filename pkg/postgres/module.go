package postgres

import (
	"github.com/kintuda/tech-challenge-pismo/pkg/account"
	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewConnectionPool,
		NewPostgresRepository,
		fx.Annotate(NewRepository[AccountRepositoryPg], fx.As(new(account.AccountRepository))),
		fx.Annotate(NewRepository[TransactionRepositoryPg], fx.As(new(transaction.TransactionRepository))),
	),
)
