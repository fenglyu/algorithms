#include <stdlib.h>
#include <string.h>
#include "list.h"
#include "set.h"



/* set_init */
void set_init(Set *set, int (*match)(const void *key1, const void *key2), void (*destory)(void *data)){
    list_init(set, destory);
    set->match = match;

    return;
}


int set_insert(Set *set, const void *data){
    if(set_is_member(set, data)){
        return 1;
    }

    return list_ins_next(set, list_tail(set), data);
}

int set_remove(Set *set, void **data){
    ListElem *member, *prev;

    for(member = list_head(set); member != NULL; member = list_next(member)){
        if(set->match(data, list_data(member)))
            break;
        prev = member;
    }

    return list_rem_next(set, prev, data);
}

int set_union(Set *set, const Set *set1, const Set *set2){
    ListElem *member;
    void *data;

    set_init(set, set1->match, set1->destory);

    for(member = list_head(set1); member != NULL; member = list_next(member)){
        data = list_data(member);
        if(list_ins_next(set, list_tail(set), data) != 0){
            set_destory(set);
            return -1;
        }
    }

    for(member = list_head(set2); member != NULL; member = list_next(member)){
        if(set_is_member(set1, list_data(member))){
            continue;
        }

        data = list_data(member);
        if(list_ins_next(set, list_tail(set), data) != 0 ){
            set_destory(set);
            return -1;
        }
    }

    return 0;
}

int set_intersection(Set *set, const Set *set1, const Set *set2){

    ListElem *member;
    void *data;

    set_init(set, set1->match, set1->destory);

    for(member = list_head(set1); member != NULL; member = list_next(member)){
        if (set_is_member(set2, list_data(member))){
            data = list_data(member);
           if( list_ins_next(set, list_tail(set), data) != 0){
               set_destory(set);
               return -1;
            }
        }else{
            continue;
        }
    }

    return 0;
}

int set_difference(Set *set, const Set *set1, const Set *set2){
    ListElem *member;
    void *data;

    set_init(set, set1->match, set1->destory);

    for(member = list_head(set1); member != NULL; member = list_next(member)){
        if(!set_is_member(set2, list_data(member))){
            data = list_data(member);

            // use list_ins_next instead of set_insert because list_ins_next
            // is O(1), while set_insert is O(n), data is guaranteed in set
            if( list_ins_next(set,list_tail(set), data) != 0 ){
                set_destory(set);
                return -1;
            }
        }
    }

    return 0;
}

int set_is_member(const Set *set, const void *data){
    ListElem *member;
    for(member = list_head(set); member != NULL; member = list_next(member)){
       if(set->match(list_data(member), data))
           return 1;
    }

    return 0;
}

/* if set1 is a subset of set2 */
int set_is_subset(const Set *set1, const Set *set2){

    ListElem *member;

    if(set_size(set1) > set_size(set2))
        return 0;

    for (member = list_head(set1); member != NULL; member = list_next(member)){
        if(!set_is_member(set2, list_data(member)))
           return 0;
    }

    return 1;
}

int set_is_equal(const Set *set1, const Set *set2){
    if(set_size(set1) != set_size(set2)){
        return 0;
    }

    return set_is_subset(set1, set2);
}
