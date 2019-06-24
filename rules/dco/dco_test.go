package dco

import (
	"testing"

	"github.com/vbatts/git-validation/compliance"

	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/stretchr/testify/assert"
)

func TestValidateDCO_Fail(t *testing.T) {
	dco, err := (&Kind{}).Rule(&compliance.RuleConfig{})
	assert.NoError(t, err)

	result, err := dco.Check(nil, &object.Commit{Message: "foo"})
	assert.NoError(t, err)
	assert.False(t, result.Pass)
}

func TestValidateDCO_Ignore(t *testing.T) {
	dco, err := (&Kind{}).Rule(&compliance.RuleConfig{})
	assert.NoError(t, err)

	result, err := dco.Check(nil, &object.Commit{ParentHashes: []plumbing.Hash{
		plumbing.ZeroHash, plumbing.ZeroHash,
	}})

	assert.NoError(t, err)
	assert.True(t, result.Pass)
}

func TestValidateDCO_Pass(t *testing.T) {
	dco, err := (&Kind{}).Rule(&compliance.RuleConfig{})
	assert.NoError(t, err)

	result, err := dco.Check(nil, &object.Commit{Message: "Signed-off-by: Máximo Cuadros <mcuadros@gmail.com>"})
	assert.NoError(t, err)
	assert.True(t, result.Pass)
}
