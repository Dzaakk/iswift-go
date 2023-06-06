package register

import (
	dto "iswift-go-project/internal/register/dto"
	userDto "iswift-go-project/internal/user/dto"
	userUseCase "iswift-go-project/internal/user/usecase"
	mail "iswift-go-project/pkg/mail/sendgrid"
)

type RegisterUseCase interface {
	Register(userDto userDto.UserRequestBody) error
}

type RegisterUseCaseImpl struct {
	userUseCase userUseCase.UserUseCase
	mail        mail.Mail
}

func NewRegisterUseCase(
	userUseCase userUseCase.UserUseCase,
	mail mail.Mail,
) RegisterUseCase {
	return &RegisterUseCaseImpl{userUseCase, mail}
}

func (ru *RegisterUseCaseImpl) Register(userDto userDto.UserRequestBody) error {
	// Check ke dalam module user
	user, err := ru.userUseCase.Create(userDto)

	if err != nil {
		return err
	}

	//kirim email melalui sendgrid
	email := dto.CreateEmailVerification{
		SUBJECT:           "Kode Verifikasi",
		EMAIL:             user.Email,
		VERIFICATION_CODE: user.CodeVerified,
	}

	go ru.mail.SendVerifivationEmail(user.Email, email)

	return nil
}
