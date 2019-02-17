%{
package parsing

import (
    "fmt"
)
%}

%union {
   value string
   valuearr []string
}

%token <value> STRING
%type <value> argument reference innerqoute innerdqoute
%type <valuearr> arguments
%token EOC REF LBRACE RBRACE DQUOTE SQUOTE
%token STRING 

%%

block:
    /* empty */
    | commands
    ;

commands:
        commands EOC command
        | command
        ;

command:
        STRING arguments {
            exec_command($1, $2);
        };

arguments:
        /* empty */ { $$ = nil; }
        | arguments argument {
            $$ = append($$, $2);
        }
        | argument {
            $$ = append($$, $1);
        };

argument:
        innerdqoute {
            $$ = $1;
        }
        | SQUOTE innerqoute SQUOTE {
            $$ = $2;
        }
        | DQUOTE innerdqoute DQUOTE {
            $$ = $2;
        }
        ;

innerqoute:
        /* empty */ {
            $$ = "";
        }
        | STRING {
            $$ = $1;
        }
        ;

innerdqoute:
        /* empty */ {
            $$ = "";
        }
        | STRING {
            $$ = $1;
        }
        | innerqoute reference {
            $$ = $1 + $2
        }
        ;

reference:
        REF STRING {
            $$  = os.Getenv($2)
        }
        | REF LBRACE STRING RBRACE {
            $$  = os.Getenv($3)
        }
        ;
