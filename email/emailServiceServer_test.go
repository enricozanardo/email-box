package email

import (
	"testing"
	pb_email "github.com/onezerobinary/email-box/proto"
	"github.com/goinggo/tracelog"
	"time"
	"github.com/spf13/viper"
)

func TestEmailServiceServer_SendEmail(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	viper.SetConfigName("config")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		tracelog.Errorf(err, "main", "main", "Error reading config file")
	}

	recipient := pb_email.Recipient{"ezanardo@onezerobinary.com", "1234", 0}

	ok := SendConfirmRegistrationEmail(recipient.Email, recipient.Token)

	time.Sleep(10 * time.Second)

	if !ok {
		t.Errorf("Error to send the email")
	}

	tracelog.Trace("emailServiceSerer_test","TestEmailServiceServer_SendEmail","Email Successfully sent")
}
