docker-template = ./templates/psql-docker.yml

init:
	go mod init github.com/luizarnoldch/mesapara2_backend
update:
	go mod tidy
start:
	go run main.go
db-up:
	docker-compose -f ${docker-template} up -d db_backend
db-down:
	docker-compose -f ${docker-template} down --volumes --rmi 'local' --remove-orphans
	docker image prune -af
db-stop:
	docker-compose -f ${docker-template} stop -d db_backend
db-start:
	docker-compose -f ${docker-template} start -d db_backend
