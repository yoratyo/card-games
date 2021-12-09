# Card Games

_December 9th 2021_

_This is a simple API written in Go to demonstrate deck for playing cards.
I'm using MySQL as the database. The API designed to make it easy to extend and reusable to create another playing cards like poker and blackjack._
---
The main feature of this API is to create decks of playing cards.
Those decks of cards can be viewed and can be drawn from a deck—removing the card from it and leaving fewer remaining cards in the deck.
♠️♥️♣️♦️

## Usage ⚡️

The API structure:
```
POST   /deck                  # create new deck (optional parameters: cards, shuffle)
GET    /deck/:deck_id         # fetch deck by id
PATCH  /deck/:deck_id/draw    # draw cards from deck (optional parameters: count)
```
_The `deck_id` using format UUID._

## Run 🛠

> All commands are to be run from the root project directory.
> This project uses Go version `1.16`.
> Current `go` executables should detect and install dependencies correctly.
> This project using environment variable file (.env) to store properties of MySQL connection and App port.

On root project, view file `.env.example` to see list of field, than create new `.env` file on your local.
```
#Database
#MySQL
DB_HOST=
DB_PORT=
DB_NAME=
DB_USER=
DB_PASSWORD=

#App
APP_PORT=
```

To manually install dependencies, run:
```
go get -d -v ./...
go install -v ./...
```

To start the API, run:
```
go run main.go
```

To create mocks file, run:
```
./generateMock.sh
```

To test all available test files, run:
```
go test -v ./...
```

## Room for improvement 🚧

_Some stuff that I'd continue if there had been more time_

- Simplify conversion between DAO and DTO to be more efficient
- Currently just tested with DB MySQL, next will try to use other DB. I'm already try to use SQLite, but it comes a problems when it can't support some complex query.
- Add more tests