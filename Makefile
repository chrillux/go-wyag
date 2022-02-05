.PHONY: build clean cat-file hash-object log

build:
	go build -o wyag main.go

clean:
	rm wyag

cat-file:
	go run main.go cat-file -t 028de97e4bde13ca950e33ac5c485001578e9629

hash-object:
	go run main.go hash-object main.go

log:
	go run main.go log 1b9d56f40d94dba86bce7d91eb09db5de65223a4

ls-tree:
	go run main.go ls-tree 74db73dfc84b01ec0ddbdaaf18da87465e76b705

# checkout is not part of "make all" because it has the side effect of creating files on disk so
# it needs to be run separately.
checkout:
	go run main.go checkout 58792402d1a4afd6cd54c9d6709001564796d069 foo

all: cat-file hash-object log ls-tree test

test:
	go test ./...
