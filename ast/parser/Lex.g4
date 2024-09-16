lexer grammar Lex;

// 忽略空白字符
// WS: [ \t\r\u000C]+ -> channel(HIDDEN);
WS: [ \t\r\u000C]+ -> skip;

NEW_LINE : [\r]? '\n' -> channel(HIDDEN);

LINE_COMMENT : '//' ~[\r\n]* -> channel(HIDDEN);

BLOCK_COMMENT : '/*' .*? '*/' -> channel(HIDDEN);

// 关键字
MODULE: 'module';

HASH: '#';

INCLUDE: 'include';

STRUCT: 'struct';

ENUM: 'enum';

OPTIONAL: 'optional';

REQUIRED: 'required';

VECTOR: 'vector';

MAP: 'map';

ULONG: 'unsigned long';

LONG: 'long';

FLOAT: 'float';

BYTE: 'byte';

UINT: 'unsigned int';

INT: 'int';

DOUBLE: 'double';

STRING: 'string';

VOID: 'void';

BOOL: 'bool';

TRUE: 'true';

FALSE: 'false';

INTERFACE: 'interface';

IN: 'in';

OUT: 'out';

LEFT_BRACE: '{';

RIGHT_BRACE: '}';

LEFT_BRACKET: '[';

RIGHT_BRACKET: ']';

LEFT_PAREN: '(';

RIGHT_PAREN: ')';

SEMICOLON: ';';

COLON: ':';

COMMA: ',';

EQUAL: '=';

LEFT_ANGLE_BRACKET: '<';

RIGHT_ANGLE_BRACKET: '>';

// 通用字符
STRING_LITERAL: '"' .+? '"';

NUMBER: [-]?'0x'?[0-9]+;

IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

DOUBLE_COLON: '::';