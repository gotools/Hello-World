// reverse.go

package reverse

//reverse text
func Reverse(s string) string {
	runes := []int(s)
	n := len(runes)
	h := n/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
