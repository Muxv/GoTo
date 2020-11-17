package main

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// translate a number to a unique url
func genKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}

	l := len(keyChar)
	// turn a decimal number n to a l-system number
	// max-length -> 20
	s := make([]byte, 20)
	i := len(s)

	for n > 0 && i >= 0 {
		i--
		s[i] = keyChar[n%l]
		n = n / l
	}
	return string(s[i:])
}
