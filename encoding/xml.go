package encoding

import (
	"encoding/xml"
)

// XmlMarshal is a shortcut of xml.Marshal
//
// xml.Marshal 的快捷方式
func XmlMarshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// XmlUnmarshal is a shortcut of xml.Unmarshal
//
// xml.Unmarshal 的快捷方式
func XmlUnmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}
