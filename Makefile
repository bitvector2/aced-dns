NAME := testgo
PKG := github.com/bitvector2/$(NAME)
FILES := $(shell find . -type f -name "*.go" | egrep -v "./vendor")

.PHONY: clean docker-clean prep build docker-build test docker-test push docker-push

all: build

clean:
	rm $(NAME)

docker-clean:
	docker rm $(shell docker ps -aq) || true
	docker rmi $(shell docker images -aq) || true

prep:
	go fix $(PKG)
	go vet $(PKG)
	goimports -w $(FILES)
	golint ${PKG}
	@echo "** Code prepared"

build: prep
	go build

docker-build:
	docker build -t $(NAME):latest .

run: build
	./$(NAME) -logtostderr -kubeconfig $(HOME)/.kube/config -outputdir /tmp

test: build
	./$(NAME) -h || true

docker-test: docker-build
	docker run -i -t $(NAME):latest

push: test
	git add vendor
	git commit -am "updated $(NAME) sources" || true
	git push

docker-push: docker-test
	docker tag $(NAME):latest quay.cnqr.delivery/containerhosting/$(NAME):latest
	docker push quay.cnqr.delivery/containerhosting/$(NAME):latest
