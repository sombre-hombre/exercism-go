package protein

// FromCodon translate codon into protein
func FromCodon(codon string) string {
	switch codon {
	case "AUG":
		return "Methionine"
	case "UUU", "UUC":
		return "Phenylalanine"
	case "UUA", "UUG":
		return "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine"
	case "UAU", "UAC":
		return "Tyrosine"
	case "UGU", "UGC":
		return "Cysteine"
	case "UGG":
		return "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "STOP"
	}

	panic("Unknown codon " + codon)
}

// FromRNA translate RNA sequences into proteins
func FromRNA(protein string) []string {
	var result []string
	for i := 0; i < len(protein); i += 3 {
		p := FromCodon(protein[i : i+3])
		if p == "STOP" {
			break
		}
		result = append(result, p)
	}

	return result
}
