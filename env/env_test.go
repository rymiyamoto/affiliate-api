package env

import (
	"testing"

	"github.com/rymiyamoto/affiliate-api/model/common"
	"github.com/stretchr/testify/assert"
)

func Test_Env(t *testing.T) {
	t.Parallel()

	tests := []struct {
		key    string
		env    string
		expect string
	}{
		{key: "GO_ENV", env: common.EnvProd, expect: common.EnvProd},
		{key: "GO_ENV", env: common.EnvDev, expect: common.EnvDev},
		{key: "GO_ENV", env: common.EnvTest, expect: common.EnvTest},
	}

	for _, test := range tests {
		t.Setenv(test.key, test.env)
		ret, err := Value(test.key)
		assert.Nil(t, err)
		assert.Equal(t, ret, test.expect)
	}
}
