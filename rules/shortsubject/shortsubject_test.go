package shortsubject

import (
	"strings"
	"testing"

	"github.com/vbatts/git-validation/compliance"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func TestRuleCheck(t *testing.T) {
	short, err := (&Kind{}).Rule(&compliance.RuleConfig{})
	assert.NoError(t, err)

	testCases := []struct {
		msg  string
		pass bool
	}{
		{"", false},
		{"foo", true},
		{strings.Repeat("0", 90), false},
	}

	for _, tc := range testCases {
		result, err := short.Check(nil, &object.Commit{Message: tc.msg})
		assert.NoError(t, err)
		assert.Equal(t, tc.pass, result.Pass, tc.pass)
	}
}

func TestKindRule(t *testing.T) {
	dco, err := (&Kind{}).Rule(&compliance.RuleConfig{})
	assert.NoError(t, err)

	assert.Equal(t, dco.ID(), "short-subject")
}
