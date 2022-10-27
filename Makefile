CP ?= cp
MV ?= mv
RM ?= rm

all: run

run: server.go
	go run server.go

libAlgorithms.so:
	cd cpp_module/cpp_module && $(MAKE)

build: libAlgorithms.so
	go build -o server.out server.go

clean:
	cd cpp_module/cpp_module && $(MAKE) clean
	$(RM) server.out *.json *.xml
