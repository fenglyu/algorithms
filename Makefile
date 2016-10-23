CC=gcc
CFLAGS=-I. -Wall -g 
DEPS = list.h
OBJ= list.o t_list.o

#t_list: t_list.o list.o
#	$(CC) -o t_list t_list.o list.o

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

t_list: $(OBJ)
	$(CC) -o $@ $^ $(CFLAGS)
