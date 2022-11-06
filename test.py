from typing import List

mapping = ("0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz" )

def phone_mnemonic(phone_num: str) -> List[str]:
    def phone_mnemonic_helper(digit: int) -> None:
        if digit == len(phone_num):
            print("partial when appending: ", partial_mnemonic)
            mnemonics.append(''.join(partial_mnemonic))
        else:
            for c in mapping[int(phone_num[digit])]:
                # print(c)
                partial_mnemonic[digit] = c
                # print("partial when updated: ", partial_mnemonic)
                phone_mnemonic_helper(digit + 1)
    
    mnemonics: List[str] = []
    partial_mnemonic = ['0'] * len(phone_num)
    # print("partial when created: ", partial_mnemonic)
    phone_mnemonic_helper(0)
    return mnemonics

print(phone_mnemonic("123"))