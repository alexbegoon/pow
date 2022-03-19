package quotes

import "math/rand"

var quotes = []string{
	"The fool doth think he is wise, but the wise man knows himself to be a fool.",
	"It is better to remain silent at the risk of being thought a fool, than to talk and remove all doubt of it.",
	"Whenever you find yourself on the side of the majority, it is time to reform (or pause and reflect).",
}

func GetQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
