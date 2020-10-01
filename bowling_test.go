package go_bowling

import "testing"


func Test_gutter_balls(t *testing.T) {
	bowling := Bowling{}
	rollNTimes(&bowling, 20, 0)
	verifyScoreIs(t, &bowling, 0)
}

func Test_all_ones(t *testing.T) {
	bowling := Bowling{}
	rollNTimes(&bowling, 20, 1)
	verifyScoreIs(t, &bowling, 20)
}

func Test_spare(t *testing.T) {
	bowling := Bowling{}

	bowling.roll(5)
	bowling.roll(5)
	bowling.roll(1)
	rollNTimes(&bowling, 17, 0)
	verifyScoreIs(t, &bowling, 12)
}

func Test_strike(t *testing.T) {
	bowling := Bowling{}

	bowling.roll(10)
	bowling.roll(2)
	bowling.roll(1)
	rollNTimes(&bowling, 17, 0)
	verifyScoreIs(t, &bowling, 16)
}

func Test_all_strikes(t *testing.T) {
	bowling := Bowling{}

	rollNTimes(&bowling, 12, 10)
	verifyScoreIs(t, &bowling, 300)
}

func verifyScoreIs(t *testing.T, bowling *Bowling, totalScore int) {
	result := bowling.score()
	if result != totalScore {
		t.Errorf("Name: Expected %d but got %d", totalScore, result)
	}
}

func rollNTimes(bowling *Bowling, times int, pins int) {
	for i := 0; i < times; i++ {
		bowling.roll(pins)
	}
}


