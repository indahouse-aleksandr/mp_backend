migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
	cp /go/bin/migrate /var/www/app/bin/

sqlc-install:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.21.0
	cp /go/bin/sqlc /var/www/app/bin/

go-vendor:
	cd ../
	go mod vendor

migrate-create:
	../bin/migrate create -ext="sql" -dir="../migrations" -seq -digits="5" new

migrate-up-all:
	../bin/migrate -path="../migrations" -database="postgresql://postgres:postgres@mp_postgres:5432/store?sslmode=disable" -verbose up

migrate-down-all:
	../bin/migrate -path="../migrations" -database="postgresql://postgres:postgres@mp_postgres:5432/store?sslmode=disable" -verbose down

sqlc-generate:
	../bin/sqlc --file "../dev_local/sqlc/sqlc.yaml" generate


help: Makefile
	@echo "Choose a command:"
	@echo " migrate-install"
	@echo " migrate-create"
	@echo " migrate-up-all"
	@echo " qlc-install"
	@echo " go-vendor" 
	@echo " migrate-down-all" 
	@echo " sqlc-generate" 
