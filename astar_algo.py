from min_heap import MinHeap
class Node:
    def __init__(self, nodeId, hValue):
        self.id = nodeId # use node idx as id
        self.g = float("inf") # distance from the start node to the current node
        self.h = hValue # estimate distance from the current node to the target node
        self.f = float("inf") # total cost from the start node to the end node
        self.previousNode = None # use for backtrack

def AstartAlgorithm(graph, start, target, hValues):
    nodes = initializeNodes(graph, hValues)

    startNode = nodes[start]
    targetNode = nodes[target]

    # set distance to start node itself to 0
    startNode.g = 0
    startNode.f = startNode.g + startNode.h
  
   # init open list and close list
    openList = MinHeap([startNode])  # nodes to be expanded
    # closeList = set()  # nodes have expanded

    # repeat until the openList is empty
    while not openList.isEmpty():
        # remove the node with the lowest f-value
        currentNode = openList.remove()
        if currentNode == targetNode:
            break

        # populate all current nodes neighbors
        neighbors = graph[currentNode.id]
        for neighbor in neighbors:
            neighborIdx, distanceToNeighbor = neighbor
            neighborNode = nodes[neighborIdx]

            # check if neighbor in close list
            # if neighborNode in closeList:
            #     continue

            # check if find a better path
            newNeighborG = currentNode.g + distanceToNeighbor 
            if newNeighborG >= neighborNode.g:
                continue

            # update neighbor's g, h, f and previousNode
            neighborNode.previousNode = currentNode
            neighborNode.g = newNeighborG
            neighborNode.f = neighborNode.g + neighborNode.h

            # check if the neighbor in the openList
            if not openList.containsNode(neighborNode):
                openList.insert(neighborNode) 
            else:
                openList.update(neighborNode)

        # put current node to close list 
        # closeList.add(currentNode)

    return backtrackPath(targetNode)

def initializeNodes(graph, hValues):
    nodes = []
    for i in range(len(graph)):
        nodes.append(Node(i, hValues[i]))
    return nodes

def backtrackPath(targetNode):
    if targetNode.previousNode is None:
        return []
    
    currentNode = targetNode
    path = []

    while currentNode is not None:
        path.append(currentNode)
        currentNode = currentNode.previousNode
    
    return path[::-1]


if __name__ == "__main__":
    hValues = [20, 16, 6, 10, 4, 0]
    graph = [
            # for vertex 0
            [
                [1, 2], [3, 6]
            ],
            # for vertex 1
            [
                [0, 2], [2, 5]
            ],
            # for vertex 2
            [
                [1, 5], [3, 7], [4, 6], [5, 9]
            ],
            # for vertex 3
            [
                [0, 6], [2, 7], [4, 10]
            ],
            # for vertex 4
            [
                [2, 6], [3, 10], [5, 6]
            ],
            # for vertex 5
            [
                [2, 9], [4, 6]
            ],
        ]

    # startNode = 1, targetNode = 5
    path = AstartAlgorithm(graph, 0, 5, hValues)
    print("The shortest path")
    print("order, nodeId, f-value")
    for idx, node in enumerate(path):
        print(idx+1, node.id, node.f)

"""
The shortest path
order, nodeId, f-value
(1, 0, 20)
(2, 1, 18)
(3, 2, 13)
(4, 5, 16)
"""