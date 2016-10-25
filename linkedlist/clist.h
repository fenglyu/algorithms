#ifndef _CLIST
#include<stdlib.h>

typedef struct _clistelem{
    void *data;
    struct _clistelem *next;
}ClistElem;

// Circle list
typedef struct _clist{
    int size;
    int (*match)(const void *a, const void *b);
    void (*destory)(void *data);
    ClistElem *head;
}CList;



#endif

