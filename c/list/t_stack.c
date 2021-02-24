#include <stdio.h>

#include "stack.h"

#define LEN 20

int main(void){

    int arr[LEN];
    int i;

    Stack *stack = malloc(sizeof(Stack));
    if(stack == NULL)
        return -1;

    stack_init(stack, NULL);

    for (i=0; i< LEN; i++){
        arr[i] = i+LEN;
        stack_push(stack, &arr[i]);
    }

    printS(stack);

    int *out;
    stack_pop(stack, (void **)&out);
    printf("pop out %d \n", *out);
    stack_pop(stack, (void **)&out);
    printf("pop out %d \n", *out);
    stack_pop(stack, (void **)&out);
    printf("pop out %d \n", *out);
    stack_pop(stack, (void **)&out);
    printf("pop out %d \n", *out);
    stack_pop(stack, (void **)&out);
    printf("pop out %d \n", *out);
    stack_pop(stack, (void **)&out);
    printf("pop out %d \n", *out);

    int in = 50;
    printf("push in %d \n", in);
    stack_push(stack, (void *)&in);

    int b = 52;
    printf("push in %d \n", b);
    stack_push(stack, (void *)&b);

    int c = 55;
    printf("push in %d \n", c);
    stack_push(stack, (void *)&c);

    printS(stack);

    stack_destory(stack);

    free(stack);

    return 0;
}
