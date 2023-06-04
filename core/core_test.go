package core_test

import (
	"github.com/gapisani/arrows/core"
	"testing"
)

func TestMain(t *testing.T) {
    g := core.Grid{}
    g.Init(10, 10)
    *g.GetCell(5, 5) = core.Cell(&core.Source{})
    *g.GetCell(6, 5) = core.Cell(&core.Wire{})
    g.Update()
}
