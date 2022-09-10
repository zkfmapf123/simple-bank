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
	
run:
	go run ./src/main.go

unit-test:
	go test ./test/unit

e2e-test:
	go test ./test/e2e

e2e-bash:
	sh ./e2e_bash_profile.sh


## API Curl

api-base:
	curl -X GET "http://localhost:8080/"

api-account:
	curl -d '{"owner":"leedonggyu", "currency":"KR", "timezone":"Asia/Seoul"}' \
	-X POST "http://localhost:8080/account" \

api-get:
	curl -X GET "http://localhost:8080/account/32"

api-list:
	curl -X GET "http://localhost:8080/account?offset=0&limit=10"