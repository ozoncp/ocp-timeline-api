package db

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

var DSN = "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable"
var dbKey = "dbKeyForContext"

func Connect() *sqlx.DB {
	db, err := sqlx.Open("pgx", DSN)

	if err != nil {
		log.Fatalf("failled to connect db %v", err)
	}

	return db
}

func NewContext(ctx context.Context, db *sqlx.DB) context.Context {
	ctxDb := context.WithValue(ctx, &dbKey, db)

	return ctxDb
}

/*func NewInterceptorWithDB(db *sqlx.DB) grpc.UnaryServerInterceptor {
	return func (ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
		return handler(NewContext(ctx, db), req)
	}
}*/
