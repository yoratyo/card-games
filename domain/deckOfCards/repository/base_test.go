package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yoratyo/card-games/config"
	"github.com/yoratyo/card-games/domain/deckOfCards"
	mock "github.com/yoratyo/card-games/mocks"
	"testing"
)

type RepositorySuite struct {
	suite.Suite
	*require.Assertions
	ctrl       *gomock.Controller
	resource   config.Resource
	database   *mock.MockGormw
	helper     *mock.MockHelper
	repository deckOfCards.Repository
	ctx        *gin.Context
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (s *RepositorySuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.database = mock.NewMockGormw(s.ctrl)
	s.helper = mock.NewMockHelper(s.ctrl)
	s.resource = config.Resource{
		DB:     s.database,
		Helper: s.helper,
	}

	s.repository, _ = New(s.resource)
	s.ctx = &gin.Context{}
}

func (s *RepositorySuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *RepositorySuite) TestRepository_New() {
	s.Run("when instantiate", func() {
		result, err := New(s.resource)
		s.Nil(err)
		s.Equal(&repository{
			resource: s.resource,
		}, result)
	})
}
