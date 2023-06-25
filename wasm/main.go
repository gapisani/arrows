package main

import (
    "reflect"
    "syscall/js"

    "github.com/gapisani/arrows/core"
)

var (
    grid core.Grid
)

type CellType uint
const (
    None CellType = iota
    And
    Angled
    Block
    Cross
    Flash
    FrwdSide
    Get
    MemCell
    Not
    Source
    Wire
    Xor
    DoubleMemCell
    Random
    Double
    Unknown
)

// Enum
func cell2type(cell core.Cell) CellType {
    switch cell.(type) {
    case *core.None:
        return None
    case *core.And:
        return And
    case *core.Angled:
        return Angled
    case *core.Block:
        return Block
    case *core.Cross:
        return Cross
    case *core.Flash:
        return Flash
    case *core.FrwdSide:
        return FrwdSide
    case *core.Get:
        return Get
    case *core.MemCell:
        return MemCell
    case *core.Not:
        return Not
    case *core.Source:
        return Source
    case *core.Wire:
        return Wire
    case *core.Xor:
        return Xor
    case *core.DoubleMemCell:
        return DoubleMemCell
    case *core.Random:
        return Random
    case *core.Double:
        return Double
    default:
        return Unknown
    }
}

func type2cell(cellType CellType) core.Cell {
    switch cellType {
    case None:
        return &core.None{}
    case And:
        return &core.And{}
    case Angled:
        return &core.Angled{}
    case Block:
        return &core.Block{}
    case Cross:
        return &core.Cross{}
    case Flash:
        return &core.Flash{}
    case FrwdSide:
        return &core.FrwdSide{}
    case Get:
        return &core.Get{}
    case MemCell:
        return &core.MemCell{}
    case Not:
        return &core.Not{}
    case Source:
        return &core.Source{}
    case Wire:
        return &core.Wire{}
    case Xor:
        return &core.Xor{}
    case DoubleMemCell:
        return &core.DoubleMemCell{}
    case Random:
        return &core.Random{}
    case Double:
        return &core.Double{}
    default:
        return nil
    }
}

func cell2js(cell core.Cell) js.Value {
    return js.ValueOf(map[string]interface{}{
        "stringType": reflect.TypeOf(cell).String(),
        "powered": cell.Check(),
        "type": uint(cell2type(cell)),
        "dir": uint(cell.Dir()),
    })
}

func main() {
    // GridInit(int w, h)
    js.Global().Set("GridInit", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        grid.Init(uint(args[0].Int()), uint(args[1].Int()))
        return nil
    }))

    // GetCell(int x, y) Cell
    js.Global().Set("GetCell", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        cell := grid.GetCell(uint(args[0].Int()), uint(args[1].Int()))
        if(cell == nil) {
            return nil
        } else {
            return cell2js(*cell)
        }
    }))

    // SetCell(int x, y, cellType celltype, direction)
    js.Global().Set("SetCell", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        cell := grid.GetCell(uint(args[0].Int()), uint(args[1].Int()))
        (*cell) = type2cell(CellType(args[2].Int()))
        (*cell).SetDir(core.Direction(args[3].Int()))
        grid.RecountUpdate()
        return nil
    }))

    // RecountUpdate()
    js.Global().Set("RecountUpdate", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        grid.RecountUpdate()
        return nil
    }))

    // GridDimensions() [W, H]
    js.Global().Set("GridDimensions", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        w, h := grid.Dimensions()
        return []uint{w, h}
    }))

    // Update() [[x, y], ...]
    js.Global().Set("Update", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        points := grid.Update()
        jsArray := js.Global().Get("Array").New(len(points))
        for i := 0; i < len(points); i++ {
            subArray := js.Global().Get("Array").New(2)
            subArray.SetIndex(0, js.ValueOf(points[i].X))
            subArray.SetIndex(1, js.ValueOf(points[i].Y))
            jsArray.SetIndex(i, subArray)
        }
        return jsArray
    }))

    // LinkCells(x1, y1, x2, y2)
    // Links cell at pos1 to pos2
    js.Global().Set("LinkCells", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        cell := *grid.GetCell(uint(args[0].Int()), uint(args[1].Int()))
        *grid.GetCell(uint(args[0].Int()), uint(args[1].Int())) = cell
        *grid.GetCell(uint(args[2].Int()), uint(args[3].Int())) = cell
        return nil
    }))


    // Constants
    js.Global().Set("NONE", uint(None))
    js.Global().Set("AND", uint(And))
    js.Global().Set("ANGLED", uint(Angled))
    js.Global().Set("BLOCK", uint(Block))
    js.Global().Set("CROSS", uint(Cross))
    js.Global().Set("FLASH", uint(Flash))
    js.Global().Set("FRWD_SIDE", uint(FrwdSide))
    js.Global().Set("GET", uint(Get))
    js.Global().Set("MEM_CELL", uint(MemCell))
    js.Global().Set("NOT", uint(Not))
    js.Global().Set("SOURCE", uint(Source))
    js.Global().Set("WIRE", uint(Wire))
    js.Global().Set("XOR", uint(Xor))
    js.Global().Set("DOUBLE_MEMCELL", uint(DoubleMemCell))
    js.Global().Set("RANDOM", uint(Random))
    js.Global().Set("DOUBLE", uint(Double))
    js.Global().Set("UNKNOWN", uint(Unknown))

    js.Global().Set("NORTH", uint(core.NORTH))
    js.Global().Set("EAST",  uint(core.EAST))
    js.Global().Set("WEST",  uint(core.WEST))
    js.Global().Set("SOUTH", uint(core.SOUTH))

    // Don't quit from main
    ch := make(chan int)
    <-ch
}
