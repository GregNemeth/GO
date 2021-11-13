package magic8ball

import (
	"fmt"
	"math/rand"
	"time"
)

func magic8ball() {
	rand.Seed(time.Now().UnixNano())
	answers := [8]string{
		"Not today!",
		"Perhaps",
		"I'm not sure about that",
		"Try again later!",
		"Today is definetely the day for that",
		"I'm really not feeling it!",
		"Yes, definetely",
		"Sure!",
	}
	fmt.Println("Magic8ball says:", answers[rand.Intn(len(answers))])
}
