// go build namegen.go
// USAGE: namegen.exe -length=5 -vowelsPos=2,4 -number=50 -random=false -alphabetmode=false > out.txt

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type interval []int

var (
	random                       *bool
	alphabetmode                 *bool
	length                       *int
	number                       *float64
	vowelsPos                    interval
	vowels, consonants, alphabet []string
)

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

func (i *interval) Set(value string) error {
	// If we wanted to allow the flag to be set multiple times,
	// accumulating values, we would delete this if statement.
	// That would permit usages such as
	//	-deltaT 10s -deltaT 15s
	// and other combinations.
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}

	for _, num := range strings.Split(value, ",") {
		value, err := strconv.Atoi(num)

		if err != nil {
			return err
		}

		*i = append(*i, value)
	}

	return nil
}

func isIntInInterval(i int, list interval) bool {
	for _, b := range list {
		if b == i {
			return true
		}
	}
	return false
}

func isStringInSlice(s string, list []string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

func doesUserConfirm() bool {
	var response string

	_, err := fmt.Scanln(&response)

	if err != nil {
		log.Fatal(err)
	}

	okResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	koResponses := []string{"n", "N", "no", "No", "NO"}

	if isStringInSlice(response, okResponses) {
		return true
	} else if isStringInSlice(response, koResponses) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return doesUserConfirm()
	}
}

func genRndNames(totNames float64) {
	if *alphabetmode {
		for i := 0.0; i < totNames; i++ {
			name := ""

			pos := 1
			for pos <= *length {
				if isIntInInterval(pos, vowelsPos) {
					name = name + vowels[rand.Intn(5)]
				} else {
					name = name + alphabet[rand.Intn(26)]
				}

				pos++
			}

			if name == "" {
				return
			}

			fmt.Println(name)
		}
	} else {
		for i := 0.0; i < totNames; i++ {
			name := ""

			pos := 1
			for pos <= *length {
				if isIntInInterval(pos, vowelsPos) {
					name = name + vowels[rand.Intn(5)]
				} else {
					name = name + consonants[rand.Intn(21)]
				}

				pos++
			}

			if name == "" {
				return
			}

			fmt.Println(name)
		}
	}
}

func genNames(totNames float64) {
	charpos := make([]int, *length)

	//fmt.Println(totNames)
	//fmt.Println(charpos)

	pointer := len(charpos) - 1

	if *alphabetmode {
		for i := 0.0; i < totNames; i++ {
			name := ""

			pos := 0
			for pos < *length {
				name = name + alphabet[charpos[pos]]

				pos++
			}

			if name == "" {
				return
			}

			fmt.Println(name)

			incrementPositions(charpos, pointer)
		}
	} else {
		for i := 0.0; i < totNames; i++ {
			name := ""

			pos := 0
			for pos < *length {
				if isIntInInterval(pos+1, vowelsPos) {
					name = name + vowels[charpos[pos]]
				} else {
					name = name + consonants[charpos[pos]]
				}

				pos++
			}

			if name == "" {
				return
			}

			fmt.Println(name)

			incrementPositions(charpos, pointer)
		}
	}
}

func incrementPositions(charpos []int, pointer int) {
	//fmt.Println(pointer)

	if pointer < 0 {
		return
	}

	if !(*alphabetmode) && isIntInInterval(pointer+1, vowelsPos) && charpos[pointer] < 4 {
		charpos[pointer]++
	} else if !(*alphabetmode) && isIntInInterval(pointer+1, vowelsPos) && charpos[pointer] >= 4 {
		pointer--
		incrementPositions(charpos, pointer)
		pointer++
		charpos[pointer] = 0
	} else if !(*alphabetmode) && !isIntInInterval(pointer+1, vowelsPos) && charpos[pointer] < 20 {
		charpos[pointer]++
	} else if !(*alphabetmode) && !isIntInInterval(pointer+1, vowelsPos) && charpos[pointer] >= 20 {
		pointer--
		incrementPositions(charpos, pointer)
		pointer++
		charpos[pointer] = 0
	} else if (*alphabetmode) && charpos[pointer] < 25 {
		charpos[pointer]++
	} else if (*alphabetmode) && charpos[pointer] >= 25 {
		pointer--
		incrementPositions(charpos, pointer)
		pointer++
		charpos[pointer] = 0
	}
}

func init() {
	// Tie the command-line flag to the intervalFlag variable and
	// set a usage message.
	flag.Var(&vowelsPos, "vowelsPos", "Comma-separated list of integer numbers that define the vowel positions in strings/names (Example: -vowelPos=2,4,6).")

	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	random = flag.Bool("random", true, "Enable random strings/names generation (Default: true). Ignored if -number=0.")
	alphabetmode = flag.Bool("alphabetmode", false, "Enable strings/names generation using whole alphabet (Default: false).")
	length = flag.Int("length", 6, "Integer number that defines strings/names characters length (Default: 6).")
	number = flag.Float64("number", 10, "Integer number indicating the number of strings/names to be generated (Default: 10, 0 for all the possible combinations).")

	flag.Parse()

	vowels = []string{"a", "e", "i", "o", "u"}
	consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	fmt.Println("-------------------------------------------------------")
	fmt.Println("---- STRINGS/NAMES GENERATOR BY Francesco Giglioli ----")
	fmt.Println("----------------- github.com/giglio91 -----------------")
	fmt.Println("-------------------------------------------------------\n")
	//fmt.Println(vowels)
	//fmt.Println(consonants)
	if *number == 0 {
		fmt.Println("Random Mode: DISABLED (-number=0)")
	} else {
		fmt.Println("Random Mode: ", *random)
	}
	fmt.Println("Alphabet Mode:", *alphabetmode)
	fmt.Println("Strings/names length: ", *length)
	fmt.Println("Vowels positions: ", vowelsPos)

	var totNames float64
	if *alphabetmode {
		totNames = math.Pow(26, float64(*length))
	} else {
		totNames = math.Pow(5, float64(len(vowelsPos))) * math.Pow(21, float64(*length-len(vowelsPos)))
	}

	fmt.Printf("Number of possible combinations: %.0f\n", totNames)

	if *number == 0 {
		if totNames > 200000.0 {
			fmt.Println("I will generate a HUGE amount of strings/names. It will take time and resources. Do you want to proceed? [yes/no]")

			if doesUserConfirm() {
				genNames(totNames)
			}
		} else {
			fmt.Println("")
			genNames(totNames)
		}
	} else {
		fmt.Println("Number of strings/names actually generated: ", *number, "\n")

		if *random {
			genRndNames(*number)
		} else {
			genNames(*number)
		}
	}
}
