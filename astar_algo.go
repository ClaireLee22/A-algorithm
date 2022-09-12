package main

import (
	"fmt"
	"math"
)

func AstartAlgorithm(graph [][][]int, hValues []int, start int, target int) []*Node {
	nodes := initializeNodes(graph, hValues)

	startNode := nodes[start]
	targetNode := nodes[target]

	// set distance to start node itself to 0
	startNode.g = 0
	startNode.f = startNode.g + startNode.h

	// init open list and close list
	openList := newMinHeap([]*Node{startNode})
	// closeList := make(map[int]bool)

	// repeat until the openList is empty
	for !openList.isEmpty() {
		// remove the node with the lowest f-value
		currentNode := openList.remove()
		if currentNode == targetNode {
			break
		}

		// populate all current nodes neighbors
		neighbors := graph[currentNode.id]
		for _, neighbor := range neighbors {
			neighborIdx, distanceToNeighbor := neighbor[0], neighbor[1]
			neighborNode := nodes[neighborIdx]

			// check if neighbor in close list
			// if _, found := closeList[neighborIdx]; found {
			// 	continue
			// }

			// check if find a better path
			newNeighborG := currentNode.g + distanceToNeighbor
			if newNeighborG >= neighborNode.g {
				continue
			}

			// update neighbor's g, h, f and previousNode
			neighborNode.g = newNeighborG
			neighborNode.f = neighborNode.g + neighborNode.h
			neighborNode.previousNode = currentNode

			// check if the neighbor in the openList
			if openList.containsNode(neighborNode) {
				openList.update(neighborNode)
			} else {
				openList.insert(neighborNode)
			}
		}
		// put current node to close list
		// closeList[currentNode.id] = true
	}
	return backtrackPath(targetNode)
}

func initializeNodes(graph [][][]int, hValues []int) []*Node {
	nodes := []*Node{}
	for i := range graph {
		nodes = append(nodes, &Node{
			id:           i,
			g:            math.MaxInt32,
			h:            hValues[i],
			f:            math.MaxInt32,
			previousNode: nil,
		})
	}
	return nodes
}

func backtrackPath(targetNode *Node) []*Node {
	if targetNode.previousNode == nil {
		return []*Node{}
	}

	currentNode := targetNode
	path := []*Node{}
	for currentNode != nil {
		path = append(path, currentNode)
		currentNode = currentNode.previousNode
	}

	return reverse(path)
}

func reverse(array []*Node) []*Node {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

type Node struct {
	id           int // use node idx as id
	g            int // distance from the start node to the current node
	h            int // estimate distance from the current node to the target node
	f            int // # total cost from the start node to the end node
	previousNode *Node
}

type MinHeap struct {
	array              []*Node
	nodePositionInHeap map[int]int
}

func newMinHeap(array []*Node) *MinHeap {
	nodePositionInHeap := map[int]int{}
	for idx, node := range array {
		nodePositionInHeap[node.id] = idx
	}

	heap := &MinHeap{
		array:              array,
		nodePositionInHeap: nodePositionInHeap,
	}
	heap.buildHeap(array)
	return heap
}

func (h *MinHeap) buildHeap(array []*Node) {
	lastParentNode := (len(array) - 2) / 2
	for currentIdx := lastParentNode; currentIdx >= 0; currentIdx-- {
		h.siftDown(currentIdx, len(array)-1)
	}
}

func (h *MinHeap) remove() *Node {
	if h.isEmpty() {
		return nil
	}

	h.swap(0, len(h.array)-1)
	nodeToRemove := h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	delete(h.nodePositionInHeap, nodeToRemove.id)
	h.siftDown(0, len(h.array)-1)
	return nodeToRemove
}

func (h *MinHeap) insert(node *Node) {
	h.array = append(h.array, node)
	h.nodePositionInHeap[node.id] = len(h.array) - 1
	h.siftUp(len(h.array) - 1)
}

func (h *MinHeap) siftDown(currentIdx, endIdx int) {
	childOneIdx := currentIdx*2 + 1
	for childOneIdx <= endIdx {
		childTwoIdx := currentIdx*2 + 2
		if childTwoIdx > endIdx {
			childTwoIdx = -1
		}

		idxToSwap := childOneIdx
		if childTwoIdx != -1 && h.array[childTwoIdx].f < h.array[childOneIdx].f {
			idxToSwap = childTwoIdx
		}

		if h.array[idxToSwap].f < h.array[currentIdx].f {
			h.swap(idxToSwap, currentIdx)
			currentIdx = idxToSwap
			childOneIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp(currentIdx int) {
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && h.array[currentIdx].f < h.array[parentIdx].f {
		h.swap(parentIdx, currentIdx)

		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

func (h *MinHeap) containsNode(node *Node) bool {
	_, found := h.nodePositionInHeap[node.id]
	return found
}

func (h *MinHeap) isEmpty() bool {
	return len(h.array) == 0
}

func (h *MinHeap) update(node *Node) {
	// update is occured when a slower f is found
	h.siftUp(h.nodePositionInHeap[node.id])
}

func (h *MinHeap) swap(i, j int) {
	h.nodePositionInHeap[h.array[i].id] = j
	h.nodePositionInHeap[h.array[j].id] = i
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	hValues := []int{20, 16, 6, 10, 4, 0}
	graph := [][][]int{
		// for vertex 0
		{
			{1, 2}, {3, 6}, // {vertexIdx, distance} ex {1, 2} distance from vertex 0 to vertex 1 is 2
		},
		// for vertex 1
		{
			{0, 2}, {2, 5},
		},
		// for vertex 2
		{
			{1, 5}, {3, 7}, {4, 6}, {5, 9},
		},
		// for vertex 3
		{
			{0, 6}, {2, 7}, {4, 10},
		},
		// for vertex 4
		{
			{2, 6}, {3, 10}, {5, 6},
		},
		// for vertex 5
		{
			{2, 9}, {4, 6},
		},
	}

	// startNode: 0, targetNode: 5
	path := AstartAlgorithm(graph, hValues, 0, 5)
	fmt.Println("The shortest path")
	fmt.Println("order, nodeId, f-value")
	for idx, node := range path {
		fmt.Println(idx+1, node.id, node.f)
	}
}
