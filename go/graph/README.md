# Algorithms Course - Graph Theory Tutorial from a Google Engineer
> [youtube](https://www.youtube.com/watch?v=09_LlHjoEiY)

```bash
⭐️ Course Contents ⭐️
⌨️ (0:00:00) Graph Theory Introduction
⌨️ (0:13:53) Problems in Graph Theory
⌨️ (0:23:15) Depth First Search Algorithm
⌨️ (0:33:18) Breadth First Search Algorithm
⌨️ (0:40:27) Breadth First Search grid shortest path
⌨️ (0:56:23) Topological Sort Algorithm
⌨️ (1:09:52) Shortest/Longest path on a Directed Acyclic Graph (DAG)
⌨️ (1:19:34) Dijkstra's Shortest Path Algorithm
⌨️ (1:43:17) Dijkstra's Shortest Path Algorithm | Source Code
⌨️ (1:50:47) Bellman Ford Algorithm
⌨️ (2:05:34) Floyd Warshall All Pairs Shortest Path Algorithm
⌨️ (2:20:54) Floyd Warshall All Pairs Shortest Path Algorithm | Source Code
⌨️ (2:29:19) Bridges and Articulation points Algorithm
⌨️ (2:49:01) Bridges and Articulation points source code
⌨️ (2:57:32) Tarjans Strongly Connected Components algorithm
⌨️ (3:13:56) Tarjans Strongly Connected Components algorithm source code
⌨️ (3:20:12) Travelling Salesman Problem | Dynamic Programming
⌨️ (3:39:59) Travelling Salesman Problem source code | Dynamic Programming
⌨️ (3:52:27) Existence of Eulerian Paths and Circuits
⌨️ (4:01:19) Eulerian Path Algorithm
⌨️ (4:15:47) Eulerian Path Algorithm | Source Code
⌨️ (4:23:00) Prim's Minimum Spanning Tree Algorithm
⌨️ (4:37:05) Eager Prim's Minimum Spanning Tree Algorithm
⌨️ (4:50:38) Eager Prim's Minimum Spanning Tree Algorithm | Source Code
⌨️ (4:58:30) Max Flow Ford Fulkerson | Network Flow
⌨️ (5:11:01) Max Flow Ford Fulkerson | Source Code
⌨️ (5:27:25) Unweighted Bipartite Matching | Network Flow
⌨️ (5:38:11) Mice and Owls problem | Network Flow
⌨️ (5:46:11) Elementary Math problem | Network Flow
⌨️ (5:56:19) Edmonds Karp Algorithm | Network Flow
⌨️ (6:05:18) Edmonds Karp Algorithm | Source Code
⌨️ (6:10:08) Capacity Scaling | Network Flow
⌨️ (6:19:34) Capacity Scaling | Network Flow | Source Code
⌨️ (6:25:04) Dinic's Algorithm | Network Flow
⌨️ (6:36:09) Dinic's Algorithm | Network Flow | Source Code
```

Exercises

17.1 Prove that any acyclic connected graph that has V vertices has V – 1 edges.
```
You are given a definition: "an acyclic, connected graph is called a tree."

You are then presented with a claim that this definition is equivalent to "has 𝑉−1 edges and is acyclic" and also equivalent to "has 𝑉−1 edges and is connected." This claim is not proved in the text, and indeed the purpose of the exercise is to prove part of this claim. Thus you may not use the claim when doing the exercise, and must rely on only the initial definition of tree that you were given.

The usual approach is induction on 𝑉. If 𝑉=1, then obviously the claim holds.

Now suppose we have proven the claim for 𝑉=𝑛. We consider an acyclic connected graph with 𝑛+1 vertices and seek to show it has 𝑛 edges. Such a graph must have a leaf (vertex of degree 1). Deleting that vertex and its accompanying edge will produce a graph that is also acyclic and connected. By the inductive hypothesis, this smaller graph has 𝑛−1 edges, so the original graph has 𝑛 edges.
```

Image 17.2 Give all the connected subgraphs of the graph

                           0-1 0-2 0-3 1-3 2-3.

Image 17.3 Write down a list of the nonisomorphic cycles of the graph in Figure 17.1. For example, if your list contains 3-4-5-3, it should not contain 3-5-4-3, 4-5-3-4, 4-3-5-4, 5-3-4-5, or 5-4-3-5.


17.4 Consider the graph

              3-7 1-4 7-8 0-5 5-2 3-8 2-9 0-6 4-9 2-6 6-4.
Determine the number of connected components, give a spanning forest, list all the simple paths with at least three vertices, and list all the nonisomorphic cycles (see Exercise 17.3).


Image 17.5 Consider the graphs defined by the following four sets of edges:

        0-1 0-2 0-3 1-3 1-4 2-5 2-9 3-6 4-7 4-8 5-8 5-9 6-7 6-9 7-8
        0-1 0-2 0-3 0-3 1-4 2-5 2-9 3-6 4-7 4-8 5-8 5-9 6-7 6-9 7-8
        0-1 1-2 1-3 0-3 0-4 2-5 2-9 3-6 4-7 4-8 5-8 5-9 6-7 6-9 7-8
        4-1 7-9 6-2 7-3 5-0 0-2 0-8 1-6 3-9 6-3 2-8 1-5 9-8 4-5 4-7
Which of these graphs are isomorphic to one another? Which of them are planar?


17.6 Consider the more than 68 billion graphs referred to in the caption to Figure 17.4. What percentage of them has fewer than nine vertices?


Image 17.7 How many different subgraphs are there in a given graph with V vertices and E edges?


17.8 Give tight upper and lower bounds on the number of connected components in graphs that have V vertices and E edges.


Image 17.9 How many different undirected graphs are there that have V vertices and E edges?


ImageImage 17.10 If we consider two graphs to be different only if they are not isomorphic, how many different graphs are there that have V vertices and E edges?


17.11 How many V -vertex graphs are bipartite?
