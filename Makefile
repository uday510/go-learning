APP_NAME = golang
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%H:%M:%S_UTC')
GIT_COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
GO_VERSION = $(shell go version | awk '{print $$3}')

LDFLAGS = -X=main.buildTime=$(BUILD_TIME) \
          -X=main.gitCommit=$(GIT_COMMIT) \
          -X=main.goVersion=$(GO_VERSION)

PLATFORMS = darwin/amd64 linux/amd64 windows/amd64

.PHONY: all build run clean info release watch test cross

all: build run

build:
	@echo "Building..."
	@go build -ldflags "$(LDFLAGS)" -o $(APP_NAME)

run:
	@echo "Running..."
	@./$(APP_NAME)

clean:
	@echo "Cleaning..."
	@rm -f $(APP_NAME)
	@rm -rf dist coverage.*

info:
	@echo "Build Time:  $(BUILD_TIME)"
	@echo "Git Commit:  $(GIT_COMMIT)"
	@echo "GitHub URI:  $$(git config --get remote.origin.url)"
	@echo "Go Version:  $(GO_VERSION)"	@echo "Go Version:  $(GO_VERSION)"

release:
	@mkdir -p dist
	@go build -ldflags "$(LDFLAGS)" -o dist/$(APP_NAME)-$(GIT_COMMIT)
	@echo "Release built: dist/$(APP_NAME)-$(GIT_COMMIT)"

watch:
	@echo "Watching for changes..."
	@reflex -r '\.go$$' -- sh -c 'clear && make'

test:
	@echo "Running tests..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out

cross:
	@echo "Cross-compiling..."
	@mkdir -p dist
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d/ -f1); \
		GOARCH=$$(echo $$platform | cut -d/ -f2); \
		OUTPUT=dist/$(APP_NAME)-$$GOOS-$$GOARCH; \
		[ "$$GOOS" = "windows" ] && OUTPUT=$$OUTPUT.exe; \
		echo "Building $$OUTPUT..."; \
		GOOS=$$GOOS GOARCH=$$GOARCH go build -ldflags "$(LDFLAGS)" -o $$OUTPUT; \
	done
	@echo "Cross-platform builds complete!"