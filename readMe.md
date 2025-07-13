BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S_UTC'); go build -ldflags "-X=main.buildTime=$BUILD_TIME" -o card

./card