run-mysql:
	docker run --name mysql8 -p 3306:3306/tcp -e MYSQL_ROOT_PASSWORD=1234 -d mysql:8

sh-mysql:
	docker exec -it mysql8 /bin/sh

## Docker Command
## docker exec comamnd to run command in container
## mysql.schema_migrations 테이블안에서 관리
db-up:
	migrate -source file:./db/migration -database "mysql://root:1234@tcp(localhost:3306)/simple_bank" --verbose up

db-down:
	migrate -source file:./db/migration -database "mysql://root:1234@tcp(localhost:3306)/simple_bank" --verbose down
	
