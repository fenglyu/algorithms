CC=gcc
CFLAGS=-I. -Wall -g 
DEPS = list.h
OBJ= list.o t_list.o

DLIST_SRC  = dlist.c t_dlist.c
DLIST_DEPS = dlist.h
DLIST_OBJ  = dlist.o t_dlist.o
CLIST_OBJ = clist.o t_clist.o


#t_list: t_list.o list.o
#	$(CC) -o t_list t_list.o list.o

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

t_list: $(OBJ)
	$(CC) -o $@ $^ $(CFLAGS)


dlist.o : dlist.c dlist.h

t_dlist.o : t_dlist.c dlist.h

t_dlist: $(DLIST_OBJ)
	$(CC) -o $@ $^ $(CFLAGS)
# valgrind --leak-check=full --show-leak-kinds=all ./t_dlist
#
#
t_clist.o : t_clist.c clist.h

t_clist: $(CLIST_OBJ)
	$(CC) -o $@ $^ $(CFLAGS)

clean:
	rm -f $(DLIST_OBJ) $(OBJ) t_list t_dlist
