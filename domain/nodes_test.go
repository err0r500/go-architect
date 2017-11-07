package domain

import "testing"
import "github.com/stretchr/testify/assert"

var one = Node{
	LOC:     1,
	Path:    "one",
	Parents: []Node{two, five},
}
var two = Node{
	LOC:     10,
	Path:    "two",
	Parents: []Node{three, four},
}
var three = Node{
	LOC:     100,
	Path:    "three",
	Parents: []Node{four},
}
var four = Node{
	LOC:     1000,
	Path:    "four",
	Parents: []Node{five},
}
var five = Node{
	Path: "five",
	LOC:  10000,
}

func TestNode_getAllParentsLoc(t *testing.T) {
	assert.Equal(t, 11110, one.getAllParentsLoc(), "getParentLoc() failed")
	assert.Equal(t, 11100, two.getAllParentsLoc(), "getParentLoc() failed")
}

func TestNode_getAllParentsCount(t *testing.T) {
	assert.Equal(t, 4, one.getAllParentsCount(), "getParentLoc() failed")
	assert.Equal(t, 3, two.getAllParentsCount(), "getParentLoc() failed")
}
