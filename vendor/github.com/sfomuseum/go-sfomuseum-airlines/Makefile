vendor-deps:
	go mod vendor

fmt:
	go fmt *.go
	go fmt cmd/*.go
	go fmt flysfo/*.go
	go fmt sfomuseum/*.go
	go fmt wikipedia/*.go

tools: 	
	go build -o bin/lookup cmd/lookup/main.go

test-flysfo:
	@make tools
	./bin/lookup -source flysfo B6 DI EI BF HA IG JL LH MH NZ OZ QF SK SN SQ AA AV DL NH AM HX KE FJ PR AY LX CA SU AZ UX CZ AF KL RJ KE CX TG SE
