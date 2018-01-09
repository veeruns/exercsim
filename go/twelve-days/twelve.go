//package twelve returns the twelve days of christmas lyrics
package twelve

import (
	"bytes"
	"fmt"
)

//twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree
var gifts = map[int][2]string{1: {"first", "a Partridge in a Pear Tree"},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

//Verse function returns line of song based on input number
func Verse(input int) string {
	if input > 12 {
		return "Input cannot be greater than 12"
	}
	header := fmt.Sprintf("On the %s day of Christmas my true love gave to me", gifts[input][0])
	var op bytes.Buffer
	op.WriteString(header)
	op.WriteString(", ")
	for i := input; i >= 1; i-- {

		//op.WriteString(gifts[i][1])
		if i != 1 {
			op.WriteString(gifts[i][1])
			op.WriteString(", ")
		} else {
			if input != 1 {
				op.WriteString("and ")
				op.WriteString(gifts[i][1])
				op.WriteString(".")
			} else {
				op.WriteString(gifts[i][1])
				op.WriteString(".")
			}
		}
	}
	return op.String()
}

//Song function returns whole song
func Song() string {
	var op bytes.Buffer
	for i := 1; i <= 12; i++ {
		op.WriteString(Verse(i))
		op.WriteString("\n")

	}
	return op.String()
}
