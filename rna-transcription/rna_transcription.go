package strand

import (
	"strings"
)

// ToRNA translates a dna strand to a rna strand
func ToRNA(dna string) string {
	re := strings.NewReplacer("C", "G", "G", "C", "T", "A", "A", "U")
	return re.Replace(dna)
}
