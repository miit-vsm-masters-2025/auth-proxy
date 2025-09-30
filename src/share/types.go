package share

import (
	"database/sql"

	"github.com/valkey-io/valkey-go"
	"go.uber.org/zap"
)

type AppContext struct {
	Valkey         *valkey.Client
	Logger         *zap.SugaredLogger
	PostgresClient *sql.DB
}
