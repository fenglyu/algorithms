#ifndef CLIST_H
#define CLIST_H
#include<stdlib.h>

typedef struct _clistelem{
    void *data;
    struct _clistelem *next;
}CListElem;

// Circle list
typedef struct _clist{
    int size;
    int (*match)(const void *a, const void *b);
    void (*destory)(void *data);
    CListElem *head;
}CList;


void clist_init(CList *list, void (*destory)(void *data));
void clist_destory(CList *list);
int clist_ins_next(CList *list, CListElem *element, const void *data);
int clist_rem_next(CList *list, CListElem *element, void **data);

#define clist_size(clist) ((clist)->size)
#define clist_head(clist) ((clist)->head)
#define clist_data(element) ((element)->data)
#define clist_next(element) ((element)->next)


#endif

