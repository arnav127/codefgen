all:
	go build -o codef
	./install.sh
clean:
	rm codef