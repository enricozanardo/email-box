package email

import (
	"testing"
	pb_email "github.com/onezerobinary/email-box/proto"
	"github.com/goinggo/tracelog"
)

func TestEmailServiceServer_SendEmail(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	recipient := pb_email.Recipient{"ezanardo@onezerobinary.com", "1234", 0}

	ok := SendConfirmRegistrationEmail(recipient.Email, recipient.Token)

	if !ok {
		t.Errorf("Error to send the email")
	}

	tracelog.Trace("emailServiceSerer_test","TestEmailServiceServer_SendEmail","Email Successfully sent")
}
