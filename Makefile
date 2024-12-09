
.PHONY: build
build:
	@env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o build/rplanner -v ./backend/
	@cd frontend; npm run build

.PHONY: backend
backend:
	@go build -o build/rplanner -v ./backend/
