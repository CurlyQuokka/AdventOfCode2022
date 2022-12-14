package hpsgame

import (
	"log"
	"strings"
)

const (
	rockValue     uint8 = 1
	paperValue    uint8 = 2
	scissorsValue uint8 = 3

	Loose outcome = 0
	Draw  outcome = 3
	Win   outcome = 6

	separator = " "
)

type gameSymbol struct {
	value    uint8
	winsWith *gameSymbol
	loosesTo *gameSymbol
}

type symbols struct {
	rock     *gameSymbol
	paper    *gameSymbol
	scissors *gameSymbol
}

func initSymbols() *symbols {
	s := &symbols{
		rock: &gameSymbol{
			value: rockValue,
		},
		paper: &gameSymbol{
			value: paperValue,
		},
		scissors: &gameSymbol{
			value: scissorsValue,
		},
	}
	s.rock.loosesTo = s.paper
	s.rock.winsWith = s.scissors
	s.paper.loosesTo = s.scissors
	s.paper.winsWith = s.rock
	s.scissors.loosesTo = s.rock
	s.scissors.winsWith = s.paper

	return s
}

type Game struct {
	skirmishes []*skirmish
	gsymbols   *symbols
}

func NewGame(input []string) *Game {
	game := &Game{
		skirmishes: []*skirmish{},
		gsymbols:   initSymbols(),
	}
	for _, line := range input {
		game.skirmishes = append(game.skirmishes, newSkirmish(line, game.gsymbols))
	}
	return game
}

type outcome uint8

type skirmish struct {
	Attack   *gameSymbol
	Response *gameSymbol
	Outcome  outcome
}

func newSkirmish(in string, s *symbols) *skirmish {
	signs := strings.Split(in, separator)
	return &skirmish{
		Attack:   decodeSymbol(signs[0], s),
		Response: decodeSymbol(signs[1], s),
		Outcome:  decodeOutcome(signs[1]),
	}
}

func decodeSymbol(in string, s *symbols) *gameSymbol {
	switch in {
	case "A", "X":
		return s.rock
	case "B", "Y":
		return s.paper
	case "C", "Z":
		return s.scissors
	default:
		log.Fatal("Unknown sign:", in)
	}
	return nil
}

func decodeOutcome(in string) outcome {
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

func (s *skirmish) getOutcome() outcome {
	if s.Response == s.Attack {
		return Draw
	} else if s.Response.winsWith == s.Attack {
		return Win
	}
	return Loose
}

func getSymbolForOutcome(attack *gameSymbol, o outcome) *gameSymbol {
	switch o {
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
			score += uint(sk.getOutcome()) + uint(sk.Response.value)
		} else {
			score += uint(sk.Outcome) + uint(getSymbolForOutcome(sk.Attack, sk.Outcome).value)
		}
	}
	return score
}
