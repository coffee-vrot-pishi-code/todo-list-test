migrate create -ext sql -dir ./schema -seq init

migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
