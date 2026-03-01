// Tafexpr.g4
grammar Tafexpr;

// Tokens
MUL: '*';
DIV: '/';
MOD: '%';
ADD: '+';
SUB: '-';
INTEGER: [0-9]+;
DOUBLE: [0]'.'[0-9]+|[1-9][0-9]*'.'[0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;
LBR : '[';
RBR : ']';
CON : '.';
NULL_TOKEN:'null';

//Logical Operators
LESSER_THAN : '<';
LESSER_THAN_EQUAL : '<=';
EQUAL:'==';
UNEQUAL:'!=';
GREATER_THAN: '>';
GREATER_THAN_EQUAL: '>=';

LOGICAL_AND : '&&';
LOGICAL_OR : '||';
LOGICAL_NOT : '!';




fragment LOWERCASE  :   [a-z];
fragment UPPERCASE  :   [A-Z];
fragment UNDERSCORE :   '_';
DOLLAR :   '$';  
fragment SINGLELETTER   :   ( LOWERCASE | UPPERCASE);
fragment VARIABLECON : '.';
fragment BOOL_TRUE:'true'|'TRUE';
fragment BOOL_FALSE:'false'|'FALSE';

/*STRING:   '"' (  ~[\\"]  )* '"';*/
STRING
  : UnterminatedStringLiteral '"'
  ;

UnterminatedStringLiteral
  : '"' (~["\\\r\n] | '\\' (. | EOF))*
  ;

/*fragment ESCAPED_QUOTE : '\\"';
STRING :   '"' ( ESCAPED_QUOTE | ~('\n'|'\r') )*? '"';*/

BOOLEAN : BOOL_TRUE|BOOL_FALSE;
NUMBER : INTEGER|DOUBLE;

VARIABLE_NAME : 
    DOLLAR SINGLELETTER (SINGLELETTER|INTEGER|UNDERSCORE)*
;

taf_expression: (libfunc | expression) EOF;
libfunc:  'randomDoubleInRange' '(' expression ',' expression ')' #HandleRandomDoubleInRange;

expression
   : 
       SUB expression #HandleNegation
       | LOGICAL_NOT expression #HandleLogicalNegation
   | expression op=(LOGICAL_AND|LOGICAL_OR) expression #HandleLogical
   | expression op=(MUL|DIV|MOD) expression #MulDiv
   | expression op=(ADD|SUB) expression #AddSub
   | expression op=(LESSER_THAN|LESSER_THAN_EQUAL|EQUAL|GREATER_THAN|GREATER_THAN_EQUAL|UNEQUAL) expression #LogicalOperation
   | expression '.' 'length' #HandleLength
   | expression '.' 'findOneByXPATH' '(' expression ')'#HandleFindOneByXPATH
   | expression '.' 'findOneStringByXPATH' '(' expression ')'#HandleFindOneStringByXPATH
   | expression '.' 'findOneDoubleByXPATH' '(' expression ')'#HandleFindOneDoubleByXPATH
   | expression '.' 'findOneIntegerByXPATH' '(' expression ')'#HandleFindOneIntegerByXPATH
   | expression '.' 'findOneBooleanByXPATH' '(' expression ')'#HandleFindOneBooleanByXPATH
   | expression '.' 'findByXPATH' '(' expression ')'#HandleFindByXPATH
   | expression '.' 'extractOneByREGEX' '(' expression ')'#HandleExtractOneByREGEX
   | expression '.' 'replaceAllStringOccurrences' '(' expression  ',' expression ')'#HandleReplaceAllStringOccurrences
   | expression '.' 'toString'#HandleToString
   | expression '.' 'toBoolean'#HandleToBoolean
   | expression '.' 'toInteger'#HandleToInteger
   | expression '.' 'toDouble'#HandleToDouble
      | expression '.' 'containsString' '(' expression ')'#HandleContainsString
   | expression '.' 'startsWith' '(' expression ')'#HandleStartsWith
   | expression '.' 'endsWith' '(' expression ')'#HandleEndsWith
   | expression '.' 'trimLeft'#HandleTrimLeft
   | libfunc #HandleLibfunc
   | expression '.' 'trimRight'#HandleTrimRight
   | expression '.' 'trim'#HandleTrim
   | INTEGER                             #Number
   | DOUBLE                              #DoubleValue
   | parenthesisExpression              #OrderedEvaluation
   | var_expression #HandleVarExpression
   | BOOLEAN #HandleBool 
   | NULL_TOKEN #HandleNull
   | STRING #HandleString
     | json #HandleJson
   ;
   


var_expression:  VARIABLE_NAME indx_expr;
indx_expr: ('[' index_expression ']' ('[' index_expression ']')*)? ('.' var_path)?;
var_path: jsonpath_expr ('.' jsonpath_expr)*;

jsonpath_expr: identifierWithQualifier
                 | PROP
                 ;

identifierWithQualifier :  PROP '[' index_expression ']' ('[' index_expression ']')*
                        ;

index_expression : expression #IndexExpression
;


PROP: SINGLELETTER (SINGLELETTER|INTEGER|UNDERSCORE)*;


parenthesisExpression: 
                      '(' (parenthesisExpression | expression) ')' 
                      ;
EXCEPT_QUOTE: ( ~[\\"] );

/*fragment ANYCHAR : .;
fragment COLON: ':';*/

any :  ':' | LBR | RBR | '{' | '}' | '"'  ;

/*string_interpolation: any*? match_interpolation any*? match_interpolation any*?;
interpolation:  string_interpolation+ EOF #HandleStringInterpolation;
match_interpolation: parenthesisExpression#HandleInterpolation;
*/

json
   : obj #HandleObject 
   | arr #HandleArray
   ;

obj
   : '{' pair (',' pair)* '}' #HandleObjectData
   | '{' '}' #HandleEmptyObjectData
   ;

pair
   : STRING ':' value #HandleObjectPair
   ;

arr
   : '[' value (',' value)* ']' 
   | '[' ']' 
   ;

/*value
   : STRING #HandleString
   | SUB DOUBLE #HandleNegation
   | SUB INTEGER #HandleNegation
   | DOUBLE #DoubleValue 
   | INTEGER #Number
   | var_expression #HandleVarExpression
   | obj  #HandleObject
   | arr #HandleArray
   | BOOLEAN #HandleBool
   | NULL_TOKEN  #HandleNull
   ;
*/
   value : 
    json #HandleJJ
|expression #HandleFoo;

   /*
   
      | INTEGER                             #Number
   | DOUBLE                              #DoubleValue
   | parenthesisExpression              #OrderedEvaluation
   | var_expression #HandleVarExpression
   | BOOLEAN #HandleBool 
   | NULL_TOKEN #HandleNull
   | STRING #HandleString
     | json #HandleJson
    */

/*
JSON_STRING
   : '"' (ESC | SAFECODEPOINT)* '"'
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;


fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;


fragment HEX
   : [0-9a-fA-F]
   ;


fragment SAFECODEPOINT
   : ~ ["\\\u0000-\u001F]
   ;

*/

JSON_NUMBER
   :  NUMBER
   ;


fragment INT
   // integer part forbis leading 0s (e.g. `01`)
   : '0' | [1-9] [0-9]*
   ;

// no leading zeros

fragment EXP
   // exponent number permits leading 0s (e.g. `1e01`)
   : [Ee] [+\-]? [0-9]+
   ;

// \- since - means "range" inside [...]
WS
   : [ \t\n\r] + -> skip
   ;

   // Catch-all for unknown tokens
UNKNOWN: . ;