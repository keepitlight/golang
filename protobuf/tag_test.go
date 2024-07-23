package protobuf

import (
	"testing"

	"github.com/keepitlight/golang/types"
)

func TestTagParse(t *testing.T) {
	const testTag1 = "varint,6,opt,name=level,proto3,enum=config.v1.LogLevel"
	if tag, err := TagParse(testTag1); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != VarInt {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 6 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "level" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	} else if tag.Enum != "config.v1.LogLevel" {
		t.Errorf("unexpected enum: %v", tag.Enum)
	}
	const testTag2 = "varint,4,opt,name=max_size,json=maxSize,proto3"
	if tag, err := TagParse(testTag2); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != VarInt {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 4 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "max_size" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Json != "maxSize" {
		t.Errorf("unexpected enum: %v", tag.Json)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
	const testTag3 = "bytes,3,opt,name=namespace,proto3"
	if tag, err := TagParse(testTag3); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != Bytes {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 3 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "namespace" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
	const testTag4 = "bytes,2,rep,name=acl,proto3"
	if tag, err := TagParse(testTag4); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != Bytes {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 2 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Repeat {
		t.Errorf("unexpected repeat: %v", tag.Repeat)
	} else if tag.Name != "acl" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
}
func TestParse(t *testing.T) {
	const testTag1 = `protobuf:"varint,6,opt,name=level,proto3,enum=config.v1.LogLevel"`
	if tag, _, _, err := Parse(testTag1); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != VarInt {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 6 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "level" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	} else if tag.Enum != "config.v1.LogLevel" {
		t.Errorf("unexpected enum: %v", tag.Enum)
	}
	const testTag2 = `protobuf:"varint,4,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	if tag, _, _, err := Parse(testTag2); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != VarInt {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 4 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "max_size" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Json != "maxSize" {
		t.Errorf("unexpected enum: %v", tag.Json)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
	const testTag3 = `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	if tag, _, _, err := Parse(testTag3); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != Bytes {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 3 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "namespace" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
	const testTag4 = `protobuf:"bytes,2,rep,name=acl,proto3" json:"acl,omitempty"`
	if tag, _, _, err := Parse(testTag4); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != Bytes {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 2 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Repeat {
		t.Errorf("unexpected repeat: %v", tag.Repeat)
	} else if tag.Name != "acl" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
	const testTag5 = `protobuf:"bytes,2,rep,name=mappings,proto2" json:"mappings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto2" protobuf_val:"bytes,2,opt,name=value,proto3"`
	if tag, k, v, err := Parse(testTag5); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if k == nil {
		t.Errorf("unexpected key nil")
	} else if v == nil {
		t.Errorf("unexpected value nil")
	} else if tag.Type != Bytes {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 2 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Repeat {
		t.Errorf("unexpected repeat: %v", tag.Repeat)
	} else if tag.Name != "mappings" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V2 {
		t.Errorf("unexpected version: %v", tag.Version)
	} else if k.Type != Bytes {
		t.Errorf("unexpected key type: %v", k.Type)
	} else if k.Index != 1 {
		t.Errorf("unexpected key index: %v", k.Index)
	} else if !k.Optional {
		t.Errorf("unexpected key optional: %v", k.Optional)
	} else if k.Name != "key" {
		t.Errorf("unexpected key name: %v", k.Name)
	} else if k.Version != V2 {
		t.Errorf("unexpected key version: %v", k.Version)
	} else if v.Type != Bytes {
		t.Errorf("unexpected value type: %v", v.Type)
	} else if v.Index != 2 {
		t.Errorf("unexpected value index: %v", v.Index)
	} else if !v.Optional {
		t.Errorf("unexpected value optional: %v", v.Optional)
	} else if v.Name != "value" {
		t.Errorf("unexpected value name: %v", v.Name)
	} else if v.Version != V3 {
		t.Errorf("unexpected value version: %v", v.Version)
	}
}
func TestParseTag(t *testing.T) {
	const testTag1 = `protobuf:"varint,6,opt,name=level,proto3,enum=config.v1.LogLevel"`
	if tag, err := types.ParseTag[*Tag](testTag1, TagName); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != VarInt {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 6 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "level" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	} else if tag.Enum != "config.v1.LogLevel" {
		t.Errorf("unexpected enum: %v", tag.Enum)
	}
	const testTag2 = `protobuf:"varint,4,opt,name=max_size,json=maxSize,proto3" json:"max_size,omitempty"`
	if tag, err := types.ParseTag[*Tag](testTag2, TagName); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != VarInt {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 4 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "max_size" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Json != "maxSize" {
		t.Errorf("unexpected enum: %v", tag.Json)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
	const testTag3 = `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	if tag, err := types.ParseTag[*Tag](testTag3, TagName); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != Bytes {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 3 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Optional {
		t.Errorf("unexpected optional: %v", tag.Optional)
	} else if tag.Name != "namespace" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
	const testTag4 = `protobuf:"bytes,2,rep,name=acl,proto3" json:"acl,omitempty"`
	if tag, err := types.ParseTag[*Tag](testTag4, TagName); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tag == nil {
		t.Errorf("unexpected nil")
	} else if tag.Type != Bytes {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 2 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Repeat {
		t.Errorf("unexpected repeat: %v", tag.Repeat)
	} else if tag.Name != "acl" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V3 {
		t.Errorf("unexpected version: %v", tag.Version)
	}
	const testTag5 = `protobuf:"bytes,2,rep,name=mappings,proto2" json:"mappings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto2" protobuf_val:"bytes,2,opt,name=value,proto3"`
	if tags, err := types.ParseTags(testTag5); err != nil {
		t.Errorf("unexpected error: %v", err)
	} else if tags == nil {
		t.Errorf("unexpected nil")
	} else if tag, f1 := types.Tag[*Tag](tags, TagName); !f1 {
		t.Errorf("unexpected error: tag not found")
	} else if k, f2 := types.Tag[*Tag](tags, KeyTagName); !f2 {
		t.Errorf("unexpected error: tag key not found")
	} else if v, f3 := types.Tag[*Tag](tags, ValTagName); !f3 {
		t.Errorf("unexpected error: tag value not found")
	} else if k == nil {
		t.Errorf("unexpected key nil")
	} else if v == nil {
		t.Errorf("unexpected value nil")
	} else if tag.Type != Bytes {
		t.Errorf("unexpected type: %v", tag.Type)
	} else if tag.Index != 2 {
		t.Errorf("unexpected index: %v", tag.Index)
	} else if !tag.Repeat {
		t.Errorf("unexpected repeat: %v", tag.Repeat)
	} else if tag.Name != "mappings" {
		t.Errorf("unexpected name: %v", tag.Name)
	} else if tag.Version != V2 {
		t.Errorf("unexpected version: %v", tag.Version)
	} else if k.Type != Bytes {
		t.Errorf("unexpected key type: %v", k.Type)
	} else if k.Index != 1 {
		t.Errorf("unexpected key index: %v", k.Index)
	} else if !k.Optional {
		t.Errorf("unexpected key optional: %v", k.Optional)
	} else if k.Name != "key" {
		t.Errorf("unexpected key name: %v", k.Name)
	} else if k.Version != V2 {
		t.Errorf("unexpected key version: %v", k.Version)
	} else if v.Type != Bytes {
		t.Errorf("unexpected value type: %v", v.Type)
	} else if v.Index != 2 {
		t.Errorf("unexpected value index: %v", v.Index)
	} else if !v.Optional {
		t.Errorf("unexpected value optional: %v", v.Optional)
	} else if v.Name != "value" {
		t.Errorf("unexpected value name: %v", v.Name)
	} else if v.Version != V3 {
		t.Errorf("unexpected value version: %v", v.Version)
	}
}
