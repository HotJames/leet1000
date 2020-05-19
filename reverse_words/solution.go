package reverse_words

func ReverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
