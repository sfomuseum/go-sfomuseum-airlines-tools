cli:
	go build -mod vendor -o bin/lookup cmd/lookup/main.go
	go build -mod vendor -o bin/build-sfomuseum-data cmd/build-sfomuseum-data/main.go
	go build -mod vendor -o bin/build-flysfo-data cmd/build-flysfo-data/main.go

rebuild:
	go build -mod vendor -o bin/build-flysfo-data cmd/build-flysfo-data/main.go
	go build -mod vendor -o bin/build-sfomuseum-data cmd/build-sfomuseum-data/main.go
	bin/build-flysfo-data
	bin/build-sfomuseum-data
	go build -mod vendor -o bin/lookup cmd/lookup/main.go

test-flysfo:
	./bin/lookup -source flysfo B6 DI EI BF HA IG JL LH MH NZ OZ QF SK SN SQ AA AV DL NH AM HX KE FJ PR AY LX CA SU AZ UX CZ AF KL RJ KE CX TG SE
