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
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
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

//ReturnValue ...
type ReturnValue struct {
	Value Object
}

//Type is of ReturnValue
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

//Inspect is of ReturnValue
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

//Error is
type Error struct {
	Message string
}

//Type is of Error
func (e *Error) Type() ObjectType { return ERROR_OBJ }

//Inspect is of Error
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
