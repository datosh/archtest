package archtest

// import (
// 	"strings"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestPackage_ShouldNotDependOn(t *testing.T) {
// 	t.Run("Succeeds on non dependencies", func(t *testing.T) {
// 		mockT := new(testingT)
// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/nodependency")

// 		assertNoError(t, mockT)
// 	})

// 	t.Run("Fails on dependencies", func(t *testing.T) {
// 		mockT := new(testingT)
// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/dependency")

// 		assertError(t, mockT,
// 			"github.com/matthewmcnew/archtest/examples/testpackage",
// 			"github.com/matthewmcnew/archtest/examples/dependency")
// 	})

// 	t.Run("Supports testing against packages in the go root", func(t *testing.T) {
// 		mockT := new(testingT)
// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage").
// 			ShouldNotDependOn("crypto")

// 		assertError(t, mockT,
// 			"github.com/matthewmcnew/archtest/examples/testpackage",
// 			"crypto")
// 	})

// 	t.Run("Fails on transative dependencies", func(t *testing.T) {
// 		mockT := new(testingT)
// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/transative")

// 		assertError(t, mockT,
// 			"github.com/matthewmcnew/archtest/examples/testpackage",
// 			"github.com/matthewmcnew/archtest/examples/dependency",
// 			"github.com/matthewmcnew/archtest/examples/transative")
// 	})

// 	t.Run("Supports multiple packages at once", func(t *testing.T) {
// 		mockT := new(testingT)
// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/dontdependonanything", "github.com/matthewmcnew/archtest/examples/testpackage").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/nodependency", "github.com/matthewmcnew/archtest/examples/dependency")

// 		assertError(t, mockT,
// 			"github.com/matthewmcnew/archtest/examples/testpackage",
// 			"github.com/matthewmcnew/archtest/examples/dependency")
// 	})

// 	t.Run("Supports wildcard matching", func(t *testing.T) {
// 		mockT := new(testingT)
// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/...").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/nodependency")

// 		assertNoError(t, mockT)

// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage/nested/...").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/...")

// 		assertError(t, mockT, "github.com/matthewmcnew/archtest/examples/testpackage/nested/dep", "github.com/matthewmcnew/archtest/examples/nesteddependency")
// 	})

// 	t.Run("Supports checking imports in test files", func(t *testing.T) {
// 		mockT := new(testingT)

// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage/...").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/testfiledeps/testonlydependency")

// 		assertNoError(t, mockT)

// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage/...").
// 			IncludeTests().
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/testfiledeps/testonlydependency")

// 		assertError(t, mockT,
// 			"github.com/matthewmcnew/archtest/examples/testpackage/nested/dep",
// 			"github.com/matthewmcnew/archtest/examples/testfiledeps/testonlydependency",
// 		)
// 	})

// 	// t.Run("Supports checking imports from test packages", func(t *testing.T) {
// 	// 	mockT := new(testingT)

// 	// 	Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage/...").
// 	// 		ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/testfiledeps/testpkgdependency")

// 	// 	assertNoError(t, mockT)

// 	// 	Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage/...").
// 	// 		IncludeTests().
// 	// 		ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/testfiledeps/testpkgdependency")

// 	// 	assertError(t, mockT,
// 	// 		"github.com/matthewmcnew/archtest/examples/testpackage/nested/dep_test",
// 	// 		"github.com/matthewmcnew/archtest/examples/testfiledeps/testpkgdependency",
// 	// 	)
// 	// })

// 	t.Run("Supports Ignoring packages", func(t *testing.T) {
// 		mockT := new(testingT)

// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage/nested/dep").
// 			Ignoring("github.com/matthewmcnew/archtest/examples/testpackage/nested/dep").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/nesteddependency")

// 		assertNoError(t, mockT)
// 	})

// 	t.Run("Ignored packages ignore ignored transitive packages", func(t *testing.T) {
// 		mockT := new(testingT)

// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage").
// 			Ignoring("github.com/this/is/verifying/multiple/exclusions", "github.com/matthewmcnew/archtest/examples/...").
// 			Ignoring("github.com/this/is/verifying/chaining").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/transative")

// 		assertNoError(t, mockT)
// 	})

// 	// t.Run("Fails on packages that do not exist", func(t *testing.T) {
// 	// 	mockT := new(testingT)
// 	// 	Package(mockT, "github.com/matthewmcnew/archtest/dontexist/sorry").
// 	// 		ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/dependency")

// 	// 	assertError(t, mockT)

// 	// 	mockT = new(testingT)
// 	// 	Package(mockT, "DONT__WORK").
// 	// 		ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/dependency")

// 	// 	assertError(t, mockT)

// 	// 	mockT = new(testingT)
// 	// 	Package(mockT, "github.com/matthewmcnew/archtest/dontexist/...").
// 	// 		ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/dependency")

// 	// 	assertError(t, mockT)
// 	// })
// }

// func TestPackage_ShouldNotDependDirectly(t *testing.T) {
// 	t.Run("Fails on direct dependencies", func(t *testing.T) {
// 		mockT := new(testingT)
// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage").
// 			ShouldNotDependDirectlyOn("github.com/matthewmcnew/archtest/examples/dependency")

// 		assertError(t, mockT,
// 			"github.com/matthewmcnew/archtest/examples/testpackage",
// 			"github.com/matthewmcnew/archtest/examples/dependency")
// 	})

// 	t.Run("Fails on transative dependencies", func(t *testing.T) {
// 		mockT := new(testingT)
// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage").
// 			ShouldNotDependDirectlyOn("github.com/matthewmcnew/archtest/examples/transative")

// 		assertNoError(t, mockT)
// 	})
// }

// func assertNoError(t *testing.T, mockT *testingT) {
// 	t.Helper()
// 	if mockT.errored() {
// 		t.Fatalf("archtest should not have failed but, %s", mockT.message())
// 	}
// }

// func assertError(t *testing.T, mockT *testingT, dependencyTrace ...string) {
// 	t.Helper()
// 	if !mockT.errored() {
// 		t.Fatal("archtest did not fail on dependency")
// 	}

// 	if dependencyTrace == nil {
// 		return
// 	}

// 	s := strings.Builder{}
// 	s.WriteString("Error:\n")
// 	for i, v := range dependencyTrace {
// 		s.WriteString(strings.Repeat("\t", i))
// 		s.WriteString(v + "\n")
// 	}

// 	if mockT.message() != s.String() {
// 		t.Errorf("expected %s got error message: %s", s.String(), mockT.message())
// 	}
// }

// type testingT struct {
// 	errors [][]interface{}
// }

// func (t *testingT) Error(args ...interface{}) {
// 	t.errors = append(t.errors, args)
// }

// func (t testingT) errored() bool {
// 	if len(t.errors) != 0 {
// 		return true
// 	}

// 	return false
// }

// func (t *testingT) message() interface{} {
// 	return t.errors[0][0]
// }

// func TestPackageTest_expand(t *testing.T) {
// 	pt := PackageTest{}

// 	expanded := pt.expand([]string{"github.com/matthewmcnew/archtest/..."})

// 	assert.Equal(t, expanded, []string{
// 		"github.com/matthewmcnew/archtest",
// 		"github.com/matthewmcnew/archtest/examples/dependency",
// 		"github.com/matthewmcnew/archtest/examples/dontdependonanything",
// 		"github.com/matthewmcnew/archtest/examples/nesteddependency",
// 		"github.com/matthewmcnew/archtest/examples/nodependency",
// 		"github.com/matthewmcnew/archtest/examples/testfiledeps/testonlydependency",
// 		"github.com/matthewmcnew/archtest/examples/testfiledeps/testpkgdependency",
// 		"github.com/matthewmcnew/archtest/examples/testpackage",
// 		"github.com/matthewmcnew/archtest/examples/testpackage/nested/dep",
// 		"github.com/matthewmcnew/archtest/examples/transative",
// 	})
// }

// func BenchmarkCheckingImportsInTestFiles(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		mockT := new(testingT)

// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage/...").
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/testfiledeps/testonlydependency")

// 		// assertNoError(t, mockT)

// 		Package(mockT, "github.com/matthewmcnew/archtest/examples/testpackage/...").
// 			IncludeTests().
// 			ShouldNotDependOn("github.com/matthewmcnew/archtest/examples/testfiledeps/testonlydependency")

// 		// assertError(t, mockT,
// 		// 	"github.com/matthewmcnew/archtest/examples/testpackage/nested/dep",
// 		// 	"github.com/matthewmcnew/archtest/examples/testfiledeps/testonlydependency",
// 		// )
// 	}
// }
