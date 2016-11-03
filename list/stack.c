#include <stdlib.h>

#include "list.h"
#include "stack.h"

/* stack_push */
int stack_push(Stack *stack, const void *data){
    return list_ins_next(stack, NULL, data);
}

int stack_pop(Stack *stack, void **data){
    return list_rem_next(stack, NULL, data);
}
