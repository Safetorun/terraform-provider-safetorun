build:
	cd safetorun && go build
	cd web/src && GOOS=js GOARCH=wasm go build -o  ../json.wasm
	go build

test:
	cd web/src && go test

tidy:
	cd safetorun && go mod tidy
	cd web/src && go mod tidy
	cd tf-safetorun && go mod tidy
	go mod tidy
update:
	cd safetorun && go get -u && go mod tidy
	cd web/src && go get -u && go mod tidy
	cd tf-safetorun && go get -u && go mod tidy
	go get -u && go mod tidy