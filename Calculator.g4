grammar Calculator;

root
    : expr
    ;

expr
    : mult_expr ((PLUS | MINUS) mult_expr)*
    ;

mult_expr
    : pow_expr ((TIMES | DIV) pow_expr)*
    ;

pow_expr
    : sign_atom (POW sign_atom)*
    ;

sign_atom
    : MINUS atom
    | atom
    ;

atom
    : LPAREN expr RPAREN
    | constant
    ;

constant
    : NUMBER
    ;

LPAREN
   : '('
   ;

RPAREN
   : ')'
   ;

PLUS
   : '+'
   ;

MINUS
   : '-'
   ;


TIMES
   : '*'
   ;

DIV
   : '/'
   ;

POW
   : '^'
   ;

NUMBER
    : ('0' .. '9')+ ('.' ('0' .. '9')+)?
    ;

WS
    : [ \r\n\t]+ -> skip
    ;
