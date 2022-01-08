##  Chapter 17. Graph Properties and Types
```
adjacency-matrix representation and the adjacency-lists representation, and various approaches to implementing basic ADT functions.
```


> We begin by working through the basic definitions of graphs and the properties of graphs, examining the standard nomenclature that is used to describe them. Following that, we define the basic ADT (abstract data type) interfaces that we use to study graph algorithms and the two most important data structures for representing graphs—the adjacency-matrix representation and the adjacency-lists representation, and various approaches to implementing basic ADT functions. Then, we consider client programs that can generate random graphs, which we can use to test our algorithms and to learn properties of graphs. All this material provides a basis for us to introduce graph-processing algorithms that solve three classical problems related to finding paths in graphs, which illustrate that the difficulty of graph problems can differ dramatically even when they might seem similar. We conclude the chapter with a review of the most important graph-processing problems that we consider in this book, placing them in context according to the difficulty of solving them.



Definition 17.1

A graph is a set of vertices plus a set of edges that connect pairs of distinct vertices (with at most one edge connecting any pair of vertices).


Property 17.1 A graph with V vertices has at most V (V – 1)/2 edges.

Proof: The total of V2 possible pairs of vertices includes V self-loops and accounts twice for each edge between distinct vertices, so the number of edges is at most (V2– V)/2=V (V – 1)/2.Image

```
17.51 Develop an implementation for the GRAPHconstruct function described in Exercise 17.49 that use a compact representation based on the following data structures:

         struct node { int cnt; int* edges; };
         struct graph { int V; int E; node *adj; };
A graph is a vertex count, an edge count, and an array of vertices. A vertex contains an edge count and an array with one vertex index corresponding to each adjacent edge. Implement GRAPHshow for this representation.
```


We adopt as standard this definition of a graph (which we first encountered in Chapter 5), but note that it embodies two technical simplifications. First, it disallows duplicate edges (mathematicians sometimes refer to such edges as parallel edges, and a graph that can contain them as a multigraph). Second, it disallows edges that connect vertices to themselves; such edges are called self-loops. Graphs that have no parallel edges or self-loops are sometimes referred to as simple graphs.


For some applications, we might consider the elimination of parallel edges and self-loops to be a data-processing problem that our implementations must address. For other applications, ensuring that a given set of edges represents a simple graph may not be worth the trouble. Throughout the book, whenever it is more convenient to address an application or to develop an algorithm by using an extended definition that includes parallel edges or self-loops, we shall do so. For example, self-loops play a critical role in a classical algorithm that we will examine in Section 17.4; and parallel edges are common in the applications that we address in Chapter 22. Generally, it is clear from the context whether we intend the term “graph” to mean “simple graph” or “multigraph” or “multigraph with self-loops.”


Placing the vertices of a given graph on the plane and drawing them and the edges that connect them is known as graph drawing. The possible vertex placements, edge-drawing styles, and aesthetic constraints on the drawing are limitless. Graph-drawing algorithms that respect various natural constraints have been studied heavily and have many successful applications (see reference section). For example, one of the simplest constraints is to insist that edges do not intersect. A planar graph is one that can be drawn in the plane without any edges crossing. Determining whether or not a graph is planar is a fascinating algorithmic problem that we discuss briefly in Section 17.8. Being able to produce a helpful visual representation is a useful goal, and graph drawing is a fascinating field of study, but successful drawings are often difficult to realize. Many graphs that have huge numbers of vertices and edges are abstract objects for which no suitable drawing is feasible.


For some applications, such as graphs that represent maps or circuits, a graph drawing can carry considerable information because the vertices correspond to points in the plane and the distances between them are relevant. We refer to such graphs as Euclidean graphs. For many other applications, such as graphs that represent relationships or schedules, the graphs simply embody connectivity information, and no particular geometric placement of vertices is ever implied. We consider examples of algorithms that exploit the geometric information in Euclidean graphs in Chapters 20 and 21, but we primarily work with algorithms that make no use of any geometric information, and stress that graphs are generally independent of any particular representation in a drawing or in a computer.

