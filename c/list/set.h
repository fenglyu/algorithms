#ifndef SET_H
#define SET_H

#include <stdlib.h>
#include "list.h"


typedef List Set;

/* public interface */

void set_init(Set *set, int (*match)(const void *key1, const void *key2), void (*destory)(void *data));

#define  set_destory list_destory

int set_insert(Set *set, const void *data);

int set_remove(Set *set, void **data);

int set_union(Set *set, const Set *set1, const Set *set2);

int set_intersection(Set *set, const Set *set1, const Set *set2);

int set_difference(Set *set, const Set *set1, const Set *set2);

int set_is_member(const Set *set, const void *data);

int set_is_subset(const Set *set, const Set *set1);

int set_is_equal(const Set *set, const Set *set1);

#define set_size list_size

#define printS printL

#endif
