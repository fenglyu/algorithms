#ifndef LINKEDLIST_H
#define LINKEDLIST_H

typedef struct node *list;

struct node{
	unsigned char elem;
	list next;
};

list make_node(unsigned char elem);

void free_node(list p);

list search(unsigned char elem);

void insert(list p);

void delete(list p);
 
void traverse(void (*visit)(list));

void destroy(void);

void push(list p);

list pop(void);

list reverse(list head);

#endif
