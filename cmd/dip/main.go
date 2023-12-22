package main

import "fmt"

type Person struct {
	name string
}

type RelationshipType int

const (
	parent RelationshipType = iota
	child
	// grandfather
	// grandchild
)

// Low level type
type Relationship struct {
	from    *Person
	relType RelationshipType
	to      *Person
}

// High level type
type Family struct {
	relationships []Relationship
}

func (f *Family) FindAllChildren(name string) []*Person {
	p := make([]*Person, 0)
	for i, v := range f.relationships {
		if v.relType == child &&
			v.to.name == name {
			p = append(p, f.relationships[i].from)
		}
	}
	return p
}

// High level type
type FamilyTreeChatBot struct {
	relations []Relationship // voilates DIP
}

// voilates DIP - see how deeply is FamilyTreeChatBot coupled to Relationship struct
// Two issues.
// 1. DIP
// 2. Such implementation could be contained in the low level model!
func (tree *FamilyTreeChatBot) FindAllChildren(name string) []*Person {
	p := make([]*Person, 0)
	for i, v := range tree.relations {
		if v.relType == child &&
			v.to.name == name {
			p = append(p, tree.relations[i].from)
		}
	}
	return p
}

type RelationshipBrowser interface {
	FindAllChildren(name string) []*Person
}

// High level type compliant DIP
// Introduce an abstraction
type BetterFamilyTreeChatBot struct {
	browser RelationshipBrowser
}

func main() {
	p1 := Person{"John"}
	p2 := Person{"Chris"}
	p3 := Person{"Matt"}

	relations := []Relationship{
		{&p2, child, &p1},
		{&p3, child, &p1},
		{&p1, parent, &p2},
		{&p1, parent, &p3},
	}

	chatBot := FamilyTreeChatBot{relations}
	children := chatBot.FindAllChildren("John")
	for _, v := range children {
		fmt.Println("John has a children called (Old)", v.name)
	}

	// Better as-per DIP
	betterChatBot := BetterFamilyTreeChatBot{&Family{relations}}
	children = betterChatBot.browser.FindAllChildren("John")
	for _, v := range children {
		fmt.Println("John has a children called (New)", v.name)
	}
}
