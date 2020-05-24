# Trees And Graphs

## Types of Trees
A tree is a data structure composed of nodes.
- Each tree has a root node.
- Root node has zero or more child nodes
- Each child node has zero or more child nodes etc.
- A tree cannot contain cycles

Trees and Graphs have many ambigous details. Watch out for the following issues:
- **Trees vs. Binary Tree.** Binary tree has up to two children.
- **Binary Tree vs. Binary Search Tree.** BST: for each node `n` `all left <= n < all right` must hold. 
- **Balanced vs. Unbalanced.** Balanced means balanced enough to ensure O(log n) for insert and find ops. Common balanced trees are `red-black trees` and `AVL trees`.
- **Complete Binary Trees.** Every level is filled except maybe the last. Last is filled left to right.
- **Full Binary Trees.** Every node has 0 or 2 children.
- **Perfect Binary Trees.** Both `full` and `complete`.



## Binary Tree Traversal
- **In-Order Traversal.** left -> current -> right
- **Pre-Order Traversal** current -> left -> right
- **Post-Order Traversal** left -> right -> current

## Binary Heaps (Min/Max Heaps)
A min-heap is a `complete` binary tree where the root is the min value, and each node is smaller than its children.

Key Operations:
- `insert` insert at the bottom rightmost spot and swap with parent until conditions met.
- `extract_min` root node is min.


## Tries (Prefix Trees)
A trie is a tree in which each node stores a character and each path may represent a word. The node `*` is often used to indicate an end to a complete word.

## Graphs
Graphs are a collection of nodes which may or may not have connections/edges between them.
A graph can:
- Be directed(one way) or undirected(two way).
- Consist of multiple isolated sub-graphs. A connected graph is graph that has a path between every pair of vertices.
- have cycles. A graph without cycles is called a acyclic graph
- be represented in `adjacency list` or `adjacency matrices`
- be searched by BFS(breadth-first search), DFS(depth-first search), and bidirectional search (2 simultaneous BFS) algorithms.