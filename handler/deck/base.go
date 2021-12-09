package deck

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/card-games/config"
	"github.com/yoratyo/card-games/domain"
	"net/http"
)

type (
	Handler interface {
		CreateNewDeck(c *gin.Context)
		OpenDeck(c *gin.Context)
		DrawCardsFromDeck(c *gin.Context)
	}
	handler struct {
		domain   domain.Domain
		resource config.Resource
	}
)

func NewHandler(domain domain.Domain, resource config.Resource) (Handler) {
	return &handler{
		domain:   domain,
		resource: resource,
	}
}

func addError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}
