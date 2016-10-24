#include <stdio.h>
#include "list.h"


//void destory(void *data){
//   printf("destory");
//    free(data);
//}

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
    ListElem *e2 = list_head(list);
    i = 0;
    DListElem *t;
    while(e2 != NULL){
        if(i++%9 == 0)
            t = e2;
        printf("%d\n", *(int *)e2->data);
        e2 = list_next(e2);
    }
    printf("list size: %d\n", list_size(list));
    printf(" the removed value is %d\n", *c);



    list_destory(list);
    if(list != NULL)
        free(list);
    return 0;
}
