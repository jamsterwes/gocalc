// Code generated from .\Calculator.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Calculator

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCalculatorVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalculatorVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitMult_expr(ctx *Mult_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitPow_expr(ctx *Pow_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitSign_atom(ctx *Sign_atomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}
