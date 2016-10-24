#ifndef LIST_H
#define LIST_H

#include<stdlib.h>


typedef struct _ListElem ListElem;

typedef struct _List List;

struct _ListElem{
    struct _ListElem *next;
    void *data;
};

struct _List{
    int size;
    int (*match)(const void *key1, const void *key2);
    void (*destory)(void *data);
    ListElem *head;
    ListElem *tail;
};


void list_init(List *list, void (*destory)(void *data));
void list_destory(List *list);
int list_ins_next(List *list, ListElem *elem, const void *data);
int list_rem_next(List *list, ListElem *elem, void **data);

int list_for_reverse(List *list);
int list_recurisively_invocation_reverse(List *list);

#define list_head(list) ((list)->head)
#define list_tail(list) ((list)->tail)
#define list_is_head(list, element) ((list)->head == (element) ? 1 : 0)
#define list_is_tail(list, element) ((element)->next == NULL ? 1 : 0)
#define list_size(list) ((list)->size)

#define list_data(element) ((element)->data)
#define list_next(element) ((element)->next)

#endif
