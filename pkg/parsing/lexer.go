//line lexer.rl:1
package parsing

import (
	"fmt"
)

//line lexer.go:12
const chell_start int = 1
const chell_first_final int = 1
const chell_error int = 0

const chell_en_main int = 1

//line lexer.rl:26

type lexer struct {
	data        []byte
	p, pe, cs   int
	ts, te, act int
}

func newLexer(data []byte) *lexer {
	lex := &lexer{
		data: data,
		pe:   len(data),
	}

//line lexer.go:35
	{
		lex.cs = chell_start
		lex.ts = 0
		lex.te = 0
		lex.act = 0
	}

//line lexer.rl:40
	return lex
}

func (lex *lexer) Lex(out *yySymType) int {
	eof := lex.pe
	tok := 0

//line lexer.go:51
	{
		if (lex.p) == (lex.pe) {
			goto _test_eof
		}
		switch lex.cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		}
		goto st_out
	tr0:
//line lexer.rl:48
		lex.te = (lex.p) + 1

		goto st1
	tr2:
//line lexer.rl:50
		lex.te = (lex.p) + 1
		{
			tok = DQUOTE
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr3:
//line lexer.rl:52
		lex.te = (lex.p) + 1
		{
			tok = REF
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr4:
//line lexer.rl:49
		lex.te = (lex.p) + 1
		{
			tok = SQUOTE
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr6:
//line lexer.rl:51
		lex.te = (lex.p) + 1
		{
			tok = EOC
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr7:
//line lexer.rl:53
		lex.te = (lex.p) + 1
		{
			tok = LBRACE
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr8:
//line lexer.rl:54
		lex.te = (lex.p) + 1
		{
			tok = RBRACE
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	tr9:
//line lexer.rl:55
		lex.te = (lex.p)
		(lex.p)--
		{
			out.value = string(lex.data[lex.ts:lex.te])
			tok = STRING
			{
				(lex.p)++
				lex.cs = 1
				goto _out
			}
		}
		goto st1
	st1:
//line NONE:1
		lex.ts = 0

		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof1
		}
	st_case_1:
//line NONE:1
		lex.ts = (lex.p)

//line lexer.go:117
		switch lex.data[(lex.p)] {
		case 9:
			goto tr0
		case 32:
			goto tr0
		case 34:
			goto tr2
		case 36:
			goto tr3
		case 39:
			goto tr4
		case 59:
			goto tr6
		case 123:
			goto tr7
		case 125:
			goto tr8
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st2
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		lex.cs = 0
		goto _out
	st2:
		if (lex.p)++; (lex.p) == (lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch lex.data[(lex.p)] {
		case 45:
			goto st2
		case 95:
			goto st2
		}
		switch {
		case lex.data[(lex.p)] < 65:
			if 48 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 57 {
				goto st2
			}
		case lex.data[(lex.p)] > 90:
			if 97 <= lex.data[(lex.p)] && lex.data[(lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto tr9
	st_out:
	_test_eof1:
		lex.cs = 1
		goto _test_eof
	_test_eof2:
		lex.cs = 2
		goto _test_eof

	_test_eof:
		{
		}
		if (lex.p) == eof {
			switch lex.cs {
			case 2:
				goto tr9
			}
		}

	_out:
		{
		}
	}

//line lexer.rl:59

	return tok
}

func (lex *lexer) Error(e string) {
	fmt.Println("error:", e)
}
