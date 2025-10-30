program : function^;

function : functionheader '{' '\n' functioncontent '}' '\n';

functionheader : 'function ' identifier '(' ')' ;

functioncontent : block+;

block : (statement+) | ifblock | whileblock | forblock ;

ifblock : 'if(' conditionalexpression '){\n' statement+ '}\n';

whileblock : 'while(' conditionalexpression '){\n' statement+ '}\n';

forblock : 'for(' forloop '){\n' statement+ '}\n';

forloop : declaration ';' conditionalexpression ';' incrementdecrement;

incrementdecrement : identifier ('++' | '--');

conditionalexpression : conditionalexpone (conditionaljoin conditionalexpone)*;

conditionalexpone : (identifier conditionaloperation operand);

conditionaloperation : '<' | '>' | '===' | '!==' | '<=' | '>=';

conditionaljoin : '&&' | '||';

statement : declaration | assignment | functioncall | returnstatement;

declaration : declarationtype ' ' identifier ' = ' expression ';\n';

assignment : identifier ' = ' expression ';\n';

functioncall : identifier '(' ')' ';\n';

returnstatement : 'return ' expression ';\n';

expression : operand (operator (operand | '(' expression ')'))*;

operand : identifier | integer | float | stringliteral | booleanliteral;

operator : '+' | '-' | '*' | '/' | '%';

declarationtype : 'let' | 'const' | 'var';

stringliteral : '"' [a-z]* '"';

booleanliteral : 'true' | 'false';

identifier : [a-z] [a-z0-9]*;

integer : [0-9]+;

float : [0-9]+ '.' [0-9]+;

whitespace : ' ';
