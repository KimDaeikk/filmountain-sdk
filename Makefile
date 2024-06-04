build:
	@go build -o ./bin/ -ldflags="all=-extldflags=-Wl,--allow-multiple-definition"

run:
	@./bin/filmountain-sdk miner-contract add t0118000

start:
	@$(MAKE) -s build
	@$(MAKE) -s run