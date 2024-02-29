NAME=volt

all: clean
	go build -o $(NAME) main.go

clean:
	rm -Rf $(NAME)