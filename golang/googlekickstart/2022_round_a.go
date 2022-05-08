package googlekickstart

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func SpeedTyping() {
	const maxCapacity = 1024 * 1024
	var buf = make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	numTestCases, err := strconv.Atoi(ReadLine())
	PrintError(err)

	for i := 0; i < numTestCases; i++ {

		result, err := NumLettersToDelete(ReadLine(), ReadLine())
		if err != nil {
			fmt.Printf("Case #%v: IMPOSSIBLE\n", i+1)
		} else {
			fmt.Printf("Case #%v: %v\n", i+1, result)
		}
	}

}

func ReadLine() string {
	scanner.Scan()
	line := scanner.Text()
	return line
}

func PrintError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

// How many letters to delete from p to get i
func NumLettersToDelete(target, p string) (int, error) {
	if len(target) > len(p) {
		return 0, errors.New("IMPOSSIBLE")
	}

	if len(target) == len(p) && target != p {
		return 0, errors.New("IMPOSSIBLE")
	}

	if target == p {
		return 0, nil
	}

	// Loop over target and p, if the elements are not equal move along p and add to count
	i, j, result := 0, 0, 0
	for i = 0; i < len(target); i++ {

		for j < len(p) && target[i] != p[j] {
			result++
			j++
			if j == len(p) {
				// ran out of letters to try
				return 0, errors.New("IMPOSSIBLE")
			}
		}
		if i == len(target)-1 && j == len(p)-1 {
			return result, nil
		}
		j++
	}

	// left over letters res
	if j < len(p) {
		result += len(p[j:])
	} else {
		return 0, errors.New("IMPOSSIBLE")
	}

	return result, nil
}
