package interpreter

import (
    . "../parser"
    "github.com/antlr/antlr4/runtime/Go/antlr"
    "math"
    "strconv"
)


type CalculatorInterpreter struct {
    BaseCalculatorVisitor
}

func NodeExists(node interface{}) bool {
    switch t := node.(type) {
    case antlr.TerminalNode:
        return true
    default:
        _ = t
        return false
    }
}

func (ci *CalculatorInterpreter) VisitRoot(ctx *RootContext) float64 {
    expr, _ := ctx.Expr().(*ExprContext)
    return ci.VisitExpr(expr)
}

func (ci *CalculatorInterpreter) VisitExpr(ctx *ExprContext) float64 {
    left, _ := ctx.Mult_expr(0).(*Mult_exprContext)
    init := ci.VisitMult_expr(left)
    for i, mult_expr := range ctx.AllMult_expr()[1:] {
        expr, _ := mult_expr.(*Mult_exprContext)
        val := ci.VisitMult_expr(expr)
        if NodeExists(ctx.PLUS(i)) {
            init += val
        } else {
            init -= val
        }
    }
    return init
}

func (ci *CalculatorInterpreter) VisitMult_expr(ctx *Mult_exprContext) float64 {
    left, _ := ctx.Pow_expr(0).(*Pow_exprContext)
    init := ci.VisitPow_expr(left)
    for i, mult_expr := range ctx.AllPow_expr()[1:] {
        expr, _ := mult_expr.(*Pow_exprContext)
        val := ci.VisitPow_expr(expr)
        if NodeExists(ctx.TIMES(i)) {
            init *= val
        } else {
            init /= val
        }
    }
    return init
}

func (ci *CalculatorInterpreter) VisitPow_expr(ctx *Pow_exprContext) float64 {
    left, _ := ctx.Sign_atom(0).(*Sign_atomContext)
    init := ci.VisitSign_atom(left)
    for _, mult_expr := range ctx.AllSign_atom()[1:] {
        expr, _ := mult_expr.(*Sign_atomContext)
        val := ci.VisitSign_atom(expr)
        init = math.Pow(init, val)
    }
    return init
}

func (ci *CalculatorInterpreter) VisitSign_atom(ctx *Sign_atomContext) float64 {
    atom, _ := ctx.Atom().(*AtomContext)
    val := ci.VisitAtom(atom)
    if NodeExists(ctx.MINUS()) {
        return -1 * val
    } else {
        return val
    }
}

func (ci *CalculatorInterpreter) VisitAtom(ctx *AtomContext) float64 {
    switch t := ctx.Expr().(type) {
    case *ExprContext:
        expr, _ := ctx.Expr().(*ExprContext)
        return ci.VisitExpr(expr)
    default:
        _ = t
        constant, _ := ctx.Constant().(*ConstantContext)
        return ci.VisitConstant(constant)
    }
}

func (ci *CalculatorInterpreter) VisitConstant(ctx *ConstantContext) float64 {
    out, _ := strconv.ParseFloat(ctx.NUMBER().GetText(), 64)
    return out
}

func CalculatorInterpret(inputText string) float64 {
    input := antlr.NewInputStream(inputText)
    lexer := NewCalculatorLexer(input)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    calcParser := NewCalculatorParser(tokens)

    root, _ := calcParser.Root().(*RootContext)

    return new(CalculatorInterpreter).VisitRoot(root)
}
