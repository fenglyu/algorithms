#include<stdio.h>

struct node { int cnt; int* edges; };
struct graph { int V; int E; node *adj; };
