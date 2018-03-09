package email

import (
	"context"
	pb_email "github.com/onezerobinary/email-box/proto"
	"errors"
)

type EmailServiceServer struct {

}

func (s *EmailServiceServer) SendEmail(ctx context.Context, recipient *pb_email.Recipient) (*pb_email.EmailResponse, error) {

	emailResponse := pb_email.EmailResponse{}

	ok := SendConfirmRegistrationEmail(recipient.Email, recipient.Token)

	if !ok {
		emailResponse.Code = 400
		return &emailResponse, errors.New("It was not possible to send the email")
	}

	return &emailResponse, nil
}