package main

import (
	"fmt"

	"./console"
	"./objects"
)

// Game : game state object info
type Game struct {
	isAiMove          bool
	treshold          int
	winsAiCounter     int
	winsPlayerCounter int
	deck              objects.Deck
	aiHand            objects.Hand
	userHand          objects.Hand
}

// ShowGlobalScore : show counters of wins
func (game *Game) ShowGlobalScore() {
	fmt.Printf("Текущий счет (Вы / компьютер) - %d:%d\n", game.winsPlayerCounter, game.winsAiCounter)
}

// MakeAiTurn : AI player makes his move
func (game *Game) MakeAiTurn() {
	game.aiHand = objects.GetInitialHand(&game.deck)
	count := 0
	for game.aiHand.Calculate() < game.treshold {
		game.deck.Take(&game.aiHand)
		count++
	}
	print("Ход противника. ")
	if count > 0 {
		println("Добрано карт: ", count)
	} else {
		println("Карты не добирались.")
	}
}

// MakeUserTurn : User makes his turn
func (game *Game) MakeUserTurn() {
	game.userHand = objects.GetInitialHand(&game.deck)
	println("Ваш ход.")
	for game.userHand.Calculate() < 21 {
		print("Текущая рука: ", game.userHand.GetRepresentation())
		choice := console.Ask("Взять еще карту? (Y/N)", []rune{'y', 'n'})
		if choice == 0 {
			game.deck.Take(&game.userHand)
		} else {
			break
		}
	}
	if game.userHand.Calculate() > 21 {
		println("Перебор: ", game.userHand.GetRepresentation())
	} else if game.userHand.Calculate() == 21 {
		println("Очко!: ", game.userHand.GetRepresentation())
	}
}

// CheckResult : compare hands, show winner
func (game *Game) CheckResult() {
	println("Рука противника:", game.aiHand.GetRepresentation())
	userScore := game.userHand.Calculate()
	aiScore := game.aiHand.Calculate()

	switch {
	// both has 21 !
	case userScore == aiScore && aiScore == 21:
		if game.isAiMove {
			game.playerLost()
		} else {
			game.playerWins(21)
		}

	// user lost
	case userScore > 21 && aiScore <= 21:
		game.playerLost()
	// user wins
	case userScore <= 21 && aiScore > 21:
		game.playerWins(userScore)
	// draw
	case userScore > 21 && aiScore > 21:
		println("Ничья.")

	// user lost
	case userScore < aiScore:
		game.playerLost()
	// user wins
	case userScore > aiScore:
		game.playerWins(userScore)
	// draw
	default:
		println("Ничья.")
	}

	game.isAiMove = !game.isAiMove
	game.deck.Shuffle()
}

// playerWins : show message, add score to player, shuffle deck.
func (game *Game) playerWins(playerScore int) {
	println("Вы победили!")
	game.winsPlayerCounter++
	game.treshold = (game.treshold + playerScore) / 2
}

// playerWins : show message, add score to computer, shuffle deck.
func (game *Game) playerLost() {
	println("Вы проиграли...")
	game.winsAiCounter++
}

// NewGame : create new state game object
func NewGame() Game {
	return Game{
		isAiMove:          false,
		treshold:          17,
		winsAiCounter:     0,
		winsPlayerCounter: 0,
		deck:              objects.NewDeck(),
	}
}
