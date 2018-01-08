//package tournament implements Tournament
package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

func Tally(r io.Reader, buffer io.Writer) error {
	p := make([]byte, 512)
	r.Read(p)
	tally := make(map[string][5]int)
	line := strings.Split((string)(p), "\n")
	for _, value := range line {
		if value != "" {
			match := strings.Split(value, ";")
			if len(match) == 3 {
				//fmt.Printf("Length of match is %s %d\n", match, len(match))
				team1 := match[0]
				team2 := match[1]
				result := match[2]
				if strings.Compare(result, "win") != 0 && strings.Compare(result, "loss") != 0 && strings.Compare(result, "draw") != 0 {
					//fmt.Printf("Result is %s %d\n", result, strings.Compare(result, "win"))
					return errors.New("Bad result")
				}

				if result == "win" {
					c := tally[team1][0] + 1
					win := tally[team1][1] + 1
					drawn := tally[team1][2]
					loss := tally[team1][3]
					//points := win*3 + drawn*1
					tally[team1] = [5]int{c, win, drawn, loss, 0}
					cL := tally[team2][0] + 1
					winL := tally[team2][1]
					drawnL := tally[team2][2]
					lossL := tally[team2][3] + 1
					//pointsL := win*3 + drawn*1
					tally[team2] = [5]int{cL, winL, drawnL, lossL, 0}
				} else if result == "loss" {
					c := tally[team1][0] + 1
					win := tally[team1][1]
					drawn := tally[team1][2]
					loss := tally[team1][3] + 1
					//points := win*3 + drawn*1
					tally[team1] = [5]int{c, win, drawn, loss, 0}
					cL := tally[team2][0] + 1
					winL := tally[team2][1] + 1
					drawnL := tally[team2][2]
					lossL := tally[team2][3]
					//pointsL := win*3 + drawn*1
					tally[team2] = [5]int{cL, winL, drawnL, lossL, 0}
				} else if result == "draw" {
					c := tally[team1][0] + 1
					win := tally[team1][1]
					drawn := tally[team1][2] + 1
					loss := tally[team1][3]
					// := win*3 + drawn*1
					tally[team1] = [5]int{c, win, drawn, loss, 0}
					cL := tally[team2][0] + 1
					winL := tally[team2][1]
					drawnL := tally[team2][2] + 1
					lossL := tally[team2][3]
					//pointsL := win*3 + drawn*1
					tally[team2] = [5]int{cL, winL, drawnL, lossL, 0}
				}
			}
		}
	}
	//var retstring string
	//retstring = append("Team                           | MP |  W |  D |  L |  P")
	var b1 bytes.Buffer
	//  retstring = append

	for key, value := range tally {
		points := value[1]*3 + value[2]*1
		cL := tally[key][0]
		winL := tally[key][1]
		drawnL := tally[key][2]
		lossL := tally[key][3]

		tally[key] = [5]int{cL, winL, drawnL, lossL, points}

	}
	hack := map[int]string{}
	hackkeys := []int{}
	for hkeys, val := range tally {
		hack[val[4]] = hkeys
		hackkeys = append(hackkeys, val[4])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(hackkeys)))

	b1.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	//op := fmt.Sprintf("%-31s|%2s|%2s|%2s|%2s|%2s\n", "Team", "MP", "W", "D", "L", "P")

	//b1.WriteString(op)
	for _, val := range hackkeys {
		team := hack[val]
		op := fmt.Sprintf("%-31s| %2d | %2d | %2d | %2d | %2d\n", team, tally[team][0], tally[team][1], tally[team][2], tally[team][3], tally[team][4])
		b1.WriteString(op)
	}
	fmt.Println(b1.String())
	io.WriteString(buffer, b1.String())
	/*for key, value := range tally {
		points := value[1]*3 + value[2]*1
		op := fmt.Sprintf("%32s|%4d|%4d|%4d|%4d|%4d\n", key, value[0], value[1], value[2], value[3], points)
		b1.WriteString(op)
	}
	fmt.Println(b1.String())*/
	return nil
}
