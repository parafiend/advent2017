package puzzle07

import (
	"fmt"
	"log"
)

type Node struct {
	Name        string
	Weight      int
	Parent      *Node
	Children    map[string]*Node
	stackWeight int
}

func (node *Node) GetOldestParent() *Node {
	for node.Parent != nil {
		node = node.Parent
	}
	return node
}

func (node *Node) StackWeight() int {
	if node.stackWeight == 0 {
		node.stackWeight = node.Weight
		for _, child := range node.Children {
			node.stackWeight += child.StackWeight()
		}
	}
	return node.stackWeight
}

func (node *Node) Balanced() bool {
	balanced := true
	weights := make([]int, 0)
	sum := 0
	for _, child := range node.Children {
		childWeight := child.StackWeight()
		weights = append(weights, childWeight)
		sum += childWeight
	}

	log.Println(node.Name, sum, weights)
	for i, weight := range weights[1:] {
		if weight != weights[0] {
			log.Println("XXX", i, weight, weights[0])
			balanced = false
		}
	}
	return balanced
}

func (node *Node) FindUnbalance() *Node {
	var result *Node
	for !node.Balanced() {
		allBalanced := true
		var next *Node
		for _, child := range node.Children {
			if !child.Balanced() {
				next = child
				allBalanced = false
			}
		}
		if allBalanced {
			log.Println("----------", node, "----------")
			for _, childs := range node.Children {
				log.Println(childs.Name, childs.Weight, childs.StackWeight())
			}
			log.Println("----------")
		}
		if next == nil {
			break
		} else {
			node = next
		}
	}
	return result
}

func (node *Node) String() string {
	parent := "<nil>"
	if node.Parent != nil {
		parent = node.Parent.Name
	}
	return fmt.Sprintf("%v (w%v, p%v, %v)", node.Name, node.Weight, parent, len(node.Children))
}

type Stack struct {
	Root          *Node
	NodesByName   map[string]*Node
	ParentsByName map[string]*Node
}

func NewStack() Stack {
	return Stack{nil, make(map[string]*Node), make(map[string]*Node)}
}

func (s *Stack) CreateNode(name string, weight int, children []string) {
	s.getOrMakeNode(name, weight, children, nil)
	log.Println(s.Root, s.NodesByName)
}

func (s *Stack) getOrMakeNode(name string, weight int, children []string, parent *Node) *Node {
	var node *Node
	var found bool
	if node, found = s.NodesByName[name]; !found {
		rawNode := Node{name, weight, parent, make(map[string]*Node), 0}
		node = &rawNode
		log.Println("NewNode!", node)
		s.NodesByName[name] = node
	} else if node.Weight < 0 {
		node.Weight = weight
	}

	for _, childName := range children {
		child := s.getOrMakeNode(childName, -1, make([]string, 0), node)
		node.Children[childName] = child
		child.Parent = node
		if child == s.Root {
			s.Root = node.GetOldestParent()
			log.Println("NewRoot!", s.Root)
		}
	}

	if s.Root == nil {
		s.Root = node
		log.Println("FirstRoot!", s.Root, node)
	}
	return node
}
