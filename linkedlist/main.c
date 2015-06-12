#include<stdio.h>

#include "singleList.h"

void print_elem(list p){
	printf("%d\n", p->elem);
}

int main(void){
	list p;

	insert(make_node(10));
	insert(make_node(4));
	insert(make_node(3));

	p = search(4);
	delete(p);
	free_node(p);
	traverse(print_elem);
	destroy();
	push(make_node(100));
	push(make_node(200));
	push(make_node(250));
	push(make_node(50));
	push(make_node(17));

	reverse(p);

	while(p = pop()){
		print_elem(p);
		free_node(p);
	}

	return 0 ;
}
