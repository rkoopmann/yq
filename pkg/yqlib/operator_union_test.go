package yqlib

import (
	"testing"
)

var unionOperatorScenarios = []expressionScenario{
	{
		description: "Combine scalars",
		expression:  `1, true, "cat"`,
		expected: []string{
			"D0, P[], (!!int)::1\n",
			"D0, P[], (!!bool)::true\n",
			"D0, P[], (!!str)::cat\n",
		},
	},
	{
		description: "Combine selected paths",
		document:    `{a: fieldA, b: fieldB, c: fieldC}`,
		expression:  `.a, .c`,
		expected: []string{
			"D0, P[a], (!!str)::fieldA\n",
			"D0, P[c], (!!str)::fieldC\n",
		},
	},
}

func TestUnionOperatorScenarios(t *testing.T) {
	for _, tt := range unionOperatorScenarios {
		testScenario(t, &tt)
	}
	documentScenarios(t, "Union", unionOperatorScenarios)
}
