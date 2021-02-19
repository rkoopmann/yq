package yqlib

import (
	"testing"
)

var booleanOperatorScenarios = []expressionScenario{
	{
		description: "OR example",
		expression:  `true or false`,
		expected: []string{
			"D0, P[], (!!bool)::true\n",
		},
	},
	{
		description: "AND example",
		expression:  `true and false`,
		expected: []string{
			"D0, P[], (!!bool)::false\n",
		},
	},
	{
		document:    "[{a: bird, b: dog}, {a: frog, b: bird}, {a: cat, b: fly}]",
		description: "Matching nodes with select, equals and or",
		expression:  `[.[] | select(.a == "cat" or .b == "dog")]`,
		expected: []string{
			"D0, P[], (!!seq)::- {a: bird, b: dog}\n- {a: cat, b: fly}\n",
		},
	},
	{
		skipDoc:    true,
		expression: `false or false`,
		expected: []string{
			"D0, P[], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   `{a: true, b: false}`,
		expression: `.[] or (false, true)`,
		expected: []string{
			"D0, P[a], (!!bool)::true\n",
			"D0, P[a], (!!bool)::true\n",
			"D0, P[b], (!!bool)::false\n",
			"D0, P[b], (!!bool)::true\n",
		},
	},
	{
		description: "Not true is false",
		expression:  `true | not`,
		expected: []string{
			"D0, P[], (!!bool)::false\n",
		},
	},
	{
		description: "Not false is true",
		expression:  `false | not`,
		expected: []string{
			"D0, P[], (!!bool)::true\n",
		},
	},
	{
		description: "String values considered to be true",
		expression:  `"cat" | not`,
		expected: []string{
			"D0, P[], (!!bool)::false\n",
		},
	},
	{
		description: "Empty string value considered to be true",
		expression:  `"" | not`,
		expected: []string{
			"D0, P[], (!!bool)::false\n",
		},
	},
	{
		description: "Numbers are considered to be true",
		expression:  `1 | not`,
		expected: []string{
			"D0, P[], (!!bool)::false\n",
		},
	},
	{
		description: "Zero is considered to be true",
		expression:  `0 | not`,
		expected: []string{
			"D0, P[], (!!bool)::false\n",
		},
	},
	{

		description: "Null is considered to be false",
		expression:  `~ | not`,
		expected: []string{
			"D0, P[], (!!bool)::true\n",
		},
	},
}

func TestBooleanOperatorScenarios(t *testing.T) {
	for _, tt := range booleanOperatorScenarios {
		testScenario(t, &tt)
	}
	documentScenarios(t, "Boolean Operators", booleanOperatorScenarios)
}
