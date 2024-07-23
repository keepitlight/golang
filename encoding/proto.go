// Copyright (c) 2024 go-kratos
// Copy from go-kratos/kratos/v2/encoding/json/json.go

package encoding

import (
	"errors"
	"reflect"

	"google.golang.org/protobuf/proto"
)

// ProtoMarshal marshal protobuf message. Note JSON marshal use JsonMarshal.
//
// Protobuf 消息编码，注意 JSON 格式编码使用 JsonMarshal。
func ProtoMarshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

// ProtoUnmarshal unmarshal protobuf message. Note JSON unmarshal use JsonUnmarshal.
//
// Protobuf 消息解码，注意 JSON 格式解码使用 JsonUnmarshal。
func ProtoUnmarshal(data []byte, v interface{}) error {
	pm, err := getProtoMessage(v)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data, pm)
}

func getProtoMessage(v interface{}) (proto.Message, error) {
	if msg, ok := v.(proto.Message); ok {
		return msg, nil
	}
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return nil, errors.New("not proto message")
	}

	val = val.Elem()
	return getProtoMessage(val.Interface())
}
