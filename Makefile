FILES = main.go 1.go 2.go 3.go 4.go 5.go 6.go 7.go 8.go 9.go 10.go 11.go 12.go 13.go 14.go 15.go 16.go 17.go 18.go 19.go 20.go 21.go 22.go 23.go 24.go 25.go 26.go

all: run

run:
	go run $(FILES)

build:
	go build $(FILES)