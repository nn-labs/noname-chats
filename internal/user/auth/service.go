package auth

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"noname-realtime-support-chat/internal/user"
	"noname-realtime-support-chat/pkg/jwt"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go
type Service interface {
	Registration(ctx context.Context, dto *RegistrationDTO) (*string, error)
	Login(ctx context.Context, dto *LoginDTO) (*string, *string, error)
	Refresh(ctx context.Context, dto *RefreshDTO) (*string, *string, error)
	Logout(ctx context.Context, dto *LogoutDTO) error
}

type service struct {
	userSvc user.Service
	logger  *zap.SugaredLogger
	jwtSvc  jwt.Service
}

func NewService(userSvc user.Service, logger *zap.SugaredLogger, jwtSvc jwt.Service) (Service, error) {
	if userSvc == nil {
		return nil, errors.New("invalid user service")
	}
	if logger == nil {
		return nil, errors.New("invalid logger")
	}
	if jwtSvc == nil {
		return nil, errors.New("invalid jwt service")
	}

	return &service{userSvc: userSvc, logger: logger, jwtSvc: jwtSvc}, nil
}

func (s *service) Registration(ctx context.Context, dto *RegistrationDTO) (*string, error) {
	userDto, err := s.userSvc.CreateUser(ctx, dto.Email, dto.Name, dto.Password)
	if err != nil {
		s.logger.Errorf("failed to save user %v", err)
		return nil, err
	}

	return &userDto.ID, nil
}

func (s *service) Login(ctx context.Context, dto *LoginDTO) (*string, *string, error) {
	userDto, err := s.userSvc.GetUserByEmail(ctx, dto.Email, true)
	if err != nil {
		s.logger.Errorf("failed to find user %v", err)
		return nil, nil, err
	}

	userEntity, err := user.MapToEntity(userDto)
	if err != nil {
		s.logger.Errorf("failed to conver dto %v", err)
		return nil, nil, err
	}

	cp, err := userEntity.CheckPassword(dto.Password)
	if !cp {
		s.logger.Errorf("failed to check password %v", err)
		return nil, nil, err
	}

	accessToken, refreshToken, err := s.jwtSvc.CreateTokens(ctx, userDto.ID, userDto.Support)
	if err != nil {
		s.logger.Errorf("failed to create jwt token %v", err)
		return nil, nil, err
	}

	//user.SetOnline()

	return accessToken, refreshToken, nil
}

func (s *service) Refresh(ctx context.Context, dto *RefreshDTO) (*string, *string, error) {
	payload, err := s.jwtSvc.ParseToken(dto.Token, false)
	if err != nil {
		s.logger.Errorf("failed parse token %v", err)
		return nil, nil, err
	}

	userDto, err := s.userSvc.GetUserById(ctx, payload.Id, false)
	if err != nil {
		s.logger.Errorf("failed to find user %v", err)
		return nil, nil, err
	}

	err = s.jwtSvc.VerifyToken(ctx, payload, false)
	if err != nil {
		s.logger.Errorf("failed to verify token %v", err)
		return nil, nil, err
	}

	accessToken, refreshToken, err := s.jwtSvc.CreateTokens(ctx, payload.Id, userDto.Support)
	if err != nil {
		s.logger.Errorf("failed to create jwt token %v", err)
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func (s *service) Logout(ctx context.Context, dto *LogoutDTO) error {
	payload, err := s.jwtSvc.ParseToken(dto.Token, true)
	if err != nil {
		s.logger.Errorf("failed parse token %v", err)
		return err
	}

	err = s.jwtSvc.VerifyToken(ctx, payload, true)
	if err != nil {
		s.logger.Errorf("failed to verify token %v", err)
		return err
	}

	err = s.jwtSvc.DeleteTokens(ctx, payload)
	if err != nil {
		s.logger.Errorf("failed to delete tokens %v", err)
		return err
	}

	return nil
}