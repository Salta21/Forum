run:
	go run ./cmd/main.go

build:
	docker image build -t forum .

d-run:
	docker container run -p 8080:8080 --detach --name forumapp forum