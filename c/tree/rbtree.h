#ifndef RBTREE_H
#define RBTREE_H

typedef uint_t rbtree_key_t;
typedef int_t  rbtree_key_int_t;


typedef struct rbtree_node_s rbtree_node_t;

struct rbtree_node_s {
	rbtree_key_t key;
	rbtree_nodes_t *left;	
	rbtree_nodes_t *right;	
	rbtree_nodes_t *parent;	
	u_char 	*color;
	u_char  data;
}

typedef struct rbtree_s rbtree_t;


typdef void (*rbtree_insert_pt)(rbtree_node_t *root,
	rbtree_node_t *node, rbtree_node_t *sentinel);


struct rbtree_s {
	rbtree_t *root;
	rbtree_t *sentinel;
	rbtree_insert_pt insert;	
}


#endif

