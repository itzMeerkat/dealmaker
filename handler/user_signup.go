package handler

import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/email/model"
	"github.com/dealmaker/shared/auth/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type UserSignupDomain struct {
	base.Base
	model.CredUser
	model.JwtAuth
	model2.EmailContent
}

type UserSignupInput struct {
	LoginName string
	HashedPassword string
}

func UserSignup(c *gin.Context) {
	input := UserSignupInput{}

	err := c.Bind(&input)
	if err != nil {
		return
	}

	domain := UserSignupDomain{
		CredUser: model.CredUser{
			HashedPassword: input.HashedPassword,
			LoginName:      input.LoginName,
		},
	}

	s := factory.Factory.Get("/auth/user/signup")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}

	c.JSON(code, nil)
}
