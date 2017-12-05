



all:	clean build test1

git:
	-@git add .
	-@git commit -m "autocommit"

clean:
	-@rm -v main
build:
	go build main.go
test1:
	@echo "Test 1 -- no args"
	-@./main

test2:
	@echo "Test 2 -- correct args"
	./main 2 15 "0:12:34" c42007e018

test3:
	@echo "Test 2 -- Wrong args"
	./main 2 BB "32:0:12:34" 

test:
	go test -v
