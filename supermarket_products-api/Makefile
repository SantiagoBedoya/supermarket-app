dev:
	GIN_MODE=debug \
  PORT=3000 \
	MYSQL_URI="root:root@tcp(127.0.0.1:3306)/supermarket_products?charset=utf8mb4&parseTime=True&loc=Local" \
	ACCOUNT_API_URL=http://localhost:3001/api/v1 \
	go run main.go