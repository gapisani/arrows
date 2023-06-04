package main

import (
	"reflect"
	"syscall/js"

	"github.com/gapisani/arrows/core"
)

var (
    grid core.Grid
    width, height uint
)

type CellType uint
const (
    None CellType = iota
    And
    Block
    Flash
    Get
    MemCell
    Not
    Source
    Wire
    Xor
    Unknown
)

// Enum
func cell2type(cell core.Cell) CellType {
    switch cell.(type) {
    case *core.And:
        return And
    case *core.Block:
        return Block
    case *core.Flash:
        return Flash
    case *core.Get:
        return Get
    case *core.MemCell:
        return MemCell
    case *core.None:
        return None
    case *core.Not:
        return Not
    case *core.Source:
        return Source
    case *core.Wire:
        return Wire
    case *core.Xor:
        return Xor
    default:
        return Unknown
    }
}

func type2cell(cellType CellType) core.Cell {
    switch cellType {
    case And:
        return &core.And{}
    case Block:
        return &core.Block{}
    case Flash:
        return &core.Flash{}
    case Get:
        return &core.Get{}
    case MemCell:
        return &core.MemCell{}
    case None:
        return &core.None{}
    case Not:
        return &core.Not{}
    case Source:
        return &core.Source{}
    case Wire:
        return &core.Wire{}
    case Xor:
        return &core.Xor{}
    default:
        return nil
    }
}


func cell2js(cell core.Cell) js.Value {
    return js.ValueOf(map[string]interface{}{
        "stringType": reflect.TypeOf(cell).String(),
        "powered": cell.Check(),
        "type": uint(cell2type(cell)),
    })
}

func main() {
    // gridInit(int w, h)
    js.Global().Set("gridInit", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        grid.Init(uint(args[0].Int()), uint(args[1].Int()))
        return nil
    }))

    // getCell(int x, y) Cell
    js.Global().Set("getCell", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        cell := grid.GetCell(uint(args[0].Int()), uint(args[1].Int()))
        return cell2js(*cell)
    }))

    // setCell(int x, y, cellType celltype)
    js.Global().Set("setCell", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        cell := grid.GetCell(uint(args[0].Int()), uint(args[1].Int()))
        (*cell) = type2cell(CellType(args[2].Int()))
        return nil
    }))

    // gridDimensions() [W, H]
    js.Global().Set("Dimensions", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        return []uint{width, height}
    }))


    // Don't quit from main
    ch := make(chan int)
    <-ch
}
