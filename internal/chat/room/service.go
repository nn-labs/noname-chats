package room

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"noname-realtime-support-chat/internal/user"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go
type Service interface {
	GetRoomByName(ctx context.Context, name string) (*DTO, error)
	CreateRoom(ctx context.Context, name string, user *user.DTO) (*Room, error)
}

type service struct {
	repository Repository
	userSvc    user.Service
	logger     *zap.SugaredLogger
}

func NewService(repository Repository, userSvc user.Service, logger *zap.SugaredLogger) (Service, error) {
	if repository == nil {
		return nil, errors.New("invalid repository")
	}
	if userSvc == nil {
		return nil, errors.New("invalid user service")
	}
	if logger == nil {
		return nil, errors.New("invalid logger")
	}

	return &service{repository: repository, userSvc: userSvc, logger: logger}, nil
}

func (s *service) GetRoomByName(ctx context.Context, name string) (*DTO, error) {
	room, err := s.repository.GetRoomByName(ctx, name)
	if err != nil {
		s.logger.Errorf("failed to get room: %v", err)
		return nil, err
	}

	return MapToDTO(room), nil
}

func (s *service) CreateRoom(ctx context.Context, roomName string, u *user.DTO) (*Room, error) {
	room, err := NewRoom(roomName)
	if err != nil {
		s.logger.Errorf("failed to create new user %v", err)
		return nil, ErrFailedCreateRoom
	}

	m := &Model{
		ID:   room.ID,
		Name: room.Name,
	}

	_, err = s.repository.CreateRoom(ctx, m)
	if err != nil {
		s.logger.Errorf("failed to save user %v", err)
		return nil, err
	}

	userEntity, _ := user.MapToEntity(u)
	userEntity.SetRoom(&roomName)
	userEntity.SetFreeStatus(true)

	userDto := user.MapToDTO(userEntity)

	err = s.userSvc.UpdateUser(ctx, userDto)

	return room, nil
}