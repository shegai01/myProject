port=?8080
run:
	go run ./cmd/api/.
kill: 
	@echo "kill port at ${[port]}"
	@fuser -k ${port}/tcp 