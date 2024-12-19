lexer grammar AlcorLexer;

// Tokens
PLUS: '+';
MINUS: '-';
STAR: '*';
DIV: '/';


// Number literals
DECIMAL_LIT : ('0' | [1-9] ('_'? [0-9])*);
BINARY_LIT : '0' [bB] ('_'? [01])+;
OCTAL_LIT : '0' [oO] ('_'? [0-7])+;
HEX_LIT : '0' [xX] ('_'? [0-9a-fA-F])+;

SEP : ' ';

fragment DECIMALS: [0-9] ('_'? [0-9])*;
fragment OCTAL_DIGIT: [0-7];
fragment HEX_DIGIT: [0-9a-fA-F];
fragment BIN_DIGIT: [01];
fragment EXPONENT: [eE] [+-]? DECIMALS;


