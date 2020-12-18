package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"strings"
)

type ExpressionAtom struct {
	value string // set to () to use the equation below, handled as parentheses
	expr *Expression
}

func (a *ExpressionAtom) Evaluate() int64 {
	if a.value == "()" {
		return a.expr.Evaluate()
	} else {
		return util.MustParseInt(a.value)
	}
}

// Convert is a one-way conversion to part 2 syntax.
func (a *ExpressionAtom) Convert() ExpressionAtom {
	newAtom := *a

	if a.value == "()" {
		newExpr := a.expr.Convert()
		newAtom.expr = &newExpr
	}

	return newAtom
}

type Expression struct {
	Atoms []ExpressionAtom
	Signs []rune
}

func (e *Expression) Evaluate() int64 {
	accumulator := e.Atoms[0].Evaluate()

	atoms := e.Atoms[1:]
	signs := e.Signs

	for len(atoms) > 0 {
		if len(signs) == 0 {
			panic("invalid expression")
		}

		switch signs[0] {
		case '*':
			accumulator *= atoms[0].Evaluate()
		case '+':
			accumulator += atoms[0].Evaluate()
		}

		atoms = atoms[1:]
		signs = signs[1:]
	}

	return accumulator
}

// Convert is a one-way conversion to part 2 syntax.
func (e *Expression) Convert() Expression {
	newExpression := Expression{
		Atoms: []ExpressionAtom{e.Atoms[0].Convert()},
		Signs: make([]rune, 0),
	}

	// Addition will be treated as a sub-expression.
	for k, v := range e.Signs {
		cFirst := newExpression.Atoms[len(newExpression.Atoms) - 1]
		cSecond := e.Atoms[k+1].Convert()

		if v == '+' {
			cSecond = ExpressionAtom{
				value: "()",
				expr:  &Expression{
					Atoms: []ExpressionAtom{cFirst, cSecond},
					Signs: []rune{'+'},
				},
			}

			newExpression.Atoms = newExpression.Atoms[:len(newExpression.Atoms) - 1]
		} else {
			newExpression.Signs = append(newExpression.Signs, v)
		}

		newExpression.Atoms = append(newExpression.Atoms, cSecond)
	}

	return newExpression
}

func ParseExpression(in string) Expression {
	exp := Expression{ make([]ExpressionAtom, 0), make([]rune, 0)}

	atom := ""
	level := 0
	for _,char := range in {
		switch {
		case level > 0:
			if char == '(' {
				level++
			} else if char == ')' {
				level--

				if level == 0 {
					subExpr := ParseExpression(atom)

					exp.Atoms = append(exp.Atoms, ExpressionAtom{
						value: "()",
						expr: &subExpr,
					})

					atom = ""
					break
				}
			}

			atom += string(char)
		case char == '(':
			level++
		case char == ' ':
			if atom == "" {
				continue
			}

			if atom == "+" || atom == "*" {
				exp.Signs = append(exp.Signs, rune(atom[0]))
			} else {
				exp.Atoms = append(exp.Atoms, ExpressionAtom{atom, nil })
			}

			atom = ""
		default:
			atom += string(char)
		}
	}

	if atom != "" {
		if atom == "+" || atom == "*" {
			exp.Signs = append(exp.Signs, rune(atom[0]))
		} else {
			exp.Atoms = append(exp.Atoms, ExpressionAtom{atom, nil })
		}

		atom = ""
	}

	return exp
}

type Day18Solution struct {
	eqs []Expression
}

func (s *Day18Solution) Prepare(input string) {
	lines := strings.Split(input, "\n")
	s.eqs = make([]Expression, len(lines))

	for k,v := range lines {
		s.eqs[k] = ParseExpression(v)
	}
}

func (s *Day18Solution) Part1() string {
	accumulator := int64(0)

	for _,v := range s.eqs {
		accumulator += v.Evaluate()
	}

	return strconv.FormatInt(accumulator, 10)
}

func (s *Day18Solution) Part2() string {
	accumulator := int64(0)

	for _,v := range s.eqs {
		v = v.Convert()
		accumulator += v.Evaluate()
	}

	return strconv.FormatInt(accumulator, 10)
}

