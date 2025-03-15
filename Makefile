.PHONY: build dev


build-wasm:
	cd wasm && GOOS=js GOARCH=wasm go build -o ../client/public/main.wasm . && cd .. & \
	touch client/src/main.js

dev:
	(cd client && npm run dev && cd .. & \
	cd server && air && cd .. & \
	find wasm -name "*.templ" | entr -r templ generate & \
	find wasm -name "*.go" | entr -r make build-wasm & \
	wait) || (kill 0; exit 1)