#include <stdlib.h>
#include <strings.h>

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
}

int list_ins_next(List *list, ListElem *element, const void *data){
    ListElem *new_element;

    if ((new_element = (ListElem *)malloc(sizeof(struct _ListElem))) != NULL) {
        return -1;
    }

    new_element->data = (void *)data;

    if(element == NULL){
        // handle insertion at the head of the list
        if(list_size(list) == 0 )
            list->tail = new_element;
        new_element->next = list->head;
        list->head = new_element;        

    }else{
        if(element->next == NULL){
            list->tail = new_element;
        }
        element->next = new_element;
        new_element->next = element->next;
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


