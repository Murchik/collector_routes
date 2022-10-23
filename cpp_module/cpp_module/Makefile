CP ?= cp
MV ?= mv
RM ?= rm

CCX := g++
CCFLAGS := -Wall -fPIC
LDFLAGS :=

all: libAlgorithms.so
	$(CP) libAlgorithms.so ../../

Algorithms.o: Algorithms.cpp Algorithms.hpp
	$(CCX) $(CCFLAGS) -c $<

libAlgorithms.so: Algorithms.o
	$(CCX) $(LDFLAGS) -shared -o $@ $<

clean:
	$(RM) Algorithms.o libAlgorithms.so Algorithms.o ../../libAlgorithms.so
