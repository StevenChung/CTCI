# 4.1  Route Between  Nodes:  Given a directed graph,
# design an algorithm to find out whether there is a
# route between two nodes.

class Graph():
	def __init__(self, verticeCount):
		self.max_vertices = verticeCount
		self.vertices = [0] * self.max_vertices
		self.count = 0

	def addNode(self, node):
		if self.count < self.max_vertices:
			self.vertices[self.count] = node
			self.count += 1
		else:
			print("Graph is full")

	def getNodes(self):
		return self.vertices

class Node():
	def __init__(self, vertex, noOfVertices):
		self.vertices = [0] * noOfVertices
		self.vertex = vertex
		self.adjacentCount = 0
		self.visited = False

	def addAdjacent(self, x):
		if self.adjacentCount < len(self.vertices):
			self.vertices[self.adjacentCount] = x
			self.adjacentCount += 1
		else:
			print("full")

	def getAdjacent(self):
		return self.vertices

	def getVertex(self):
		return self.vertex

def createNewGraph():
	g = Graph(6)
	temp = [0] * 6

	temp[0] = Node("a", 3)
	temp[1] = Node("b", 0)
	temp[2] = Node("c", 0)
	temp[3] = Node("d", 1)
	temp[4] = Node("e", 1)
	temp[5] = Node("f", 0)

	temp[0].addAdjacent(temp[1])
	temp[0].addAdjacent(temp[2])
	temp[0].addAdjacent(temp[3])
	temp[3].addAdjacent(temp[4])
	temp[4].addAdjacent(temp[5])

	for i in range(6):
		g.addNode(temp[i])
	return g

def createNewGraphWithLoop():
	g = Graph(6)
	temp = [0] * 6

	temp[0] = Node("a", 1)
	temp[1] = Node("b", 1)
	temp[2] = Node("c", 1)
	temp[3] = Node("d", 1)
	temp[4] = Node("e", 2)
	temp[5] = Node("f", 0)

	temp[0].addAdjacent(temp[1])
	temp[1].addAdjacent(temp[2])
	temp[2].addAdjacent(temp[3])
	temp[3].addAdjacent(temp[4])
	temp[4].addAdjacent(temp[1])
	# temp[4].addAdjacent(temp[5])

	for i in range(6):
		g.addNode(temp[i])
	return g

g = createNewGraphWithLoop()
n = g.getNodes()
start = n[0]
end = n[5]

def breadthfirstsearch(g, start, end):
    queue = [start]
    while len(queue) > 0:
        node = queue.pop(0)
        node.visited = True
        if node.vertex == end.vertex:
            return True
        for i in range(node.adjacentCount):
            if node.vertices[i].visited == False:
                queue.append(node.vertices[i])

    return False

print(breadthfirstsearch(g, start, end))
# basically, just iterate through the graph and flag visited
