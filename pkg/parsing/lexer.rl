package parsing

import (
    "fmt"
    "strconv"
)

%%{ 
    machine chell;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;

	ALPHA = [a-zA-Z];
	ALPHANUM = [a-zA-Z0-9];

	s = [ \t];
	QUOTE_SINGLE = "'";
	QUOTE_DOUBLE = '"';
	EOC = ';';
	UNQUOTED_STRING = ALPHANUM (ALPHANUM | [\-_])*;
	REF = '$';
	LBRACE = '{';
	RBRACE = '}';
}%%

type lexer struct {
    data []byte
    p, pe, cs int
    ts, te, act int
}

func newLexer(data []byte) *lexer {
    lex := &lexer{ 
        data: data,
        pe: len(data),
    }
    %% write init;
    return lex
}

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0
    %%{ 
        main := |*
            s;
            QUOTE_SINGLE => { tok = SQUOTE; fbreak; };
            QUOTE_DOUBLE => { tok = DQUOTE; fbreak; };
            EOC => { tok = EOC; fbreak; };
            REF => { tok = REF; fbreak; };
            LBRACE => { tok = LBRACE; fbreak; };
            RBRACE => { tok = RBRACE; fbreak; };
            UNQUOTED_STRING => { out.value = string(lex.data[lex.ts:lex.te]); tok = STRING; fbreak; };
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("error:", e)
}
