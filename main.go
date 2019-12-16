package main

import (
	"./console"
)

func main() {
	Intro()
	game := NewGame()
	for {
		if game.isAiMove {
			game.MakeAiTurn()
			game.MakeUserTurn()
		} else {
			game.MakeUserTurn()
			game.MakeAiTurn()
		}
		game.CheckResult()
		game.ShowGlobalScore()

		choice := console.Ask("Продолжим? (Y/N)", []rune{'y', 'n'})
		if choice == 1 {
			break
		}
		println("---------")
	}
	print("\nНажмите любую клавишу...")
	console.GetChar()
}

// Intro : Show game introduction
func Intro() {
	println("\n------------------------------------- Игра \"Очко\", она же \"21\" ---------------------------------")
	println("Правила довольно просты.")
	println("Цель игры - набрать до 21 очка включительно.")
	println("Колода из 36 карт тасуется, после чего каждый игрок получает по две карты.")
	println("Стоимость карт: числовые значения (6-10) по номиналу, картинки 2-3-4, туз 11.")
	println("Игроки могут добирать дополнительные карты, если у них нет перебора.")
	println("В случае, если игрок первым набрал 21 очко, он выигрывает.")
	println("Удачной игры!")
	println("------------------------------------------------------------------------------------------------\n")
}
