package main

import "fmt"

type Player struct {
	elo      int
	fullname string
}

func main() {
	rating := []Player{
		{600, "Abylay Amandykov"}, //left1
		{1799, "Nur-Assyl Kanagatov"},
		{1800, "Aigali Sultankul"},
		{2400, "Ian Nepomnyashi"}, //mid1

		{2400, "Hikaru Youtuber"}, //left2
		{2510, "Sayat Zeitkazy"},  //mid2 //right2 //left3

		{2700, "Magnus Karlsen"},
		{2710, "Sanzhar Myrzash"}, //right1
	}
	legend := Player{
		2500, "Harry Kasparov",
	}
	fmt.Println(spotPlayer(rating, legend))
}

func spotPlayer(rating []Player, player Player) int {
	left, right := 0, len(rating)-1
	var middle int
	for left < right {
		middle = (left + right) / 2
		fmt.Println(middle)
		if rating[middle].elo < player.elo {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}
