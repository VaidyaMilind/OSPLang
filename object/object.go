package object

import (
	"OSPLang/ast"
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
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
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	HASH_OBJ         = "HASH"
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

/*
//NewEnvironment is ...
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

//Environment struct
type Environment struct {
	store map[string]Object
}

//Get for Environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

//Set for Environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

*/

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {

	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") { \n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }

func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}

	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type Hashable interface {
	HashKey() HashKey
}
