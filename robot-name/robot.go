package robotname

import (
	"bytes"
	"math/rand"
)

// Robot represents a robot with a name
type Robot struct {
	name string
}

var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var base = len(letters)
var usedNames = make(map[string]struct{})

// Name returns robot's name
func (r *Robot) Name() string {
	if r.name != "" {
		return r.name
	}

	r.name = getUniqueName()

	return r.name
}

// Reset resets robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func getUniqueName() string {
	var rnd int
	var name string
	for {
		rnd = rand.Int()
		name = generateName(rnd)
		if _, found := usedNames[name]; !found {
			usedNames[name] = struct{}{}
			break
		}
	}

	return name
}

func generateName(rnd int) string {
	buf := bytes.NewBufferString("")

	for i := 0; i < 2; i++ {
		buf.WriteByte(letters[rnd%base])
		rnd = rnd / base
	}
	for i := 0; i < 3; i++ {
		buf.WriteByte(byte('0' + rnd%10))
		rnd = rnd / 10
	}

	return buf.String()
}
