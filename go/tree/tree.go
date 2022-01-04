package tree

type NodeInterface interface {
	//Len() int
	//Less(i,j int) bool
	Less(v any) bool
	//Compare(i, j interface{}) int
	//Swap(i, j int)
	//Insert(r *BinaryNode, v interface{}) interface{}
}
