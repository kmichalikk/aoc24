package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct {
	lines []string
}

func (d *Day2) Init(lines []string) {
	d.lines = lines
}

func (d *Day2) SolveSimple() string {
	safe := 0
	for _, line := range d.lines {
		values := strings.Split(line, " ")
		if len(values) < 2 {
			safe++
			continue
		}

		a, _ := strconv.Atoi(values[0])
		b, _ := strconv.Atoi(values[1])
		sign := 1

		if a > b {
			sign = -1
		}

		diff := (b - a) * sign

		if diff > 3 || diff < 1 {
			continue
		}

		i := 2
		ok := true
		for i < len(values) {
			a = b
			b, _ = strconv.Atoi(values[i])
			diff = (b - a) * sign
			if diff > 3 || diff < 1 {
				ok = false
				break
			}

			i++
		}

		if ok {
			safe++
		}
	}

	return fmt.Sprint(safe)
}

func (d *Day2) SolveAdvanced() string {
	safe := 0
	for _, line := range d.lines {
		valuesStr := strings.Split(line, " ")
		values := make([]int, len(valuesStr))
		wantSign := 0
		for i, s := range valuesStr {
			values[i], _ = strconv.Atoi(s)

			if i == 0 {
				continue
			}

			if values[i]-values[i-1] > 0 {
				wantSign++
			} else {
				wantSign--
			}
		}

		if wantSign < 0 {
			wantSign = -1
		} else {
			wantSign = 1
		}

		i := 0
		errors := 0
		sign, distance := signAndDistance(values[0], values[1])
		if sign != wantSign || distance > 3 || distance < 1 {
			sign2, distance2 := signAndDistance(values[0], values[2])
			errors++
			if sign2 != wantSign || distance2 > 3 || distance2 < 1 {
				i += 2
			} else {
				i += 3
			}
		} else {
			i += 2
		}

		for i < len(values) {
			sign, distance = signAndDistance(values[i-1], values[i])
			if sign != wantSign || distance > 3 || distance < 1 {
				errors++
				if i+1 < len(values) {
					sign2, distance2 := signAndDistance(values[i-1], values[i+1])
					if sign2 != wantSign || distance2 > 3 || distance2 < 1 {
						errors++
					}
				}
				i++
			} else if i+1 < len(values) {
				sign2, distance2 := signAndDistance(values[i-1], values[i+1])
				if sign2 == wantSign && distance2 < distance && distance2 > 0 {
					errors++
					i++
				}
			}

			i++
		}

		if errors < 2 {
			safe++
		}

		//
		//a, _ := strconv.atoi(values[0])
		//b, _ := strconv.atoi(values[1])
		//sign, distanceab := signanddistance(a, b)
		//
		//if len(values) == 2 {
		//	if distanceab <= 3 {
		//		safe++
		//	}
		//
		//	continue
		//}

		//
		//c, _ := strconv.Atoi(values[2])
		//signBC, distanceBC := signAndDistance(b, c)
		//
		//if len(values) == 3 {
		//	_, distanceAC := signAndDistance(a, c)
		//
		//	if distanceAB <= 3 || distanceAC <= 3 || distanceBC <= 3 {
		//		safe++
		//	}
		//
		//	continue
		//}
		//
		//d, _ := strconv.Atoi(values[3])
		//signCD, _ := signAndDistance(c, d)
		//
		//wantSign := signAB + signBC + signCD
		//if wantSign > 0 {
		//	wantSign = 1
		//} else {
		//	wantSign = -1
		//}
		//
		//i := 3
		//errors := 0
		//for i < len(values)-1 {
		//	signAB, distanceAB = signAndDistance(a, b)
		//	signBC, distanceBC = signAndDistance(b, c)
		//	if signAB != wantSign || signBC != wantSign || distanceAB > 3 || distanceBC > 3 {
		//		errors++
		//		// b is probably unsafe, check a against c
		//		signAC, distanceAC := signAndDistance(a, c)
		//		if signAC != wantSign || distanceAC > 3 {
		//			// a is actually unsafe
		//			a, b = b, c
		//		} else {
		//			a, b = a, c
		//		}
		//	} else {
		//		a, b = b, c
		//	}
		//	c, _ = strconv.Atoi(values[i])
		//	i++
		//}
		//
		//signAB, distanceAB = signAndDistance(a, b)
		//signBC, distanceBC = signAndDistance(b, c)
		//if signAB != wantSign || signBC != wantSign || distanceAB > 3 || distanceBC > 3 {
		//	errors++
		//}
		//
		//if errors < 2 {
		//	safe++
		//}
	}

	return fmt.Sprint(safe)
}

func signAndDistance(a, b int) (int, int) {
	if diff := b - a; diff > 0 {
		return 1, diff
	} else {
		return -1, -diff
	}
}
