package service

import "math"

func (s service) calculateRatingChange(winnerRating, loserRating int) int {
	var res int
	coef := math.Abs(float64(max(winnerRating, loserRating)) / float64(min(winnerRating, loserRating)))
	res = int(baseMmrChange * coef)
	return res
}
