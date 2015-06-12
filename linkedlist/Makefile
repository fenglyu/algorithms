CFLAGS = -g -Wall
CC = gcc

main: main.o singleList.o
	gcc main.o singleList.o $(CFLAGS)

main.o: main.c singleList.h
	gcc -c main.c $(CFLAGS)

singleList.o: singleList.c singleList.h
	gcc -c singleList.c $(CFLAGS)
