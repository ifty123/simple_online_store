package auth

import (
	"context"
	"errors"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/factory"
	"github.com/ifty123/simple_online_store/internal/repository"
	"github.com/ifty123/simple_online_store/pkg/constant"
	"github.com/ifty123/simple_online_store/pkg/util"
	"github.com/ifty123/simple_online_store/pkg/util/response"
)

type Service struct {
	UserRepository repository.UserRepository
}

type AuthService interface {
	LoginByEmailAndPassword(ctx context.Context, payload *dto.EmailAndPasswordReq) (*dto.UserWithJWTResponse, error)
}

func NewService(f *factory.Factory) AuthService {
	return &Service{
		UserRepository: f.UserRepository,
	}
}

func (s *Service) LoginByEmailAndPassword(ctx context.Context, payload *dto.EmailAndPasswordReq) (*dto.UserWithJWTResponse, error) {
	var res *dto.UserWithJWTResponse

	data, err := s.UserRepository.FindByEmail(ctx, payload.Email)

	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return res, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if !(util.CompareHashPassword(payload.Password, data.Password)) {
		return res, response.ErrorBuilder(
			&response.ErrorConstant.EmailOrPasswordIncorrect, errors.New(response.ErrorConstant.EmailOrPasswordIncorrect.Response.Message),
		)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return res, response.ErrorBuilder(
			&response.ErrorConstant.InternalServerError,
			errors.New("Error when generate token"),
		)
	}

	res = &dto.UserWithJWTResponse{
		UserResponse: dto.UserResponse{
			ID:       data.ID,
			Username: data.Username,
			Email:    data.Email,
		},
		JWT: token,
	}

	return res, nil

}
