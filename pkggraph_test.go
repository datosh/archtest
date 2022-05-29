package archtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkgGraph_AddNode(t *testing.T) {
	graph := NewPkgGraph()

	graph.AddNode("foo")

	assert.Equal(t, "foo", graph.GetNode("foo").pkgPath)
}

func TestPkgGraph_GetNode_NotAdded(t *testing.T) {
	graph := NewPkgGraph()

	assert.Nil(t, graph.GetNode("foo"))
}

func TestPkgGraph_Size(t *testing.T) {
	graph := NewPkgGraph()

	graph.Load("github.com/matthewmcnew/archtest/examples/dependency")

	assert.Equal(t, 3, graph.Size())
}

func TestPkgGraph_Size_Expanding(t *testing.T) {
	graph := NewPkgGraph()

	graph.Load("github.com/matthewmcnew/archtest/examples/...")

	assert.Equal(t, 11, graph.Size())
}

func TestPkgGraph_IsDependingOn(t *testing.T) {
	assert := assert.New(t)
	graph := NewPkgGraph()

	graph.Load("github.com/matthewmcnew/archtest/examples/dependency")

	assert.True(graph.GetNode("github.com/matthewmcnew/archtest/examples/dependency").
		IsDependingOn("github.com/matthewmcnew/archtest/examples/transative"),
	)
}

func TestPkgGraph_IsDependedOnBy(t *testing.T) {
	assert := assert.New(t)
	graph := NewPkgGraph()

	graph.Load("github.com/matthewmcnew/archtest/examples/dependency")

	assert.True(graph.GetNode("github.com/matthewmcnew/archtest/examples/transative").
		IsDependedOnBy("github.com/matthewmcnew/archtest/examples/dependency"),
	)
}

func TestPkgGraph_Roots(t *testing.T) {
	assert := assert.New(t)
	graph := NewPkgGraph()

	graph.Load("github.com/matthewmcnew/archtest/examples/dependency")

	assert.Len(graph.Roots(), 1)
}
