OUT_DIR:="bin"
APP_NAME:="gotailwind"

all: clean build

tailwind:
	@echo "Running TailwindCSS build..."
	@tailwindcss -i tailwindbase.css -o server/static/main.css

tailwind-live:
	@echo "Starting tailwind-live..."
	@tailwindcss -i tailwindbase.css -o server/static/main.css --watch

clean:
	@echo "Cleaning up..."
	@gofumpt -l -w .
	@go mod tidy -v

build: tailwind
	@echo "Building..."
	@go build -o $(OUT_DIR)/$(APP_NAME)
