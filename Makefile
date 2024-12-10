
.PHONY: build
build:
	@go build -o build/rplanner -v ./backend/
	@cd frontend; npm run build

.PHONY: backend
backend:
	@go build -o build/rplanner -v ./backend/
