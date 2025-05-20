env=8081
toml=8080

.phony: build
build: 
	go build -v ./cmd/api
run:
	go run ./cmd/api/.
killenv: 
	@echo "kill env at env"
	@fuser -k ${env}/tcp 
	@lsof -i tcp:$(env) || echo "No process found on port $(env)"
# killtoml:
# 	@echo "kill toml at env"
# 	@fuser -k ${toml}/tcp 
# 	@lsof -i :$8080
# 	@lsof -i tcp:$(toml) || echo "No process found on port $(toml)"
restart: run killenv 