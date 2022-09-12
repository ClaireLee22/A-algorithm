class MinHeap:
    def __init__(self, array):
        self.nodePositionInHeap = {node.id: idx for idx, node in enumerate(array) }
        self.heap = self.buildHeap(array)
    
    # Time: O(n) | Space: O(1)
    def buildHeap(self, array):
        lastParentNodeIdx = (len(array) - 2) // 2
        for currentIdx in range(lastParentNodeIdx, -1, -1):
            self.siftDown(currentIdx, len(array)- 1, array)
        return array
    
    # Time: O(logn) | Space: O(1)
    def remove(self):
        if self.isEmpty():
            return 

        self.swap(0, len(self.heap)-1, self.heap)
        nodeToRemove = self.heap.pop()
        del self.nodePositionInHeap[nodeToRemove.id]
        self.siftDown(0, len(self.heap)-1, self.heap)
        return nodeToRemove
    
    # Time: O(logn) | Space: O(1)
    def insert(self, node):
        self.heap.append(node)
        self.nodePositionInHeap[node.id]= len(self.heap)-1
        self.siftUp(len(self.heap)-1, self.heap)

    # Time: O(logn) | Space: O(1)
    def siftDown(self, currentIdx, endIdx, heap):
        childOneIdx = currentIdx*2 + 1
        while childOneIdx <= endIdx:
            childTwoIdx = currentIdx*2 + 1 if currentIdx*2 + 1 <= endIdx else -1
            idxToSwap = childOneIdx
            if childTwoIdx != -1 and heap[childTwoIdx].f < heap[childOneIdx].f:
                idxToSwap = childTwoIdx
            
            if heap[idxToSwap].f < heap[currentIdx].f:
                self.swap(idxToSwap, currentIdx, heap)
                currentIdx = idxToSwap
                childOneIdx = currentIdx*2 + 1
            else:
                return
            
    # Time: O(logn) | Space: O(1)
    def siftUp(self, currentIdx, heap):
        parentIdx = (currentIdx - 1) // 2
        while currentIdx > 0 and heap[currentIdx].f < heap[parentIdx].f:
                self.swap(currentIdx, parentIdx, heap)
                currentIdx = parentIdx
                parentIdx = (currentIdx - 1) // 2
    
    def containsNode(self, node):
        return node.id in self.nodePositionInHeap
    
    def update(self, node):
        # update is occured when a slower f is found
        self.siftUp(self.nodePositionInHeap[node.id], self.heap)

    def isEmpty(self):
        return len(self.heap) == 0

    def swap(self, i, j, heap):
        self.nodePositionInHeap[heap[i].id] = j
        self.nodePositionInHeap[heap[j].id] = i
        heap[i], heap[j] = heap[j], heap[i]