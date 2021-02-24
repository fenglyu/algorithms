#include <stdio.h>
#include <stdlib.h>
#include "set.h"


#define LEN 20

int match(const void *key1, const void *key2){
    return *(int *)key1 == *(int *)key2 ? 1 : 0 ;
}

/*
 *   make t_set
 *   valgrind --leak-check=full --show-leak-kinds=all ./t_set
 *
 */

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
//       printf("arr1 %d => %d\n", i, arr1[i]);
//       printf("arr2 %d => %d\n", i, arr2[i]);
    }


    printf("set1:\n");
    printS(set1);
//    ListElem *member;
//    void *data;
//    for(member = list_head(set1); member!=NULL; member = list_next(member)){
//        data = list_data(member);
//        printf("%d\n", *(int *)data);
//    }
    printf("set2:\n");
    printS(set2);


    Set *set;
    set = (Set *)malloc(sizeof(Set));
    if(!set)
        return -1;
    printf("Set Union:\n");
    set_init(set, match, NULL);
    set_union(set, set1, set2);
    printS(set);
    set_destory(set);

    printf("set intersection :\n");
    set_init(set, match, NULL);
    set_intersection(set, set1, set2);
    printS(set);
    set_destory(set);

    printf("set difference (set1 - set2) : \n");
    set_init(set, match, NULL);
    set_difference(set, set1, set2);
    printS(set);
    set_destory(set);

    printf("set difference (set2 - set1) : \n");
    set_init(set, match, NULL);
    set_difference(set, set2, set1);
    printS(set);
    set_destory(set);

    printf("set is subset :\n");
    set_init(set, match, NULL);
    set_intersection(set, set1, set2);
    if(set_is_subset(set, set2)){
        printf("YES\n");
    }else{
        printf("NOOOOO\n");
    }
    printS(set);
    set_destory(set);


    printf("set equal :\n");
    set_init(set, match, NULL);
    set_intersection(set, set1, set2);
    Set *set3;
    set3 = (Set *)malloc(sizeof(Set));
    if (!set3)
        return -1;

    set_intersection(set3, set1, set2);
    if(set_is_equal(set, set3))
        printf("YES\n");
    else
        printf("NOOOOOO\n");

    set_destory(set);
    set_destory(set3);
    free(set3);


    // end
    set_destory(set1);
    set_destory(set2);

    free(set);
    free(set1);
    free(set2);

    return 0;
}
