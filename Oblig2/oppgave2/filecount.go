package main



import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"sort"
)

func scanNumberOfLines() { // counts number of lines
	file,_  := os.Open(os.Args[1])
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("\nNumber of lines in file:", lineCount, "\n")
}

func scanNumberOfRunes(runes string) { // counts number of runes
	fmt.Println("Total runes in file: ", len([]rune(runes)))
}

func dupCount(list []rune) map[string] int {
	duplicateFrequency := make(map[string] int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicateFrequency[string(item)]

		if exist {
			duplicateFrequency[string(item)] += 1 // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[string(item)] = 1 // else start counting from 1
		}
	}

	return duplicateFrequency

}

func runePrinter(m map[string]int, value int){
	for k, v := range m {
		if v == value {
			fmt.Println("Rune:", []rune(k), "- Counts:", v)
			return
		}
	}
	return
}

func main() {
	input := os.Args [1]
	b, err := ioutil.ReadFile(input) // Opens file
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert to string
	rune1 := []rune(str) // convert string to rune array
	dupMap := dupCount(rune1) // create map with duplicate count, from rune array

	var keys []int // array to store the numbers
	var strings []string // array to store the strings - useless


	// collecting keys
	for v, k := range dupMap {
		//fmt.Printf("Item : %s , Count : %d\n", string(v), k)
		keys = append(keys, k)
		strings = append(strings, v)

	}

	// sorting - highest number first (the one with most occurrences being first)
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	//print name of file
	fmt.Println("\nInformation about", input)


	// print number of lines
	scanNumberOfLines()


	// print 5 most common runes
	fmt.Println("Most common runes: ")
	fmt.Println()

	var i int
	i = 0
	for v, k := range keys {
		i++
		if i <= 5 {
		runePrinter(dupMap, k)
				v=v
		}
	}

	// new line
	fmt.Println()

	// print number of runes
	scanNumberOfRunes(str)

	fmt.Println()


}