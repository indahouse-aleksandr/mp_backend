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

help: Makefile
	@echo "Choose a command:"
	@echo " migrate-install"
	@echo " migrate-create"	
	@echo " qlc-install"
	@echo " go-vendor" 

