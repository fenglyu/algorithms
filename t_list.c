#include <stdio.h>
#include "list.h"


//void destory(void *data){
//   printf("destory");
//    free(data);
//}

void printL(List *dlist){

    ListElem *e =  dlist->head;
    while(e != NULL){
        printf("%d\n", *(int *)e->data);
        e = e->next;
    }
    return ;
}


int main(void){

    List *list;
    list = malloc(sizeof(List));
    if(list == NULL)
        return -1;
    //list_init(list, destory);
    list_init(list, NULL);
    printf("list size: %d\n", list_size(list));
    int i;
    int arr[100];
    for(i = 0; i < 10; i++){
        arr[i] = i;
//        printf("%p\n", &arr[i]);
        list_ins_next(list, NULL, &arr[i]);//arr+i );
    //    printf("list size: %d\n", list_size(list));
    }

    ListElem *element = list_head(list);
    while(element != NULL){
//        printf("%p\n", (int *)element->data);
        printf("%d\n", *(int *)element->data);
        element = list_next(element);
    }

    printf("list size: %d\n", list_size(list));
    printf("remove\n");
    int val = 222;
    ListElem *e =  list_head(list);
    list_ins_next(list, e, &val);
    printf("list size: %d\n", list_size(list));

    int *t;
    list_rem_next(list, list_next(list_next(list_head(list))), (void **)&t);
    ListElem *e1 = list_head(list);
    while(e1 != NULL){
        printf("%d\n", *(int *)e1->data);
        e1 = list_next(e1);
    }
    printf(" the removed value is %d\n", *t);

    printf("remote\n");
    printf("list size: %d\n", list_size(list));
    int *c;
    list_rem_next(list, NULL, (void **)&c);

    printf("list size: %d\n", list_size(list));
    printf(" the removed value is %d\n", *c);


    printf("Before:\n");
    printL(list);
    printf(" reverse list:\n");
    list_for_reverse(list);
    printf("After:\n");
    printL(list);

    list_destory(list);
    free(list);

    return 0;
}
