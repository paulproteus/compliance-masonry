package component

import (
	"testing"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/stretchr/testify/assert"
	"github.com/blang/semver"
)

func TestComponentGetters(t *testing.T) {
	testSatisfies := []Satisfies{{}, {}, {}, {}}
	component := Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    common.GeneralReferences{{}},
		Verifications: common.VerificationReferences{{}, {}},
		Satisfies:     testSatisfies,
		SchemaVersion: semver.MustParse("2.0.0"),
	}
	// Test the getters
	assert.Equal(t, "EC2", component.GetKey())
	assert.Equal(t, "Amazon Elastic Compute Cloud", component.GetName())
	assert.Equal(t, &common.GeneralReferences{{}}, component.GetReferences())
	assert.Equal(t, &common.VerificationReferences{{}, {}}, component.GetVerifications())
	assert.Equal(t, semver.MustParse("2.0.0"), component.GetVersion())
	assert.Equal(t, len(testSatisfies), len(component.GetAllSatisfies()))
	for idx, satisfies := range component.GetAllSatisfies() {
		assert.Equal(t, satisfies.GetControlKey(), testSatisfies[idx].GetControlKey())
		assert.Equal(t, satisfies.GetStandardKey(), testSatisfies[idx].GetStandardKey())
		assert.Equal(t, satisfies.GetNarrative(), testSatisfies[idx].GetNarrative())
		assert.Equal(t, satisfies.GetCoveredBy(), testSatisfies[idx].GetCoveredBy())
	}
}

func TestComponentSetters(t *testing.T) {
	component := Component{}
	// Test the setters.
	// Change the version.
	component.SetVersion(semver.MustParse("3.0.0"))
	assert.Equal(t, semver.MustParse("3.0.0"), component.GetVersion())
	// Change the key.
	component.SetKey("FooKey")
	assert.Equal(t, "FooKey", component.GetKey())
}