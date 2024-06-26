module github.com/araxiaonline/endgame-item-generator

go 1.22.4

replace github.com/araxiaonline/endgame-item-generator/models => ../models

require github.com/go-sql-driver/mysql v1.8.1

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
