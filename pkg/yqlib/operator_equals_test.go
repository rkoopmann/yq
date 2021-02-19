package yqlib

import (
	"testing"
)

var equalsOperatorScenarios = []expressionScenario{
	{
		skipDoc:    true,
		document:   "cat",
		document2:  "dog",
		expression: "select(fi==0) == select(fi==1)",
		expected: []string{
			"D0, P[], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   "{a: { b: {things: \"\"}, f: [1], g: [] }}",
		expression: ".. | select(. == \"\")",
		expected: []string{
			"D0, P[a b things], (!!str)::\"\"\n",
		},
	},
	{
		description: "Match string",
		document:    `[cat,goat,dog]`,
		expression:  `.[] | (. == "*at")`,
		expected: []string{
			"D0, P[0], (!!bool)::true\n",
			"D0, P[1], (!!bool)::true\n",
			"D0, P[2], (!!bool)::false\n",
		},
	},
	{
		description: "Don't match string",
		document:    `[cat,goat,dog]`,
		expression:  `.[] | (. != "*at")`,
		expected: []string{
			"D0, P[0], (!!bool)::false\n",
			"D0, P[1], (!!bool)::false\n",
			"D0, P[2], (!!bool)::true\n",
		},
	},
	{
		description: "Match number",
		document:    `[3, 4, 5]`,
		expression:  `.[] | (. == 4)`,
		expected: []string{
			"D0, P[0], (!!bool)::false\n",
			"D0, P[1], (!!bool)::true\n",
			"D0, P[2], (!!bool)::false\n",
		},
	},
	{
		description: "Dont match number",
		document:    `[3, 4, 5]`,
		expression:  `.[] | (. != 4)`,
		expected: []string{
			"D0, P[0], (!!bool)::true\n",
			"D0, P[1], (!!bool)::false\n",
			"D0, P[2], (!!bool)::true\n",
		},
	},
	{
		skipDoc:    true,
		document:   `a: { cat: {b: apple, c: whatever}, pat: {b: banana} }`,
		expression: `.a | (.[].b == "apple")`,
		expected: []string{
			"D0, P[a cat b], (!!bool)::true\n",
			"D0, P[a pat b], (!!bool)::false\n",
		},
	},
	{
		skipDoc:    true,
		document:   ``,
		expression: `null == null`,
		expected: []string{
			"D0, P[], (!!bool)::true\n",
		},
	},
	{
		description: "Match nulls",
		document:    ``,
		expression:  `null == ~`,
		expected: []string{
			"D0, P[], (!!bool)::true\n",
		},
	},
}

func TestEqualOperatorScenarios(t *testing.T) {
	for _, tt := range equalsOperatorScenarios {
		testScenario(t, &tt)
	}
	documentScenarios(t, "Equals", equalsOperatorScenarios)
}
