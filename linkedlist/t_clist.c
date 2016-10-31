#include <stdio.h>
#include "clist.h"


void printCL(CList *list){
    if(clist_size(list) == 0)
        return; 

    CListElem *element = list->head;
    do{
        printf("%d\n", *(int *)element->data);  
        element = element->next;
    }while(element != list->head);
}

int main(void){
    
    CList *list = (CList *)malloc(sizeof(CList));

    if(!list){
        return -1;
    }

    clist_init(list, NULL);

    int i;
    int arr[20];

    for(i = 0; i < 20; i++){
        arr[i] = i+10;
        clist_ins_next(list, list->head, &arr[i]);
    }

    printCL(list);

    clist_destory(list);
    free(list);

    return 0;
}
