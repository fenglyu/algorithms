CC=gcc
CFLAGS=-I. -Wall -g 
DEPS = list.h
OBJ= list.o t_list.o

DLIST_SRC  = dlist.c t_dlist.c
DLIST_DEPS = dlist.h
DLIST_OBJ  = dlist.o t_dlist.o
CLIST_OBJ = clist.o t_clist.o
STACK_OBJ = stack.o list.o t_stack.o

#t_list: t_list.o list.o
#	$(CC) -o t_list t_list.o list.o

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

t_list: $(OBJ)
	$(CC) -o $@ $^ $(CFLAGS)

list.o : list.c list.h

dlist.o : dlist.c dlist.h

t_dlist.o : t_dlist.c dlist.h

t_dlist: $(DLIST_OBJ)
	$(CC) -o $@ $^ $(CFLAGS)
# valgrind --leak-check=full --show-leak-kinds=all ./t_dlist
#

t_clist.o : t_clist.c clist.h

t_clist: $(CLIST_OBJ)
	$(CC) -o $@ $^ $(CFLAGS)

stack.o : stack.c stack.h list.h

t_stack.o: t_stack.c stack.h

t_stack: $(STACK_OBJ)
	$(CC) -o $@ $^ $(CFLAGS)

queue.o : queue.c queue.h list.h

t_queue.o : t_queue.c queue.h

t_queue: queue.o  list.o t_queue.o
	$(CC) -o $@ $^ $(CFLAGS)

clean:
	rm -f *.o t_list t_dlist t_stack t_queue
