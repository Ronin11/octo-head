include .env

GO=go

APP_NAME=octo-mantle

default: run

# Create a build go image
build:
	@echo "\n~~~~~~~ BUILDING ~~~~~~~"
	$(GO) build -o $(APP_NAME) .

# Remove built go image
clean:
	@echo "\n~~~~~~~ CLEANING ~~~~~~~"
	rm -f $(APP_NAME)

run: build
	@echo "\n~~~~~~~ RUNNING ~~~~~~~"
	export SERVER=$(NATS_SERVER); \
	./$(APP_NAME)

test:
	@echo "~~~~~~~ TESTING ~~~~~~~"
	go test



