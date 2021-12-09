# - Domain
# - deckOfCards
mkdir -p mocks/deckOfCards
mockgen --source=domain/deckOfCards/service.go -package=mock_deckOfCards --destination=mocks/deckOfCards/mock_service.go
mockgen --source=domain/deckOfCards/repository.go -package=mock_deckOfCards --destination=mocks/deckOfCards/mock_repository.go

# - Helper
mockgen --source=config/helper.go -package=mock --destination=mocks/mock_helper.go
# - Database
mockgen --source=config/database/interface.go -package=mock --destination=mocks/mock_database.go