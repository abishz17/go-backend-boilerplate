migrateup:
		migrate -path ./migrations -database "postgresql://postgres:password@localhost:5432/godb?sslmode=disable" -verbose up

migratedown:
		migrate -path ./migrations -database "postgresql://postgres:password@localhost:5432/godb?sslmode=disable" -verbose down