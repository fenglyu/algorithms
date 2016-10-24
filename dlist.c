#include <stdlib.h>
#include <string.h>
#include "dlist.h"


void dlist_init(DList *dlist, void(*destory)(void *data)){
    dlist->head = NULL;
    dlist->tail = NULL;
    dlist->size = 0;
    dlist->destory = destory;
    return ;
}


void dlist_destory(DList *dlist){
    void *data;
    while (dlist_size(dlist) > 0){
        if(dlist_remove(dlist, dlist_tail(dlist) , &data) == 0 && dlist->destory != NULL){
            dlist->destory(data);
        }
    }
    memset(dlist, 0, sizeof(DList));
    return;
}


int dlist_ins_next(DList *dlist, DListElem *element, const void *data){
    DListElem *new_element;

    // DO NOT allow a NULL element unless list is empty
    if(element == NULL && dlist_size(dlist) !=0 )
        return -1;

    if((new_element = malloc(sizeof(DListElem))) == NULL)
        return -1;

    new_element->data = (void *)data;

    if(dlist_size(dlist) == 0){
        // insert on head
        dlist->head = new_element;
        new_element->prev = NULL;
        new_element->next = NULL;
        dlist->tail = new_element;

    }else{
        new_element->next = element->next;
        new_element->prev = element;

        if (element->next == NULL){
            dlist->tail = new_element;
        }else{
            element->next->prev = new_element;
        }
        element->next = new_element;
    }

    dlist->size++;
    return 0;
}


int dlist_ins_prev(DList *dlist, DListElem *element, const void *data){

    DListElem *new_element;

    // DO NOT allow a NULL element unless list is empty
    if(element == NULL && dlist_size(dlist) != 0){
        return -1 ;
    }

    if((new_element = malloc(sizeof(struct _dlistelem))) == NULL){
        return -1;
    }

    new_element->data = (void *)data;

    if (dlist_size(dlist) == 0){
        dlist->head = new_element;
        new_element->next =NULL;
        new_element->prev = NULL;
        dlist->tail = new_element;
    }else{
        new_element->next = element;
        new_element->prev = element->prev;
        if(element->prev == NULL){
            dlist->head = new_element;
        }else{
            element->prev->next = new_element;
        }
        element->prev = new_element;
    }

    dlist->size++;
    return 0;
}


int dlist_remove(DList *dlist, DListElem *element, void **data){

    // DO NOT allow a NULL element or empty list
    if(element== NULL || dlist_size(dlist) ==  0)
        return -1;

    *data = element->data;

    /*
     *  2 pointers need to reset after element is removed,
     *  list->head move to element->next
     *  if element->next is NULL, reset dlist->tail
     *  else reset element->next->prev
     */
    if(element == dlist->head){
        dlist->head = element->next;

        if (dlist->head == NULL)
            dlist->tail = NULL;
        else
            element->next->prev = NULL;

    }else{
        element->prev->next = element->next;
        if( element->next == NULL )
            dlist->tail = element->prev;
        else
            element->next->prev = element->prev;
    }

    free(element);
    dlist->size--;

    return 0;
}
