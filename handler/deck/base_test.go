package deck

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yoratyo/card-games/config"
	"github.com/yoratyo/card-games/domain"
	mock "github.com/yoratyo/card-games/mocks"
	mock_deckOfCards "github.com/yoratyo/card-games/mocks/deckOfCards"
	"net/http/httptest"
	"testing"
)

type HandlerSuite struct {
	suite.Suite
	*require.Assertions
	ctrl     *gomock.Controller
	resource config.Resource
	helper   *mock.MockHelper
	domain   domain.Domain
	service  *mock_deckOfCards.MockService
	database *mock.MockGormw
	handler  Handler
	recorder *httptest.ResponseRecorder

	ctx *gin.Context
}

func TestService(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.helper = mock.NewMockHelper(s.ctrl)
	s.database = mock.NewMockGormw(s.ctrl)
	s.resource = config.Resource{
		DB:           s.database,
		DefaultCards: config.NewDefaultCards(),
		Helper:       s.helper,
	}
	s.service = mock_deckOfCards.NewMockService(s.ctrl)
	s.domain = domain.Domain{
		DeckOfCards: s.service,
	}
	s.recorder = httptest.NewRecorder()
	c, _ := gin.CreateTestContext(s.recorder)
	s.ctx = c

	s.handler = NewHandler(s.domain, s.resource)
}

func (s *HandlerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *HandlerSuite) TestHandler_New() {
	s.Run("when instantiate", func() {
		result := NewHandler(s.domain, s.resource)
		s.Equal(&handler{
			domain:   s.domain,
			resource: s.resource,
		}, result)
	})
}
