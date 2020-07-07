package main

import proto "github.com/golang/protobuf/proto"

var _ = proto.Marshal

type ProtoTest struct {
	field0 string `protobuf:"bytes,1,opt,name=field0" json:"field0,omitempty"`
	field1 string `protobuf:"bytes,2,opt,name=field1" json:"field1,omitempty"`
	field2 string `protobuf:"bytes,3,opt,name=field2" json:"field2,omitempty"`
	field3 string `protobuf:"bytes,4,opt,name=field3" json:"field3,omitempty"`
	field4 string `protobuf:"bytes,5,opt,name=field4" json:"field4,omitempty"`
	field5 string `protobuf:"bytes,6,opt,name=field5" json:"field5,omitempty"`
	field6 string `protobuf:"bytes,7,opt,name=field6" json:"field6,omitempty"`
	field7 string `protobuf:"bytes,8,opt,name=field7" json:"field7,omitempty"`
	field8 string `protobuf:"bytes,9,opt,name=field8" json:"field8,omitempty"`
	field9 string `protobuf:"bytes,10,opt,name=field9" json:"field9,omitempty"`
	size0  int32  `protobuf:"varint,11,opt,name=size0" json:"size0,omitempty"`
	size1  int32  `protobuf:"varint,12,opt,name=size1" json:"size1,omitempty"`
	size2  int32  `protobuf:"varint,13,opt,name=size2" json:"size2,omitempty"`
	size3  int32  `protobuf:"varint,14,opt,name=size3" json:"size3,omitempty"`
	size4  int32  `protobuf:"varint,15,opt,name=size4" json:"size4,omitempty"`
	size5  int32  `protobuf:"varint,16,opt,name=size5" json:"size5,omitempty"`
	size6  int32  `protobuf:"varint,17,opt,name=size6" json:"size6,omitempty"`
	size7  int32  `protobuf:"varint,18,opt,name=size7" json:"size7,omitempty"`
	size8  int32  `protobuf:"varint,19,opt,name=size8" json:"size8,omitempty"`
	size9  int32  `protobuf:"varint,20,opt,name=size9" json:"size9,omitempty"`
}

func (m *ProtoTest) Reset()         { *m = ProtoTest{} }
func (m *ProtoTest) String() string { return proto.CompactTextString(m) }
func (*ProtoTest) ProtoMessage()    {}

func init() {
}
