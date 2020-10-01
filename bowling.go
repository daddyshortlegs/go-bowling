package go_bowling

type Bowling struct {
	rolls []int
}

func (bowling *Bowling) roll(pins int) {
	bowling.rolls = append(bowling.rolls, pins)
}

func (bowling *Bowling) score() int {
	score := 0
	i := 0
	for frame := 0; frame < 10; frame++ {
		if isSpare(bowling, i) {
			score += 10 + bowling.rolls[i+2]
			i += 2
		} else if isStrike(bowling, i) {
			score += 10 + bowling.rolls[i+1] + bowling.rolls[i+2]
			i++
		} else {
			score += bowling.rolls[i] + bowling.rolls[i+1]
			i += 2
		}
	}

	return score
}

func isStrike(bowling *Bowling, i int) bool {
	return bowling.rolls[i] == 10
}

func isSpare(bowling *Bowling, i int) bool {
	return bowling.rolls[i]+bowling.rolls[i+1] == 10
}
