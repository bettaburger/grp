package shell

/*
command read: ls filename

lexing => first split into tokens
1) [ls] [filename]

Parser will then read each token, if token is not recognized or wrong, parser will throw an error before running.
1) ls -> command
2) filename -> target 
*/

