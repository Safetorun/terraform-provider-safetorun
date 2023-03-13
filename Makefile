build:
	cd safetorun && go build
	cd web/src && GOOS=js GOARCH=wasm go build -o  ../json.wasm

test:
	cd web/src && go test