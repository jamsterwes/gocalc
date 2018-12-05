// Code generated from .\Calculator.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Calculator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalculatorParser.
type CalculatorVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalculatorParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by CalculatorParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by CalculatorParser#mult_expr.
	VisitMult_expr(ctx *Mult_exprContext) interface{}

	// Visit a parse tree produced by CalculatorParser#pow_expr.
	VisitPow_expr(ctx *Pow_exprContext) interface{}

	// Visit a parse tree produced by CalculatorParser#sign_atom.
	VisitSign_atom(ctx *Sign_atomContext) interface{}

	// Visit a parse tree produced by CalculatorParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by CalculatorParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}
}
