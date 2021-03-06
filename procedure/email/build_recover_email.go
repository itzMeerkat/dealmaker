package email

import (
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type BuildRecoverEmailInterface interface {
	GetJwtAuth() *model.JwtAuth
	GetCredUser() *model.CredUser
	GetEmailContent() *model2.EmailContent
}

func (w *WorkerInstance) BuildRecoverEmail(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(BuildRecoverEmailInterface).GetCredUser()
	token := c.DataDomain.(BuildRecoverEmailInterface).GetJwtAuth()

	email := c.DataDomain.(BuildRecoverEmailInterface).GetEmailContent()

	c.Debugw("email jwt claim", token.TokenClaim)
	toEmail := data.LoginName+"@wustl.edu"
	c.Infow("sending email to", toEmail)

	content := "Hi, "+data.LoginName+"\nHere's your recovery token" + token.Token
	email.Title = "RECOVER PASSWORD"
	email.Body = content
	email.To = toEmail
	email.Recipient = data.LoginName
	return http.StatusOK
}

