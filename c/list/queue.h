#ifndef QUEUE_H
#define QUEUE_H

#include <stdlib.h>

#include "list.h"


typedef List Queue;

#define queue_init list_init

#define queue_destory list_destory

int queue_enqueue(Queue *queue, const void *data);

int queue_dequeue(Queue *queue, void **data);

#define queue_peek(queue) ((queue)->head == NULL ? NULL : (queue)->head->data)

#define queue_size list_size

#define printQ printL

#endif
