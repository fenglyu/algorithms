#include <stdlib.h>
#include "dlist.h"


void printL(DList *dlist){

    DListElem *e =  dlist->head;
    while(e != NULL){
        printf("%d\n", *(int *)e->data);
        e = e->next;
    }
    return ;
}

int main(){
    DList *dlist = malloc(sizeof(DList));
    if(dlist == NULL)
        return -1;

    dlist_init(dlist, NULL);

    int arr[20];
    int i;

    for(i=0; i< 10; i++){
        arr[i] = i;
        dlist_ins_next(dlist, dlist_head(dlist), &arr[i]);
    }

    for(i=10; i< 20; i++){
        arr[i] = i;
        dlist_ins_prev(dlist, dlist_head(dlist), &arr[i]);
    }

    printf(" The double linked list:\n");
    printL(dlist);


    dlist_destory(dlist);
    free(dlist);
    return 0;
}
