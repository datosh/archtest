package archtest

import (
	"container/list"
	"fmt"
	"strings"

	"golang.org/x/tools/go/packages"
)

type PackageTest struct {
	packages     []string
	ignored      map[string]interface{}
	t            TestingT
	includeTests bool
}

type TestingT interface {
	Error(args ...interface{})
}

func Package(t TestingT, packageName ...string) *PackageTest {
	return &PackageTest{packages: packageName, t: t, includeTests: false}
}

func (t PackageTest) IncludeTests() *PackageTest {
	t.includeTests = true
	return &t
}

func (t PackageTest) Ignoring(e ...string) *PackageTest {
	set := make(map[string]interface{})

	for v := range t.ignored {
		set[v] = struct{}{}
	}

	for _, v := range t.expand(e) {
		set[v] = struct{}{}
	}

	t.ignored = set
	return &t
}

func (t PackageTest) ShouldNotDependDirectlyOn(pkgs ...string) {
	t.shouldNotDependOnPackageWithFilter(func(d *dep) bool {
		return d.depth() > 1
	}, pkgs)
}

func (t *PackageTest) ShouldNotDependOn(pkgs ...string) {
	t.shouldNotDependOnPackageWithFilter(noOpFilter, pkgs)
}

func (t *PackageTest) shouldNotDependOnPackageWithFilter(filter depFilter, d []string) {
	dl := t.expand(d)
	for i := range t.findDeps(t.packages, filter) {
		if i.isDependencyOn(dl) {
			chain, _ := i.chain()
			msg := fmt.Sprintf("Error:\n%s", chain)
			t.t.Error(msg)
		}
	}
}

type depFilter func(*dep) bool

var noOpFilter depFilter = func(i *dep) bool {
	return false
}

type dep struct {
	name   string
	parent *dep
	xtest  bool
}

func (d *dep) depth() int {
	if d.parent == nil {
		return 0
	}
	return d.parent.depth() + 1
}

func (d *dep) chain() (string, int) {
	name := d.name
	if d.xtest {
		name = d.name + "_test"
	}

	if d.parent == nil {
		return name + "\n", 1
	}

	c, tabs := d.parent.chain()

	return c + strings.Repeat("\t", tabs) + name + "\n", tabs + 1
}

func (d dep) asxtest() *dep {
	d.xtest = true
	return &d
}

func (d *dep) isDependencyOn(dl []string) bool {
	if d.parent == nil {
		return false
	}

	if contains(dl, d.name) {
		return true
	}
	return false
}

func (t PackageTest) findDeps(packages []string, filter depFilter) <-chan *dep {
	c := make(chan *dep)
	go func() {
		defer close(c)

		importCache := map[string]struct{}{}
		for _, p := range t.expand(packages) {
			t.read(c, &dep{name: p, parent: nil}, importCache, filter)
		}
	}()
	return c
}

func (t *PackageTest) read(pChan chan *dep, d *dep, cache map[string]struct{}, filter depFilter) {
	queue := list.New()

	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedImports,
		Tests:      t.includeTests,
		BuildFlags: []string{},
	}

	queue.PushBack(d)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		d, _ := (front.Value).(*dep)

		if t.skip(cache, d.name) || filter(d) {
			continue
		}

		cache[d.name] = struct{}{}
		pChan <- d

		// pkg, err := context.Import(d.name, ".", importMode)
		pkgs, err := packages.Load(cfg, d.name)
		for _, pkg := range pkgs {
			if err != nil {
				e := fmt.Sprintf("Error reading: %s", d.name)
				t.t.Error(e)

				continue
			}

			for _, imported := range pkg.Imports {
				queue.PushBack(&dep{name: imported.PkgPath, parent: d})
			}
		}
	}
}

func (t PackageTest) expand(ps []string) []string {
	if !needExpansion(ps) {
		return ps
	}

	cfg := &packages.Config{
		Mode:       packages.NeedName,
		Tests:      false,
		BuildFlags: []string{},
	}

	loadedPs, err := packages.Load(cfg, ps...)
	if err != nil {
		e := fmt.Sprintf("Error reading: %s, err: %s", ps, err)
		t.t.Error(e)
		return nil
	}
	if len(loadedPs) == 0 {
		e := fmt.Sprintf("Error reading: %s, did not match any packages", ps)
		t.t.Error(e)
		return nil

	}

	ls := make([]string, 0, len(loadedPs))

	for _, p := range loadedPs {
		ls = append(ls, p.PkgPath)
	}

	return ls
}

func (t PackageTest) skip(cache map[string]struct{}, pkg string) bool {
	if _, excluded := t.ignored[pkg]; excluded ||
		pkg == "C" {
		return true
	}

	_, seen := cache[pkg]
	return seen
}

func needExpansion(ps []string) bool {
	for _, p := range ps {
		if strings.Contains(p, "...") {
			return true
		}
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
