package dco

import (
	"testing"

	"github.com/src-d/git-validate/validate"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func TestRuleCheck_Fail(t *testing.T) {
	dco, err := (&Kind{}).Rule(&validate.RuleConfig{})
	assert.NoError(t, err)

	result, err := dco.Check(nil, &object.Commit{Message: "foo"})
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.NotNil(t, result[0].Rule)
	assert.False(t, result[0].Pass)
	assert.Equal(t, result[0].Location.String(), "000000")
}

func TestRuleCheck_Ignore(t *testing.T) {
	dco, err := (&Kind{}).Rule(&validate.RuleConfig{})
	assert.NoError(t, err)

	result, err := dco.Check(nil, &object.Commit{ParentHashes: []plumbing.Hash{
		plumbing.ZeroHash, plumbing.ZeroHash,
	}})

	assert.NoError(t, err)
	assert.Len(t, result, 0)
}

func TestRuleCheck_Pass(t *testing.T) {
	dco, err := (&Kind{}).Rule(&validate.RuleConfig{})
	assert.NoError(t, err)

	result, err := dco.Check(nil, &object.Commit{Message: "Signed-off-by: Máximo Cuadros <mcuadros@gmail.com>"})
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.True(t, result[0].Pass)
}

func TestKindRule(t *testing.T) {
	dco, err := (&Kind{}).Rule(&validate.RuleConfig{})
	assert.NoError(t, err)

	assert.Equal(t, dco.ID(), "DCO")
}
