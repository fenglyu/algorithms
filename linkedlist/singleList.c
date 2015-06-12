#include<stdlib.h>

#include "singleList.h"

static list head = NULL;

void free_node(list p){ 
        free(p);
}

list make_node(unsigned char elem){
	list p = malloc(sizeof(*p));
	p->elem = elem;
	p->next=NULL;
	return p;
}

list search(unsigned char elem){
        list p;
        for(p = head; p ; p = p->next){
        	if(p->elem == elem)
			return p;
	}
	return NULL;
}

void insert(list p){ 
	p->next = head;
	head = p;	
}

void delete(list p){
	list t;
	if (p == head){
		head = p->next;
		return;
	}

	for(t = head; t ; t = t->next){
		if(t->next == p){
			t->next = p->next;
			return;
		}
	}
}


void traverse(void (*visit)(list)){
	list p;
	for(p = head; p ; p = p->next)
		visit(p);
}


void destroy(void){
	list q,p = head;
	head = NULL;

	while(p){
		q=p;
		p = p->next;
		free_node(q);
	}
}


void push(list p){
	insert(p);
}

list pop(void){
	if(head == NULL) return NULL;

	list p  = head;
	head = head->next;
	return p;
}

list reverse(list head){
	list q,p = head->next;
	if( head == NULL || p == NULL) return head;
	head->next = NULL;
	while(p){
		q = p->next;
		p->next = head;
		head = p;
		p = q;
	}
	return head;
}
