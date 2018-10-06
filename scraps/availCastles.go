var fenToAvailCastles = map[string]string{
	"K":    "K---",
	"KQ":   "KQ--",
	"KQk":  "KQk-",
	"KQkq": "KQkq",
	"Q":    "-Q--",
	"Qk":   "-Qk-",
	"Qkq":  "-Qkq",
	"k":    "--k-",
	"kq":   "--kq",
	"q":    "---q",
	"-":    "-",
}

var availCastlesToFen = map[string]string{
	"K---": "K",
	"KQ--": "KQ",
	"KQk-": "KQk",
	"KQkq": "KQkq",
	"-Q--": "Q",
	"-Qk-": "Qk",
	"-Qkq": "Qkq",
	"--k-": "k",
	"--kq": "kq",
	"---q": "q",
	"-":    "-",
}

// based off https://stackoverflow.com/questions/24893624/how-to-replace-a-letter-at-a-specific-index-in-a-string-in-go
func replaceAtIndex(in string, l string, a []int) string {
	out := []rune(in)
	for _, i := range a {
		out[i] = []rune(l)[0]
	}
	return string(out)
}
