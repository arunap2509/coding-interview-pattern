package recursion

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	// eculidien algorithm
	// if x or y is 0 return the other one
	// we dont have to check both explictly as when we do modulus and swapping x and y it will get checked
	// GCD(A, B) => A = B * Q + R, GCD(A, B) becomes GCD(B, R) in the next recursive call 
	return Gcd(y, x % y)
}