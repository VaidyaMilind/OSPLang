package ast

import "OSPLang/token"

//Node is ...
type Node interface {
	TokenLiteral() string
}

//Statement is ...
type Statement interface {
	Node
	statementNode()
}

//Expression is ...
type Expression interface {
	Node
	expressionNode()
}

//Program is ...
type Program struct {
	Statements []Statement
}

//TokenLiteral is ...
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//LetStatement is ...
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

//TokenLiteral is ...
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

//Identifier is ...
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

//TokenLiteral is ...
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

//ReturnStatement is ...
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

//TokenLiteral is ...
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
