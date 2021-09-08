module github.com/ozoncp/ocp-timeline-api

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Masterminds/squirrel v1.5.0
	github.com/Shopify/sarama v1.29.1
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.3 // indirect
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.24.0
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a // indirect
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
)

replace github.com/ozoncp/ocp-timeline-api/pkg/ocp-timeline-api => ./pkg/ocp-timeline-api
