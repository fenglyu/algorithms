#include <stdio.h>

#include "queue.h"

#define LEN 20

int main(void){

    int arr[LEN];
    int i;

    Queue *queue = malloc(sizeof(Queue));
    if(queue == NULL)
        return -1;

    queue_init(queue, NULL);

    for (i=0; i< LEN; i++){
        arr[i] = i+LEN;
        queue_enqueue(queue, &arr[i]);
    }

    printQ(queue);

    int *out;
    queue_dequeue(queue, (void **)&out);
    printf("pop out %d \n", *out);
    queue_dequeue(queue, (void **)&out);
    printf("pop out %d \n", *out);
    queue_dequeue(queue, (void **)&out);
    printf("pop out %d \n", *out);
    queue_dequeue(queue, (void **)&out);
    printf("pop out %d \n", *out);
    queue_dequeue(queue, (void **)&out);
    printf("pop out %d \n", *out);
    queue_dequeue(queue, (void **)&out);
    printf("pop out %d \n", *out);

    int in = 50;
    printf("push in %d \n", in);
    queue_enqueue(queue, (void *)&in);

    int b = 52;
    printf("push in %d \n", b);
    queue_enqueue(queue, (void *)&b);

    int c = 55;
    printf("push in %d \n", c);
    queue_enqueue(queue, (void *)&c);

    printQ(queue);

    queue_destory(queue);

    free(queue);

    return 0;
}
