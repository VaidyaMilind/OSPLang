package object

import (
	"fmt"
)

//ObjectType of object
type ObjectType string

//Object Interface Type and Inspect
type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

//Integer struct with Value int64
type Integer struct {
	Value int64
}

//Inspect is of Integer
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

//Type is of Integer
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

//Boolean struct with bool
type Boolean struct {
	Value bool
}

//Type is of Integer
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

//Inspect is of Boolean
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

//Null struct with Empty
type Null struct{}

//Type is of Null
func (n *Null) Type() ObjectType { return NULL_OBJ }

//Inspect is of Null
func (n *Null) Inspect() string { return "null" }
