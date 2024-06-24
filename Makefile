build:
	@go build -o ./bin/ -ldflags="all=-extldflags=-Wl,--allow-multiple-definition"
.PHONY: build

run:
	@./bin/filmountain-sdk miner-vault add t0118000

config:
	mkdir -p ~/.filmountain
	cp config.toml ~/.filmountain/setting.yaml

install:
	cp ./bin/filmountain-sdk /usr/local/bin/filmountain-sdk
.PHONY: install

start:
	@$(MAKE) -s build
	@$(MAKE) -s run