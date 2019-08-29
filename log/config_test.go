package log

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleConfig() {
	_, err := toml.DecodeFile("test.toml", &Conf)
	if err != nil {
		fmt.Printf("fail to get/parse test.toml: err=%v\n", err)
	}
	fmt.Printf("toml: service_name=%v, writers_count=%v", Conf.ServiceName, len(Conf.Writers))
}

func TestConfig(t *testing.T) {
	_, err := toml.DecodeFile("test.toml", &Conf)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(Conf.Writers))
	assert.Equal(t, "test", Conf.ServiceName)
}
