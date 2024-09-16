parser grammar Sdp;

options { tokenVocab=Lex; }

sdp
    : includeStatement* moduleStatement* EOF
    ;

includeStatement
    : HASH INCLUDE STRING_LITERAL
    ;

moduleStatement
    : MODULE IDENTIFIER LEFT_BRACE (structStatement|serviceStatement|enumStatement)* RIGHT_BRACE SEMICOLON 
    ;

structStatement
    : STRUCT IDENTIFIER LEFT_BRACE fieldStatement* RIGHT_BRACE SEMICOLON
    ;

fieldStatement
    : NUMBER (OPTIONAL | REQUIRED) type IDENTIFIER (EQUAL (STRING_LITERAL | TRUE | FALSE | NUMBER))? SEMICOLON
    ;

serviceStatement
    : INTERFACE IDENTIFIER LEFT_BRACE serviceMethodStatement* RIGHT_BRACE SEMICOLON
    ;

serviceMethodStatement
    : type IDENTIFIER LEFT_PAREN methodParamStatement? RIGHT_PAREN SEMICOLON
    ;

methodParamStatement
    : (IN | OUT)? type IDENTIFIER (COMMA (IN | OUT)? type IDENTIFIER)*
    ;

enumStatement
    : ENUM IDENTIFIER LEFT_BRACE (enumField)* RIGHT_BRACE SEMICOLON
    ;

enumField
    : IDENTIFIER (EQUAL NUMBER)? COMMA
    ;

enumLastField
    : IDENTIFIER COMMA
    ;

type
    : UINT
    | ULONG
    | LONG
    | FLOAT
    | BYTE
    | INT
    | DOUBLE 
    | BOOL 
    | STRING
    | VOID 
    | IDENTIFIER
    | moduleType
    | vectorType
    | mapType 
    ;

vectorType
    : VECTOR LEFT_ANGLE_BRACKET type RIGHT_ANGLE_BRACKET
    ;

// map<int,string>
mapType
    : MAP LEFT_ANGLE_BRACKET type COMMA type RIGHT_ANGLE_BRACKET
    ;
moduleType
    : IDENTIFIER DOUBLE_COLON IDENTIFIER
    ;