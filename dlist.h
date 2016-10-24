#ifndef DLIST_H
#include <stdio.h>

typedef struct _dlistelem{
    void *data;
    struct _dlistelem *prev;
    struct _dlistelem *next;
} DListElem;

typedef struct _dlist{
    int size;
    int (*match)(const void *a, const void *b);
    void (*destory)(void *val);
    DListElem *head;
    DListElem *tail;
} DList;


// function decalre
void dlist_init(DList *dlist, void(*destory)(void *data));
void dlist_destory(DList *dlist);
int dlist_ins_next(DList *dlist, DListElem *element, const void *data);
int dlist_ins_prev(DList *dlist, DListElem *element, const void *data);
int dlist_remove(DList *dlist, DListElem *element, void **data);

// macros
#define dlist_head(dlist) ((dlist)->head)
#define dlist_is_head(dlist, element) ((element) == (dlist)->head ? 1 : 0)

#define dlist_tail(dlist) ((dlist)->tail)
#define dlist_is_tail(dlist, element) ((dlist)->tail == (element) ? 1 : 0 )

#define dlist_prev(elemnt) ((elemnt)->prev)
#define dlist_size(dlist) ((dlist)->size)

#define dlist_next(element) ((element)->next)
#define dlist_data(elment) ((elment)->next)

#endif
