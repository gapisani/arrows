package main

import "C"

import (
	"github.com/gapisani/arrows/core"
)

type CellType uint
const (
    //export None
    None CellType = iota
    //export T_And
    And
    //export T_Andgled
    Angled
    //export T_Block
    Block
    //export T_Cross
    Cross
    //export T_Flash
    Flash
    //export T_FrwdLeft
    FrwdLeft
    //export T_FrwdRight
    FrwdRight
    //export T_Get
    Get
    //export T_MemCell
    MemCell
    //export T_Not
    Not
    //export T_Source
    Source
    //export T_Wire
    Wire
    //export T_Xor
    Xor
    //export T_Unknown
    Unknown
)

type Direction uint
const (
    //export D_North
    North Direction = iota
    //export D_East
    East
    //export D_South
    South
    //export D_West
    West
)

// Enum
func cell2type(cell core.Cell) CellType {
    switch cell.(type) {
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
    case *core.FrwdLeft:
        return FrwdLeft
    case *core.FrwdRight:
        return FrwdRight
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
    case Angled:
        return &core.Angled{}
    case Block:
        return &core.Block{}
    case Cross:
        return &core.Cross{}
    case Flash:
        return &core.Flash{}
    case FrwdLeft:
        return &core.FrwdLeft{}
    case FrwdRight:
        return &core.FrwdRight{}
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

var (
    g core.Grid
)

//export gridInit
func gridInit(w, h uint) {
    g.Init(w, h)
}

//export getCellDir
func getCellDir(x, y uint) Direction {
    cell := g.GetCell(x, y)
    if(cell == nil) {
        return 0
    }
    return Direction((*cell).Dir())
}

//export getCellLit
func getCellLit(x, y uint) bool {
    cell := g.GetCell(x, y)
    if(cell == nil) {
        return false
    }
    return (*cell).Check()
}

//export getCellType
func getCellType(x, y uint) bool {
    cell := g.GetCell(x, y)
    if(cell == nil) {
        return false
    }
    return (*cell).Check()
}
//export setCell
func setCell(x, y uint, t CellType, d Direction) {
    c := g.GetCell(x, y)
    (*c) = type2cell(t)
    (*c).SetDir(core.Direction(d))
}

//export Update
func Update() {
    g.Update()
}

// IDK why is it even exist
func main() {}
