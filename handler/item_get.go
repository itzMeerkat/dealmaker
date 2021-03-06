package handler

import (
	"github.com/dealmaker/factory"
	model2 "github.com/dealmaker/procedure/item/model"
	"github.com/dealmaker/shared/base"
	"github.com/gin-gonic/gin"
	"github.com/itzmeerkat/streamline"
	"net/http"
)

type ItemGetDomain struct {
	base.Base
	//model.JwtAuth
	model2.GetItemDomain
}

type ItemGetInput struct {
	model2.QueryFilter
}

type ItemGetResponse struct {
	Message string
	Items []model2.Item
}

func ItemGetHandler(c *gin.Context) {
	input := ItemGetInput{}

	err := c.Bind(&input)
	if err != nil {
		return
	}

	domain := ItemGetDomain{
		GetItemDomain: model2.GetItemDomain{
			QueryFilter: input.QueryFilter,
		},
	}

	s := factory.Factory.Get("/item/get")
	conv := streamline.NewConveyorBelt(s, c, &domain, GenLogMeta)
	conv.Debugw("input", domain)
	code := conv.Run()
	if code != http.StatusOK {
		c.AbortWithStatusJSON(code, domain.GetBase())
	}


	resp := ItemGetResponse{
		Message: domain.BaseMessage,
		Items:  domain.Result,
	}
	c.JSON(code, resp)
}