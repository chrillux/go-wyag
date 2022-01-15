.PHONY: build clean cat-file hash-object log

clean:
	rm wyag

build:
	go build -o wyag main.go

cat-file:
	go run main.go cat-file -t 028de97e4bde13ca950e33ac5c485001578e9629

hash-object:
	go run main.go hash-object main.go

log:
	go run main.go log 1b9d56f40d94dba86bce7d91eb09db5de65223a4
