.PHONY: all
all: clean build run

.PHONY: clean
clean:
	@rm -rf ./mik8s.app ./ui/build
	@echo "🗑 Clean complete!"

.PHONY: build
build:
	@cd ./ui && npm install
	@cd ./ui && npm run build
	@mkdir -p ./mik8s.app/Contents/MacOS
	@go build -o ./mik8s.app/Contents/MacOs/mik8s
	@echo "✅ Build complete!"

.PHONY: run
run:
	@open ./mik8s.app
	@echo "🚀 mik8s running!"