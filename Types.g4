// Calc.g4
grammar Types;


OPEN_PAREN: '(';
CLOSE_PAREN: ')';
OPEN_BRACKET: '[';
CLOSE_BRACKET: ']';
OPEN_BRACE: '{';
CLOSE_BRACE: '}';

ATOMIC_INT32: 'int32';
ATOMIC_UINT32: 'uint32';

IDENT: [a-zA-Z_][A-Za-z0-9_]+;
WHITESPACE: [ \r\n\t]+ -> skip;


// Rules
start : value=type EOF;

type: '[' elemType=type ']' #List 
    | '(' types=typeList ')' #Tuple
    | atom=atomic #AtomicType
    ;

typeList: pattern=type '...'  #TypeListOnlyExpansion
    | (types+=type ',')+ finalType=type #TypeListNoExpansion
    | (types+=type ',')+ pattern=type '...' #TypeListExpansion
    | #TypeListEmpty
    ;

atomic: name=ATOMIC_INT32 | name=ATOMIC_UINT32;