#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "list.h"


void list_init(List *list, void (*destory)(void *data)){
    list->size = 0;
    list->destory = destory;
    list->head = NULL;
    list->tail = NULL;

    return;
}


void list_destory(List *list){
    void *data;
    while(list_size(list) > 0){
        if(list_rem_next(list, NULL, (void **)&data) == 0 &&  list->destory != NULL){
            list->destory(data);
        }
    }

    memset(list, 0, sizeof(List));
    return;
}

int list_ins_next(List *list, ListElem *element, const void *data){
    ListElem *new_element;

    if ((new_element = (ListElem *)malloc(sizeof(struct _ListElem))) == NULL) {
        return -1;
    }

    new_element->data = (void *)data;

    if(element == NULL){
        // handle insertion at the head of the list
        if(list_size(list) == 0 ){
            list->tail = new_element;
        }
        new_element->next = list->head;
        list->head = new_element;
    }
    else {
        if(element->next == NULL){
            list->tail = new_element;
        }
        // this confusion of below 2 lines cost me the whole night to find the circle list;
        new_element->next = element->next;
        element->next = new_element;
    }

    list->size++;

    return 0;
}


int list_rem_next(List *list, ListElem *element, void **data){

    ListElem *old_element;

    if(list_size(list) == 0){
        return -1;
    }

    if (element == NULL){
        *data = list->head->data;
        old_element = list->head;
        list->head = old_element->next;
    }else{

        if(element->next == NULL)
            return -1;

        *data = element->next->data;
        old_element = element->next;
        element->next = element->next->next;

        if(element->next == NULL)
            list->tail = element;
    }

    free(old_element);

    list->size--;

    return 0;
}

int list_for_reverse(List *list){
    ListElem *p, *q, *r;
    if(list->head == NULL || list->head->next ==NULL){
        return -1;
    }

    p = list->head;
    q = list->head->next;
    list->head->next = NULL;
    while(q){
        r = q->next;
        q->next = p;
        p = q;
        q = r;
    }
    list->head = p;
    return 0;
}


List *list_recurisively_invocation_reverse(List *list){
    if (list_size(list) == 0)
        return NULL;

    return NULL;
}


//void printL(List *list){
//
//    ListElem *e =  list->head;
//    while(e != NULL){
//        printf("%-4d", *(int *)e->data);
//        e = e->next;
//    }
//    printf("\n");
//    return ;
//}

