#include <stdio.h>
#include <stdlib.h>
#include "set.h"


#define LEN 20

int match(const void *key1, const void *key2){
    return *(int *)key1 - *(int *)key2;
}

int main(void){
    
    Set *set1, *set2;
    int arr1[LEN], arr2[LEN];
    int i;

    set1 = (Set *)malloc(sizeof(Set));
    if (!set1)
        return -1;

    set_init(set1, match, NULL);

    set2 = (Set *)malloc(sizeof(Set));
    if (!set2)
        return -1;

    set_init(set2, match, NULL);
    
    for(i = 0; i < LEN; i++){
        arr1[i] = i + 10;
        arr2[i] = i + 15;
        set_insert(set1, &arr1[i]);
        set_insert(set2, &arr2[i]);
        printf("arr1 %d => %d\n", i, arr1[i]);
        printf("arr2 %d => %d\n", i, arr2[i]);
    }


    printf("set1:\n");
//    printS(set1);
    ListElem *member;
    void *data;
    for(member = list_head(set1); member!=NULL; member = list_next(member)){
        data = list_data(member);
        printf("%d\n", *(int *)data);
    }
    printf("set2:\n");
//    printS(set2);

    return 0; 
}
