package hpsgame

import (
	"log"
	"strings"
)

const (
	Rock     Sign = 1
	Paper    Sign = 2
	Scissors Sign = 3

	Loose Outcome = 0
	Draw  Outcome = 3
	Win   Outcome = 6

	separator = " "
)

type GameSymbol struct {
	sign     Sign
	winsWith *GameSymbol
	loosesTo *GameSymbol
}

type Symbols struct {
	rock     *GameSymbol
	paper    *GameSymbol
	scissors *GameSymbol
}

func initSymbols() *Symbols {
	symbols := &Symbols{
		rock: &GameSymbol{
			sign: Rock,
		},
		paper: &GameSymbol{
			sign: Paper,
		},
		scissors: &GameSymbol{
			sign: Scissors,
		},
	}
	symbols.rock.loosesTo = symbols.paper
	symbols.rock.winsWith = symbols.scissors
	symbols.paper.loosesTo = symbols.scissors
	symbols.paper.winsWith = symbols.rock
	symbols.scissors.loosesTo = symbols.rock
	symbols.scissors.winsWith = symbols.paper

	return symbols
}

type Game struct {
	skirmishes []*Skirmish
	symbols    *Symbols
}

func NewGame(input []string) *Game {
	game := &Game{
		skirmishes: []*Skirmish{},
		symbols:    initSymbols(),
	}
	for _, line := range input {
		game.skirmishes = append(game.skirmishes, NewSkirmish(line, game.symbols))
	}
	return game
}

type Sign uint8

type Outcome uint8

type Skirmish struct {
	Attack   *GameSymbol
	Response *GameSymbol
	Outcome  Outcome
}

func NewSkirmish(in string, symbols *Symbols) *Skirmish {
	signs := strings.Split(in, separator)
	return &Skirmish{
		Attack:   decodeSymbol(signs[0], symbols),
		Response: decodeSymbol(signs[1], symbols),
		Outcome:  decodeOutcome(signs[1]),
	}
}

func decodeSymbol(in string, symbols *Symbols) *GameSymbol {
	switch in {
	case "A", "X":
		return symbols.rock
	case "B", "Y":
		return symbols.paper
	case "C", "Z":
		return symbols.scissors
	default:
		log.Fatal("Unknown sign:", in)
	}
	return nil
}

func decodeOutcome(in string) Outcome {
	switch in {
	case "X":
		return Loose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		log.Fatal("Unknown sign:", in)
	}
	return 0
}

func (s *Skirmish) getOutcome() Outcome {
	if s.Attack == s.Response {
		return Draw
	} else if s.Response == s.Attack.loosesTo {
		return Win
	}

	return Loose
}

func getSymbolForOutcome(attack *GameSymbol, outcome Outcome) *GameSymbol {
	switch outcome {
	case Draw:
		return attack
	case Win:
		return attack.loosesTo
	default:
		return attack.winsWith
	}
}

func (g *Game) GetScore(isPartOne bool) uint {
	var score uint
	score = 0
	for _, sk := range g.skirmishes {
		if isPartOne {
			score += uint(sk.getOutcome()) + uint(sk.Response.sign)
		} else {
			score += uint(sk.Outcome) + uint(getSymbolForOutcome(sk.Attack, sk.Outcome).sign)
		}
	}
	return score
}
