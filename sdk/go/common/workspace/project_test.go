package workspace

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestProjectRuntimeInfoRoundtripYAML(t *testing.T) {
	t.Parallel()

	doTest := func(marshal func(interface{}) ([]byte, error), unmarshal func([]byte, interface{}) error) {
		ri := NewProjectRuntimeInfo("nodejs", nil)
		byts, err := marshal(ri)
		assert.NoError(t, err)

		var riRountrip ProjectRuntimeInfo
		err = unmarshal(byts, &riRountrip)
		assert.NoError(t, err)
		assert.Equal(t, "nodejs", riRountrip.Name())
		assert.Nil(t, riRountrip.Options())

		ri = NewProjectRuntimeInfo("nodejs", map[string]interface{}{
			"typescript":   true,
			"stringOption": "hello",
		})
		byts, err = marshal(ri)
		assert.NoError(t, err)
		err = unmarshal(byts, &riRountrip)
		assert.NoError(t, err)
		assert.Equal(t, "nodejs", riRountrip.Name())
		assert.Equal(t, true, riRountrip.Options()["typescript"])
		assert.Equal(t, "hello", riRountrip.Options()["stringOption"])
	}

	doTest(yaml.Marshal, yaml.Unmarshal)
	doTest(json.Marshal, json.Unmarshal)
}

func TestProjectArbitraryExtraData(t *testing.T) {
	t.Parallel()

	doTest := func(marshal func(interface{}) ([]byte, error), unmarshal func([]byte, interface{}) error, repr string) {
		p := Project{
			AdditionalData: map[string]interface{}{
				"resources": map[interface{}]interface{}{
					"foo": map[interface{}]interface{}{
						"type": "example:Resource",
					},
				},
			},
		}
		byts, err := marshal(p)
		assert.NoError(t, err)

		var pRoundtrip Project
		err = unmarshal(byts, &pRoundtrip)
		assert.NoError(t, err)
		assert.Equal(t, "example:Resource",
			pRoundtrip.AdditionalData["resources"].(map[interface{}]interface{})["foo"].(map[interface{}]interface{})["type"])

		assert.Equal(t, string(byts), repr)
	}

	doTest(yaml.Marshal, yaml.Unmarshal,
		`name: ""
runtime: ""
resources:
  foo:
    type: example:Resource
`)
}
