package encoding

import (
	"gopkg.in/yaml.v3"
)

// YamlMarshal is a shortcut to yaml.Marshal
//
// YamlMarshal 是 yaml.Marshal 的快捷方式
func YamlMarshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// YamlUnmarshal is a shortcut to yaml.Unmarshal
//
// YamlUnmarshal 是 yaml.Unmarshal 的快捷方式
func YamlUnmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
