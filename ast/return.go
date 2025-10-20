package ast

import (
	"monkey/token"
)

type ReturnStatement struct {
	Token       token.Token // the return statement
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
