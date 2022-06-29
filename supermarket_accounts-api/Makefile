dev:
	GIN_MODE=debug \
	PORT=3001 \
	MYSQL_URI="root:root@tcp(127.0.0.1:3306)/supermarket_accounts?charset=utf8mb4&parseTime=True&loc=Local" \
	JWT_SECRET=6014n6_5u93rm4k37 \
	go run main.go

seed:
	MYSQL_URI="root:root@tcp(127.0.0.1:3306)/supermarket_accounts?charset=utf8mb4&parseTime=True&loc=Local" \
	go run commands/accounts-seeder.go
