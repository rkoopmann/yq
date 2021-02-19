package yqlib

import (
	"testing"
)

var anchorOperatorScenarios = []expressionScenario{
	{
		description: "Get anchor",
		document:    `a: &billyBob cat`,
		expression:  `.a | anchor`,
		expected: []string{
			"D0, P[a], (!!str)::billyBob\n",
		},
	},
	{
		description: "Set anchor",
		document:    `a: cat`,
		expression:  `.a anchor = "foobar"`,
		expected: []string{
			"D0, P[], (doc)::a: &foobar cat\n",
		},
	},
	{
		description: "Set anchor relatively using assign-update",
		document:    `a: {b: cat}`,
		expression:  `.a anchor |= .b`,
		expected: []string{
			"D0, P[], (doc)::a: &cat {b: cat}\n",
		},
	},
	{
		description: "Get alias",
		document:    `{b: &billyBob meow, a: *billyBob}`,
		expression:  `.a | alias`,
		expected: []string{
			"D0, P[a], (!!str)::billyBob\n",
		},
	},
	{
		description: "Set alias",
		document:    `{b: &meow purr, a: cat}`,
		expression:  `.a alias = "meow"`,
		expected: []string{
			"D0, P[], (doc)::{b: &meow purr, a: *meow}\n",
		},
	},
	{
		description: "Set alias relatively using assign-update",
		document:    `{b: &meow purr, a: {f: meow}}`,
		expression:  `.a alias |= .f`,
		expected: []string{
			"D0, P[], (doc)::{b: &meow purr, a: *meow}\n",
		},
	},
	{
		description: "Explode alias and anchor",
		document:    `{f : {a: &a cat, b: *a}}`,
		expression:  `explode(.f)`,
		expected: []string{
			"D0, P[], (doc)::{f: {a: cat, b: cat}}\n",
		},
	},
	{
		description: "Explode with no aliases or anchors",
		document:    `a: mike`,
		expression:  `explode(.a)`,
		expected: []string{
			"D0, P[], (doc)::a: mike\n",
		},
	},
	{
		description: "Explode with alias keys",
		document:    `{f : {a: &a cat, *a: b}}`,
		expression:  `explode(.f)`,
		expected: []string{
			"D0, P[], (doc)::{f: {a: cat, cat: b}}\n",
		},
	},
	{
		description: "Explode with merge anchors",
		document:    mergeDocSample,
		expression:  `explode(.)`,
		expected: []string{`D0, P[], (doc)::foo:
    a: foo_a
    thing: foo_thing
    c: foo_c
bar:
    b: bar_b
    thing: bar_thing
    c: bar_c
foobarList:
    b: bar_b
    a: foo_a
    thing: bar_thing
    c: foobarList_c
foobar:
    c: foo_c
    a: foo_a
    thing: foobar_thing
`},
	},
	{
		skipDoc:    true,
		document:   mergeDocSample,
		expression: `.foo* | explode(.) | (. style="flow")`,
		expected: []string{
			"D0, P[foo], (!!map)::{a: foo_a, thing: foo_thing, c: foo_c}\n",
			"D0, P[foobarList], (!!map)::{b: bar_b, a: foo_a, thing: bar_thing, c: foobarList_c}\n",
			"D0, P[foobar], (!!map)::{c: foo_c, a: foo_a, thing: foobar_thing}\n",
		},
	},
	{
		skipDoc:    true,
		document:   mergeDocSample,
		expression: `.foo* | explode(explode(.)) | (. style="flow")`,
		expected: []string{
			"D0, P[foo], (!!map)::{a: foo_a, thing: foo_thing, c: foo_c}\n",
			"D0, P[foobarList], (!!map)::{b: bar_b, a: foo_a, thing: bar_thing, c: foobarList_c}\n",
			"D0, P[foobar], (!!map)::{c: foo_c, a: foo_a, thing: foobar_thing}\n",
		},
	},
	{
		skipDoc:    true,
		document:   `{f : {a: &a cat, b: &b {f: *a}, *a: *b}}`,
		expression: `explode(.f)`,
		expected: []string{
			"D0, P[], (doc)::{f: {a: cat, b: {f: cat}, cat: {f: cat}}}\n",
		},
	},
}

func TestAnchorAliaseOperatorScenarios(t *testing.T) {
	for _, tt := range anchorOperatorScenarios {
		testScenario(t, &tt)
	}
	documentScenarios(t, "Anchor and Alias Operators", anchorOperatorScenarios)
}
