package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.statements) > 0 {
        return p.statements[0].TokenLiteral()
    }
    else {
        return ""
    }
}

// struct for "let"
type LetStatement struct {
    Token token.Token
    Name  *Identifier
    Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
    return ls.Token.Literal
}

// struct for "IDENT"
type Identifier struct {
    Token token.Token
    Value string
}

func (i *Identifier) expressionNode() {}

func (i *Idnetifier) TokenLiteral() string {
    return i.Token.Literal
}
