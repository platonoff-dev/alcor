parser grammar AlcorParser;

options {
    tokenVocab = AlcorLexer;
}

numberLit : DECIMAL_LIT | BINARY_LIT | OCTAL_LIT | HEX_LIT;

expr : numberLit SEP (PLUS | MINUS | STAR | DIV) SEP numberLit;
prog : expr+;

