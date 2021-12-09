package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/card-games/config"
	"github.com/yoratyo/card-games/domain/deckOfCards"
	mock "github.com/yoratyo/card-games/mocks"
	mock_deckOfCards "github.com/yoratyo/card-games/mocks/deckOfCards"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
	*require.Assertions
	ctrl       *gomock.Controller
	resource   config.Resource
	repository *mock_deckOfCards.MockRepository
	helper     *mock.MockHelper
	service    deckOfCards.Service
	ctx        *gin.Context
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.helper = mock.NewMockHelper(s.ctrl)
	s.resource = config.Resource{
		DefaultCards: config.NewDefaultCards(),
		Helper:       s.helper,
	}
	s.repository = mock_deckOfCards.NewMockRepository(s.ctrl)

	s.service, _ = New(s.repository, s.resource)
	s.ctx = &gin.Context{}
}

func (s *ServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *ServiceSuite) TestService_New() {
	s.Run("when instantiate", func() {
		result, err := New(s.repository, s.resource)
		s.Nil(err)
		s.Equal(&service{
			repository: s.repository,
			resource:   s.resource,
		}, result)
	})
}
