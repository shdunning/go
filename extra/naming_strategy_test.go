package extra

import (
	jsoniter "github.com/shdunning/go"
	"github.com/stretchr/testify/require"
	"testing"
	"json_iterator"
)

func Test_lower_case_with_underscores(t *testing.T) {
	should := require.New(t)
	should.Equal("hello_world", LowerCaseWithUnderscores("helloWorld"))
	should.Equal("hello_world", LowerCaseWithUnderscores("HelloWorld"))
	SetNamingStrategy(LowerCaseWithUnderscores)
	output, err := jsoniter.Marshal(struct {
		UserName      string
		FirstLanguage string
	}{
		UserName:      "taowen",
		FirstLanguage: "Chinese",
	})
	should.Nil(err)
	should.Equal(`{"user_name":"taowen","first_language":"Chinese"}`, string(output))
}
