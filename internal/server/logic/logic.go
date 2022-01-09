package logic

import (
	"fmt"
)

var pastFourScores = []float64{5.0, 4.0, 2.0, 1.0}

func GetSize() float64 {
	fmt.Println(pastFourScores)
	olderScores := pastFourScores[0] + pastFourScores[1]
	recentScores := pastFourScores[2] + pastFourScores[3]

	diff := recentScores - olderScores
	if diff > 0.0 {
		size := 600.0 + diff*60.0
		if size < 2000.0 {
			return size
		}
		return 2000.0
	}
	if diff > -5.0 && diff <= 0.0 {
		return 100.0 + 18.0*diff
	}
	return 10.0
}

func Average() float64 {
	var sum float64
	for _, item := range pastFourScores {
		sum += item
	}
	average := sum / float64(len(pastFourScores))
	return average
}

func SetScore(a float64) bool {
	pastFourScores = append(pastFourScores, a)
	pastFourScores = pastFourScores[1:]
	return true
}
