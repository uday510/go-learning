APP_NAME = card
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%H:%M:%S_UTC')
GIT_COMMIT = $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
GO_VERSION = $(shell go version | awk '{print $$3}')

LDFLAGS = -X=main.buildTime=$(BUILD_TIME) \
          -X=main.gitCommit=$(GIT_COMMIT) \
          -X=main.goVersion=$(GO_VERSION)

.PHONY: all build run clean info

all: build run

build:
	@echo "Building..."
	@go build -ldflags "$(LDFLAGS)" -o $(APP_NAME)

run:
	@echo "Running..."
	@./$(APP_NAME)

clean:
	@rm -f $(APP_NAME)

info:
	@echo "Build Time:  $(BUILD_TIME)"
	@echo "Git Commit:  $(GIT_COMMIT)"
	@echo "Go Version:  $(GO_VERSION)"