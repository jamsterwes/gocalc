# gocalc
> Simple calculator written in golang with ANTLR4
### Files:
+ `interpreter/interpreter.go` - Implementation of an ANTLR4 visitor that turns an AST into a `float64` result.  
+ `parser/*.go` - ANTLR4-generated files  
+ `Calculator.g4` - ANTLR4 grammar definition file  
+ `main.go` - Simple test of interpreter  

### Todo:
- [x] Interpret expressions correctly
- [ ] Generate bitcode (via LLVM) that outputs correct result
- [ ] Link bitcode into EXE file
