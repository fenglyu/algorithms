#include <stdio.h>
#include "list.h"

#define LEN 20

int main(){

    List *list;

    list = (List *)malloc(sizeof(List));
    if(!list)
        return -1;

    list_init(list, NULL);

    int i, arr[LEN];

    for(i = 0; i < LEN; i++){
        arr[i] = i + 10;
        list_ins_next(list, NULL, &arr[i]);
    }

    printL(list);

    list_destory(list);
    free(list);

    return 0;
}
