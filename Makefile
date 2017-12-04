NAME := testgo
PKG := github.com/bitvector2/$(NAME)

.PHONY: clean docker-clean build docker-build test docker-test push docker-push

all: build

clean:
	rm $(NAME)

docker-clean:
	docker rm $(shell docker ps -aq) || true
	docker rmi $(shell docker images -aq) || true

build:
	go build

docker-build:
	docker build -t $(NAME):latest .

test: build
	./$(NAME)

docker-test: docker-build
	docker run -i -t $(NAME):latest

push: test
	git add vendor
	git commit -am "updated $(NAME) sources" || true
	git push

docker-push: docker-test
	docker tag $(NAME):latest quay.cnqr.delivery/containerhosting/$(NAME):latest
	docker push quay.cnqr.delivery/containerhosting/$(NAME):latest
