# Bump these on release
VERSION_MAJOR ?= 1
VERSION_MINOR ?= 0
VERSION_BUILD ?= 0
VERSION_RC ?= ""

RAW_VERSION=$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)
VERSION ?= $(RAW_VERSION)$(VERSION_RC)

DOCKER_OWNER ?= innovia
DOCKER_REPO=$(DOCKER_OWNER)/secrets-consumer-webhook

# Get git commit id
COMMIT_NO := $(shell git rev-parse HEAD 2> /dev/null || true)
COMMIT ?= $(if $(shell git status --porcelain --untracked-files=no),"${COMMIT_NO}-dirty","${COMMIT_NO}")
CURRENT_GIT_BRANCH ?= $(shell git branch | grep \* | cut -d ' ' -f2)

BUILD_DIR ?= ./out
$(shell mkdir -p $(BUILD_DIR))

OSARCH := "linux/amd64 linux/386 windows/amd64 windows/386 darwin/amd64 darwin/386"

# Set the version and commit
SECRETS_CONSUMER_ENV_LDFLAGS := -X github.com/innovia/secrets-consumer-env/pkg/version.version=$(VERSION) -X github.com/innovia/secrets-consumer-env/pkg/version.gitCommitID=$(COMMIT)

.PHONY: cross
cross:
	gox -osarch=$(OSARCH) -output "out/secrets-consumer-env-{{.OS}}-{{.Arch}}" -ldflags="$(SECRETS_CONSUMER_ENV_LDFLAGS)"

docker-build:
	docker build -t innovia/secrets-consumer-env:$(VERSION) . --build-arg VERSION=$(VERSION) --build-arg COMMIT=$(COMMIT)

docker-push:
	docker push innovia/secrets-consumer-env:$(VERSION)

up: docker-build docker-push

publish-latest:  ## Publish the `latest` tagged container
	@echo publish latest to $(DOCKER_REPO)
	docker tag $(DOCKER_OWNER)/secrets-consumer-webhook:$(VERSION) $(DOCKER_OWNER)/secrets-consumer-webhook:latest
	docker push $(DOCKER_REPO):latest


.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: vet
vet: ## Run go vet
	@go vet ./...


