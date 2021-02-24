IDIR=../list
CC=gcc
CFLAGS=-I$(IDIR) -Wall -g

DEPS=$(IDIR)/list.h
ODIR=obj
OBJ=reverse.o list.o
LIBS=-lm


list.o: ../list/list.c ../list/list.h
	$(CC) -c -o $@ $< $(CFLAGS)

reverse.o: reverse.c list.h
	$(CC) -c -o $@ $<  $(DEPS) $(CFLAGS)

reverse: reverse.o list.o
	$(CC) -o $@ $< $(DEPS) $(CFLAGS)


.PHONY: clean

clean:
	rm -f $(ODIR)/*.o *~ core*
