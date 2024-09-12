package service

import (
	"matchmaking/internal/entity"
	"math"
)

func calculateRatingChange(winnerRating, loserRating int) int {
	var res int
	coef := math.Abs(float64(max(winnerRating, loserRating)) / float64(min(winnerRating, loserRating)))
	res = int(baseMmrChange * coef)
	return res
}

func findLeastDiff(newUser entity.User, users ...entity.User) entity.User {
	var res entity.User
	res.TgID = -1
	currDiff := okDiff

	for _, user := range users {
		floatDiff := float64(newUser.Rating - user.Rating)
		if int(math.Abs(floatDiff)) <= currDiff {
			currDiff = newUser.Rating - user.Rating
			res = user
		}
	}
	return res
}
