package protoservice

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type IProtoService interface {
	MarshallProto() ([]byte, error)
	UnmarshallProto(text []byte) error
	GetProtoClass() protoreflect.ProtoMessage
	SetProtoClass(protoclass protoreflect.ProtoMessage)
}

type ProtoService struct {
	protoclass protoreflect.ProtoMessage
}

func NewProtoService(protoclass protoreflect.ProtoMessage) *ProtoService {
	return &ProtoService{protoclass: protoclass}
}

func (s *ProtoService) MarshallProto() ([]byte, error) {
	text, err := proto.Marshal(s.protoclass)
	if err != nil {
		return nil, err
	}
	return text, nil
}

func (s *ProtoService) UnmarshallProto(text []byte) error {
	if err := proto.Unmarshal(text, s.protoclass); err != nil {
		return err
	}
	return nil
}
func (s *ProtoService) GetProtoClass() protoreflect.ProtoMessage {
	return s.protoclass
}
func (s *ProtoService) SetProtoClass(protoclass protoreflect.ProtoMessage) {
	s.protoclass = protoclass
}
