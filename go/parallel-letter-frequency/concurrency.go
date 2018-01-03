//Package letter implements parallel word frequency
package letter

/*ConcurrentFrequency counts word occurances using parallel processes*/
func ConcurrentFrequency(wordList []string) FreqMap {
	// Map the frequency function over all the words
	// create a channel named channel
	//	channel := make(chan FreqMap, len(wordList))
	second := make(chan FreqMap, len(wordList))

	// loop through wordlist

	for _, words := range wordList {
		go parallelFreq(words, second)
	}

	// Reduce the results to a single map
	frequency := FreqMap{}

	for range wordList {
		for t1, t2 := range <-second {
			frequency[t1] += t2
		}
	}

	return frequency
}

//function parallelFreq is the function that is called repeatedly to figure out frequency
func parallelFreq(input string, outlet chan FreqMap) {
	outlet <- Frequency(input)
}
