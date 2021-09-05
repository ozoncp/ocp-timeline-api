module github.com/ozoncp/ocp-timeline-api

go 1.16

require (
	github.com/Masterminds/squirrel v1.5.0
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.3 // indirect
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.24.0
	github.com/shopspring/decimal v1.2.0 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
)

replace github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api => ./pkg/ocp-timeline-api
