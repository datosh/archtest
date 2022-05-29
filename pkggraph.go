package archtest

import "golang.org/x/tools/go/packages"

type PkgNode struct {
	dependsOn    map[string]*PkgNode
	dependedOnBy map[string]*PkgNode
	pkgPath      string
}

func NewPkgNode(pkgPath string) *PkgNode {
	n := &PkgNode{
		dependsOn:    make(map[string]*PkgNode),
		dependedOnBy: make(map[string]*PkgNode),
		pkgPath:      pkgPath,
	}
	return n
}

func (p *PkgNode) DependOn(other *PkgNode) {
	p.dependsOn[other.pkgPath] = other
	other.dependedOnBy[p.pkgPath] = p
}

func (p *PkgNode) IsDependingOn(pkgName string) bool {
	_, exists := p.dependsOn[pkgName]
	return exists
}

func (p *PkgNode) IsDependedOnBy(pkgName string) bool {
	_, exists := p.dependedOnBy[pkgName]
	return exists
}

// func (p *PkgNode) DependedOnBy(other *PkgNode) {
// 	p.dependedOnBy[other.pkgPath] = other
// 	other.dependsOn[p.pkgPath] = p
// }

type PkgGraph struct {
	lut   map[string]*PkgNode
	roots []*PkgNode
}

// TODO: Directly load varadic number of package strings
func NewPkgGraph() *PkgGraph {
	g := &PkgGraph{}
	g.lut = make(map[string]*PkgNode)
	return g
}

func (p *PkgGraph) addNode(pkgPath string) {
	if _, exists := p.lut[pkgPath]; !exists {
		p.lut[pkgPath] = NewPkgNode(pkgPath)
	}
}

func (p *PkgGraph) AddNode(pkgPath string) {
	p.addNode(pkgPath)
	p.updateRoots()
}

func (p *PkgGraph) GetNode(pkgPath string) *PkgNode {
	if node, exists := p.lut[pkgPath]; exists {
		return node
	} else {
		return nil
	}
}

func (p *PkgGraph) Size() int {
	return len(p.lut)
}

func (p *PkgGraph) Load(pkg string) error {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedImports,
	}

	pkgs, err := packages.Load(cfg, pkg)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		p.addNode(pkg.PkgPath)
		for _, dependsOn := range pkg.Imports {
			p.addNode(dependsOn.PkgPath)
			p.GetNode(pkg.PkgPath).DependOn(p.GetNode(dependsOn.PkgPath))
		}
	}

	p.updateRoots()
	return nil
}

func (p *PkgGraph) Roots() []*PkgNode {
	return p.roots
}

func (p *PkgGraph) updateRoots() {
	p.roots = make([]*PkgNode, 0)

	for _, node := range p.lut {
		if len(node.dependedOnBy) == 0 {
			p.roots = append(p.roots, node)
		}
	}
}
