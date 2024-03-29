#ifndef STACK_H

#include <stdlib.h>
#include "list.h"


typedef List Stack;

#define stack_init list_init

#define stack_destory list_destory

#define printS printL

int stack_push(Stack *stack, const void *data);

int stack_pop(Stack *stack, void **data);

#define stack_peek(stack) ((stack)->head == NULL ? NULL : (stack)->head->data)

#define stack_size list_size

#endif
