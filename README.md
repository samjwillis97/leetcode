# leetcode
Personal LeetCode Solutions for Practice

## Algorithms

- Binary Search
  - It is used to find the position of an element in the array only if the array is sorted.
- Two Pointer
  - Given an array of integers (already sorted in ascending order) find and return the indices of the two elements that, when added together, are equal to the provided target value.
  - With this, we can start with a pointer on the left and right, and move them depending if the current sum is greater or less than the target 
- Sliding Window
- Johnson and Trotter
- Breadth First Search
  - Graph Traversal Approach
  - Start at Source Node and Layer by Layer through the graph, anlyse the nodes directly related to the source node. Then, you must move on to the next level neighbour nodes
  - Use a FIFO queue data structure to store the node and mark it as "visited" until it marks all the neighbouring vertices directly related to it.
    - The nodes neighbors will be viewed in the order in which it inserts them in the node, starting with the node that was inserted first
- Depth First Search
- Backtracking
  - Its like a splitting of paths, see https://leetcode.com/problems/letter-case-permutation/discuss/379928/Python-clear-solution
  - Technique for solving problems recursively, by trying to build a solution incrementally, one piece at a time.
  - Algorithm:
    - If it is not a solution, return false
    - if it is a new solution, add to list
- DFS + Backtracking Template
```python
def dfs( parameter ):

	if stop condtion or base case:
		# base case:
		update result
	    return
	
	else:
		# general cases:
		for all possible next moves:
		
		    select one next move
			dfs( paramter with selected next move )
			undo the selection
	
		return

```


## Data Structures

- Stack
  - FILO Queue
  
## Revise

- 77_Combination