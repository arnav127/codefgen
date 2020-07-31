all:
	go build -o codefgen
	./install.sh
clean:
	rm codefgen