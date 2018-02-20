package twelve

import (
	"bytes"
	"fmt"
)

var ordinals = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Verse(n int) string {
	var gift string
	if n > 1 {
		for i := n; i > 1; i-- {
			gift += gifts[i-1] + ", "
		}
		gift += "and "
	}
	gift += gifts[0]

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me, %s.", ordinals[n-1], gift)
}

func Song() string {
	buf := bytes.NewBufferString("")
	for i := 1; i <= 12; i++ {
		buf.WriteString(Verse(i))
		buf.WriteRune('\n')
	}

	return buf.String()
}
