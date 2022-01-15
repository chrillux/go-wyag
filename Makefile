.PHONY: build clean cat-file

clean:
	rm wyag

build:
	go build -o wyag main.go

cat-file:
	go run main.go cat-file -t 028de97e4bde13ca950e33ac5c485001578e9629

hash-object:
	go run main.go hash-object main.go
