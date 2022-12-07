RM ?= rm

all: run

run: server.go
	go run server.go

clean:
	$(RM) *.json *.xml
