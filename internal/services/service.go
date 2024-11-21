package services

// Сервис общается с file proto и crypto при помощи интерфейса (мост) делаем билдера для сервиса
// билдер для интерфейса(моста)
import (
	"github.com/sirupsen/logrus"
	licensev1 "github.com/splashk1e/jet/gen"
	"github.com/splashk1e/jet/internal/config"
	"github.com/splashk1e/jet/internal/services/cryptoservice"
	"github.com/splashk1e/jet/internal/services/fileservice"
	"github.com/splashk1e/jet/internal/services/protoservice"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type IService interface {
	FileWrite(protoclass protoreflect.ProtoMessage) error
	FileRead() (protoreflect.ProtoMessage, error)
}
type Service struct {
	fileservice.IFileService
	protoservice.IProtoService
	cryptoservice.ICryptoService
}

func (s *Service) FileRead() (protoreflect.ProtoMessage, error) {
	fileText, err := s.IFileService.ReadFile()
	if err != nil {
		return nil, err
	}
	logrus.Debugf("(file service worked succesfully with the result:%s\n", fileText)
	decryptText, err := s.ICryptoService.Decrypt(fileText)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("(cryptography service worked succesfully with the result:%s\n", decryptText)
	if err := s.IProtoService.UnmarshallProto([]byte(decryptText)); err != nil {
		return nil, err
	}
	logrus.Debug("(cryptography service worked succesfully")
	return s.IProtoService.GetProtoClass(), nil
}
func (s *Service) FileWrite(protoclass protoreflect.ProtoMessage) error {
	s.IProtoService.SetProtoClass(protoclass)

	protoText, err := s.IProtoService.MarshallProto()
	if err != nil {
		return err
	}
	logrus.Debugf("protobuf service worked succesfully with the result:%s\n", protoText)
	encryptText, err := s.ICryptoService.Encrypt(string(protoText))
	if err != nil {
		return err
	}
	logrus.Debugf("cryptography service worked succesfully with the result:%s\n", encryptText)
	err = s.IFileService.WriteFile(encryptText)
	logrus.Debug("file service worked succesfully")
	return err
}
func NewService(cfg config.Config) *Service {
	return &Service{
		IFileService:   fileservice.NewFileService(cfg.FilePath),
		IProtoService:  protoservice.NewProtoService(&licensev1.License{}),
		ICryptoService: cryptoservice.NewCryptoAesService(cfg.Key),
	}
}
