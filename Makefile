
NAME=volt

INSTALL_DIR=/usr/local/bin

all: clean build

build:
	@echo "🔨 building $(NAME).."
	go build -o $(NAME) main.go

install: build
	@echo "🔨 installing $(NAME).."
	cp $(NAME) $(INSTALL_DIR)

uninstall:
	@echo "🧹 uninstalling $(NAME).."
	rm -Rf $(NAME) $(INSTALL_DIR)

clean:
	@echo "🧹 cleaning up.."
	rm -Rf $(NAME)