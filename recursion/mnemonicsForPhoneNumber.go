package recursion

import (
	"strconv"
	"strings"
)

var mapping = []string {"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz" }
var mnemonic []string
var partialMnemonic []string

func ComputeAllMnenonicsForPhoneNumber(phoneNumber string) []string {

	for i:= len(phoneNumber); i>0; i-- {
		partialMnemonic = append(partialMnemonic, "0")
	}

	computeMnemonicsHelper(0, phoneNumber)

	return mnemonic
}

func computeMnemonicsHelper(digit int, phoneNumber string) {

	if digit == len(phoneNumber) {
		mnemonic = append(mnemonic, strings.Join(partialMnemonic, ""))
	} else {
		idx, _ := strconv.Atoi(string(phoneNumber[digit]))
		for _, c := range mapping[idx] {
			partialMnemonic[digit] = string(c)
			computeMnemonicsHelper(digit + 1, phoneNumber)
		}
	}
}