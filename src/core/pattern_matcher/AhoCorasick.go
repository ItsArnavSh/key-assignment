package patternmatcher

// AhoCorasick struct encapsulates the automaton state
type AhoCorasick struct {
	goTo         [][]int // Trie transitions
	fail         []int   // Failure links
	output       []int   // Output bitmasks indicating matched patterns
	words        []string
	states       int
	alphabetSize int
}

// NewAhoCorasick constructor builds trie and automaton for given patterns
func NewAhoCorasick(words []string) *AhoCorasick {
	const maxChar = 256 // can be set to 26 for lowercase a-z or extend for ASCII
	maxStates := 0
	for _, w := range words {
		maxStates += len(w)
	}

	ac := &AhoCorasick{
		goTo:         make([][]int, maxStates+1),
		fail:         make([]int, maxStates+1),
		output:       make([]int, maxStates+1),
		words:        words,
		states:       1,
		alphabetSize: maxChar,
	}

	// Initialize goTo arrays with -1 (no transition)
	for i := range ac.goTo {
		ac.goTo[i] = make([]int, ac.alphabetSize)
		for j := 0; j < ac.alphabetSize; j++ {
			ac.goTo[i][j] = -1
		}
	}

	// Build the trie (goTo function)
	for i, word := range words {
		currentState := 0
		for _, ch := range word {
			c := int(ch)
			if ac.goTo[currentState][c] == -1 {
				ac.goTo[currentState][c] = ac.states
				ac.states++
			}
			currentState = ac.goTo[currentState][c]
		}
		// Mark output at terminal state for this word
		ac.output[currentState] |= (1 << i)
	}

	// Set failure function and extend trie to automaton using BFS
	for i := 0; i < ac.alphabetSize; i++ {
		if ac.goTo[0][i] == -1 {
			ac.goTo[0][i] = 0
		}
	}

	ac.fail[0] = 0
	queue := []int{}
	for i := 0; i < ac.alphabetSize; i++ {
		next := ac.goTo[0][i]
		if next != 0 {
			ac.fail[next] = 0
			queue = append(queue, next)
		}
	}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		for i := 0; i < ac.alphabetSize; i++ {
			next := ac.goTo[state][i]
			if next == -1 {
				continue
			}
			queue = append(queue, next)

			failure := ac.fail[state]
			for ac.goTo[failure][i] == -1 {
				failure = ac.fail[failure]
			}
			ac.fail[next] = ac.goTo[failure][i]
			ac.output[next] |= ac.output[ac.fail[next]]
		}
	}

	return ac
}

// Search runs the automaton over input text and returns matches as map[word]positions
func (ac *AhoCorasick) Search(text string) map[string][]int {
	result := make(map[string][]int)
	state := 0

	for i, ch := range text {
		c := int(ch)
		for ac.goTo[state][c] == -1 {
			state = ac.fail[state]
		}
		state = ac.goTo[state][c]

		if ac.output[state] == 0 {
			continue
		}
		for iWord := 0; iWord < len(ac.words); iWord++ {
			if (ac.output[state] & (1 << iWord)) != 0 {
				startIndex := i - len(ac.words[iWord]) + 1
				if startIndex >= 0 { // sanity check
					result[ac.words[iWord]] = append(result[ac.words[iWord]], startIndex)
				}
			}
		}
	}
	return result
}
