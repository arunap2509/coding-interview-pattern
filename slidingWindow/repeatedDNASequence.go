package slidingwindow

func FindRepeatedDnaSequences(dna string, seqLength int) []string {
	pointer := 0
	maxSequenceSeen := make(map[string]int)
	maxSequenceRepeated := make([]string, 0)

	for pointer + seqLength <= len(dna) {

		sequence := dna[pointer:pointer+seqLength]
		maxSequenceSeen[sequence]++

		if maxSequenceSeen[sequence] == 2 {
			maxSequenceRepeated = append(maxSequenceRepeated, sequence)
		}

		pointer++
	}

	return maxSequenceRepeated
}