
//line src/coffeephp/parser.y:2
package coffeephp
import __yyfmt__ "fmt"
//line src/coffeephp/parser.y:2
		
// vim:et:sw=4:ai:si:ts=4:sts=4:

import _ "fmt"
import "coffeephp/ast"

var regs = make([]int, 26)
var base int


//line src/coffeephp/parser.y:16
type CoffeeSymType struct{
	yys int
    typ TokenType
    val ast.Node
    line int
    pos  int
}

const T_DIGIT = 57346
const T_LETTER = 57347
const T_DOT = 57348
const T_IDENTIFIER = 57349
const T_EOF = 57350
const T_FLOATING = 57351
const T_NUMBER = 57352
const T_STRING = 57353
const T_ONELINE_COMMENT = 57354
const T_COMMENT = 57355
const T_NEWLINE = 57356
const T_ASSIGN = 57357
const T_NEW = 57358
const T_CLONE = 57359
const T_IF = 57360
const T_ELSEIF = 57361
const T_ELSE = 57362
const T_FOR = 57363
const T_SAY = 57364
const T_SPACE = 57365
const T_ECHO = 57366
const T_FOREACH = 57367
const T_TRY = 57368
const T_CATCH = 57369
const T_CLASS = 57370
const T_IS = 57371
const T_DOES = 57372
const T_FUNCTION_PROTOTYPE = 57373
const T_RANGE_OPERATOR = 57374
const T_BRACKET_OPEN = 57375
const T_BRACKET_CLOSE = 57376
const T_CONST = 57377
const T_RETURN = 57378
const T_BREAK = 57379
const T_CONTINUE = 57380
const T_THROW = 57381
const T_NS_SEPARATOR = 57382
const T_NAMESPACE = 57383
const T_OBJECT_OPERATOR = 57384
const T_BOOLEAN_OR = 57385
const T_BOOLEAN_AND = 57386
const UMINUS = 57387

var CoffeeToknames = []string{
	"T_DIGIT",
	"T_LETTER",
	"T_DOT",
	"T_IDENTIFIER",
	"T_EOF",
	"T_FLOATING",
	"T_NUMBER",
	"T_STRING",
	"T_ONELINE_COMMENT",
	"T_COMMENT",
	"T_NEWLINE",
	"T_ASSIGN",
	"T_NEW",
	"T_CLONE",
	"T_IF",
	"T_ELSEIF",
	"T_ELSE",
	"T_FOR",
	"T_SAY",
	"T_SPACE",
	"T_ECHO",
	"T_FOREACH",
	"T_TRY",
	"T_CATCH",
	"T_CLASS",
	"T_IS",
	"T_DOES",
	"T_FUNCTION_PROTOTYPE",
	"T_RANGE_OPERATOR",
	"T_BRACKET_OPEN",
	"T_BRACKET_CLOSE",
	"T_CONST",
	"T_RETURN",
	"T_BREAK",
	"T_CONTINUE",
	"T_THROW",
	"T_NS_SEPARATOR",
	"T_NAMESPACE",
	"T_OBJECT_OPERATOR",
	" . (T_OBJECT_OPERATOR)",
	" and",
	" or",
	" |",
	" ^",
	" &",
	" +",
	" -",
	" *",
	" /",
	" %",
	" !",
	"T_BOOLEAN_OR",
	"T_BOOLEAN_AND",
	"UMINUS",
}
var CoffeeStatenames = []string{}

const CoffeeEofCode = 1
const CoffeeErrCode = 2
const CoffeeMaxDepth = 200

//line src/coffeephp/parser.y:192
      /*  start  of  programs  */

//line yacctab:1
var CoffeeExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 2,
}

const CoffeeNprod = 28
const CoffeePrivate = 57344

var CoffeeTokenNames []string
var CoffeeStates []string

const CoffeeLast = 78

var CoffeeAct = []int{

	21, 41, 20, 15, 16, 17, 18, 19, 35, 24,
	9, 13, 13, 13, 21, 36, 20, 15, 16, 17,
	18, 19, 40, 22, 38, 21, 37, 20, 15, 16,
	17, 18, 19, 17, 18, 19, 39, 14, 20, 15,
	16, 17, 18, 19, 15, 16, 17, 18, 19, 34,
	7, 11, 11, 11, 8, 4, 3, 2, 1, 6,
	5, 10, 10, 10, 12, 23, 25, 0, 0, 0,
	26, 27, 28, 29, 30, 31, 32, 33,
}
var CoffeePact = []int{

	-1000, -1000, -1000, 3, -1000, -1000, -1000, -1000, -21, -36,
	2, 2, -1000, -1000, -1000, 2, 2, 2, 2, 2,
	2, 2, 1, -46, -1000, -1000, -18, -18, -1000, -1000,
	-1000, -5, -10, -32, -34, -38, -1000, -1000, -1000, -1000,
	-60, -1000,
}
var CoffeePgo = []int{

	0, 54, 64, 60, 59, 58, 57, 56, 55, 50,
	49, 49, 49, 49, 36,
}
var CoffeeR1 = []int{

	0, 5, 7, 6, 6, 8, 3, 3, 4, 9,
	9, 11, 12, 13, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 14, 10,
}
var CoffeeR2 = []int{

	0, 1, 0, 3, 0, 1, 1, 1, 2, 4,
	4, 2, 5, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1, 1, 1, 2, 2,
}
var CoffeeChk = []int{

	-1000, -5, -6, -7, -8, -3, -4, -9, -1, 7,
	60, 50, -2, 10, 58, 49, 50, 51, 52, 53,
	48, 46, 59, -1, 7, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -10, 7, 61, 58, 58, -14,
	60, 61,
}
var CoffeeDef = []int{

	4, -2, -2, 0, 3, 5, 6, 7, 0, 23,
	0, 0, 24, 25, 8, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 23, 22, 15, 16, 17, 18,
	19, 20, 21, 0, 0, 23, 14, 9, 10, 27,
	0, 26,
}
var CoffeeTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 54, 3, 3, 3, 53, 48, 3,
	60, 61, 51, 49, 3, 62, 43, 52, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 58,
	3, 59, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 47, 3, 3, 44, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 45, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 46,
}
var CoffeeTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 55, 56, 57,
}
var CoffeeTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var CoffeeDebug = 0

type CoffeeLexer interface {
	Lex(lval *CoffeeSymType) int
	Error(s string)
}

const CoffeeFlag = -1000

func CoffeeTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(CoffeeToknames) {
		if CoffeeToknames[c-4] != "" {
			return CoffeeToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func CoffeeStatname(s int) string {
	if s >= 0 && s < len(CoffeeStatenames) {
		if CoffeeStatenames[s] != "" {
			return CoffeeStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func Coffeelex1(lex CoffeeLexer, lval *CoffeeSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = CoffeeTok1[0]
		goto out
	}
	if char < len(CoffeeTok1) {
		c = CoffeeTok1[char]
		goto out
	}
	if char >= CoffeePrivate {
		if char < CoffeePrivate+len(CoffeeTok2) {
			c = CoffeeTok2[char-CoffeePrivate]
			goto out
		}
	}
	for i := 0; i < len(CoffeeTok3); i += 2 {
		c = CoffeeTok3[i+0]
		if c == char {
			c = CoffeeTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = CoffeeTok2[1] /* unknown char */
	}
	if CoffeeDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", CoffeeTokname(c), uint(char))
	}
	return c
}

func CoffeeParse(Coffeelex CoffeeLexer) int {
	var Coffeen int
	var Coffeelval CoffeeSymType
	var CoffeeVAL CoffeeSymType
	CoffeeS := make([]CoffeeSymType, CoffeeMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Coffeestate := 0
	Coffeechar := -1
	Coffeep := -1
	goto Coffeestack

ret0:
	return 0

ret1:
	return 1

Coffeestack:
	/* put a state and value onto the stack */
	if CoffeeDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", CoffeeTokname(Coffeechar), CoffeeStatname(Coffeestate))
	}

	Coffeep++
	if Coffeep >= len(CoffeeS) {
		nyys := make([]CoffeeSymType, len(CoffeeS)*2)
		copy(nyys, CoffeeS)
		CoffeeS = nyys
	}
	CoffeeS[Coffeep] = CoffeeVAL
	CoffeeS[Coffeep].yys = Coffeestate

Coffeenewstate:
	Coffeen = CoffeePact[Coffeestate]
	if Coffeen <= CoffeeFlag {
		goto Coffeedefault /* simple state */
	}
	if Coffeechar < 0 {
		Coffeechar = Coffeelex1(Coffeelex, &Coffeelval)
	}
	Coffeen += Coffeechar
	if Coffeen < 0 || Coffeen >= CoffeeLast {
		goto Coffeedefault
	}
	Coffeen = CoffeeAct[Coffeen]
	if CoffeeChk[Coffeen] == Coffeechar { /* valid shift */
		Coffeechar = -1
		CoffeeVAL = Coffeelval
		Coffeestate = Coffeen
		if Errflag > 0 {
			Errflag--
		}
		goto Coffeestack
	}

Coffeedefault:
	/* default state action */
	Coffeen = CoffeeDef[Coffeestate]
	if Coffeen == -2 {
		if Coffeechar < 0 {
			Coffeechar = Coffeelex1(Coffeelex, &Coffeelval)
		}

		/* look through exception table */
		xi := 0
		for {
			if CoffeeExca[xi+0] == -1 && CoffeeExca[xi+1] == Coffeestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Coffeen = CoffeeExca[xi+0]
			if Coffeen < 0 || Coffeen == Coffeechar {
				break
			}
		}
		Coffeen = CoffeeExca[xi+1]
		if Coffeen < 0 {
			goto ret0
		}
	}
	if Coffeen == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Coffeelex.Error("syntax error")
			Nerrs++
			if CoffeeDebug >= 1 {
				__yyfmt__.Printf("%s", CoffeeStatname(Coffeestate))
				__yyfmt__.Printf(" saw %s\n", CoffeeTokname(Coffeechar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Coffeep >= 0 {
				Coffeen = CoffeePact[CoffeeS[Coffeep].yys] + CoffeeErrCode
				if Coffeen >= 0 && Coffeen < CoffeeLast {
					Coffeestate = CoffeeAct[Coffeen] /* simulate a shift of "error" */
					if CoffeeChk[Coffeestate] == CoffeeErrCode {
						goto Coffeestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if CoffeeDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", CoffeeS[Coffeep].yys)
				}
				Coffeep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if CoffeeDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", CoffeeTokname(Coffeechar))
			}
			if Coffeechar == CoffeeEofCode {
				goto ret1
			}
			Coffeechar = -1
			goto Coffeenewstate /* try again in the same state */
		}
	}

	/* reduction by production Coffeen */
	if CoffeeDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Coffeen, CoffeeStatname(Coffeestate))
	}

	Coffeent := Coffeen
	Coffeept := Coffeep
	_ = Coffeept // guard against "declared and not used"

	Coffeep -= CoffeeR2[Coffeen]
	CoffeeVAL = CoffeeS[Coffeep+1]

	/* consult goto table to find next state */
	Coffeen = CoffeeR1[Coffeen]
	Coffeeg := CoffeePgo[Coffeen]
	Coffeej := Coffeeg + CoffeeS[Coffeep].yys + 1

	if Coffeej >= CoffeeLast {
		Coffeestate = CoffeeAct[Coffeeg]
	} else {
		Coffeestate = CoffeeAct[Coffeej]
		if CoffeeChk[Coffeestate] != -Coffeen {
			Coffeestate = CoffeeAct[Coffeeg]
		}
	}
	// dummy call; replaced with literal code
	switch Coffeent {

	case 1:
		//line src/coffeephp/parser.y:101
		{ }
	case 2:
		//line src/coffeephp/parser.y:105
		{ }
	case 3:
		//line src/coffeephp/parser.y:106
		{ }
	case 5:
		//line src/coffeephp/parser.y:111
		{ }
	case 6:
		//line src/coffeephp/parser.y:115
		{ }
	case 7:
		//line src/coffeephp/parser.y:116
		{ }
	case 8:
		//line src/coffeephp/parser.y:119
		{
	    CoffeeVAL.val = ast.CreateExprStatementNode(CoffeeS[Coffeept-1].val)
	}
	case 9:
		//line src/coffeephp/parser.y:125
		{  }
	case 10:
		//line src/coffeephp/parser.y:126
		{  }
	case 14:
		//line src/coffeephp/parser.y:138
		{ CoffeeVAL.val  =  CoffeeS[Coffeept-1].val }
	case 15:
		//line src/coffeephp/parser.y:140
		{ 
	            CoffeeVAL.val = ast.CreateExprNode('+', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 16:
		//line src/coffeephp/parser.y:144
		{ 
	            CoffeeVAL.val = ast.CreateExprNode('-', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 17:
		//line src/coffeephp/parser.y:148
		{ 
	            CoffeeVAL.val = ast.CreateExprNode('*', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 18:
		//line src/coffeephp/parser.y:152
		{ 
	            CoffeeVAL.val = ast.CreateExprNode('/', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 19:
		//line src/coffeephp/parser.y:156
		{ 
	            CoffeeVAL.val = ast.CreateExprNode('%', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 20:
		//line src/coffeephp/parser.y:160
		{ 
	            CoffeeVAL.val = ast.CreateExprNode('&', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 21:
		//line src/coffeephp/parser.y:164
		{ 
	            CoffeeVAL.val = ast.CreateExprNode('|', CoffeeS[Coffeept-2].val, CoffeeS[Coffeept-0].val)
	        }
	case 22:
		//line src/coffeephp/parser.y:168
		{ 
	            // $$  = -$2  
        }
	case 23:
		//line src/coffeephp/parser.y:172
		{ 
	            // $$  = regs[$1] 
        }
	case 24:
		CoffeeVAL.val = CoffeeS[Coffeept-0].val
	case 25:
		//line src/coffeephp/parser.y:180
		{
	        CoffeeVAL.val = ast.CreateNumberNode(CoffeeS[Coffeept-0].val.(string))
	    }
	case 26:
		//line src/coffeephp/parser.y:185
		{ }
	case 27:
		//line src/coffeephp/parser.y:189
		{ }
	}
	goto Coffeestack /* stack new state and value */
}