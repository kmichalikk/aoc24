package day22

import (
	"fmt"
	"strconv"
	"sync"
)

type Day22 struct {
	numbers []int64
}

func (d *Day22) Init(lines []string) {
	for _, line := range lines {
		number, _ := strconv.ParseInt(line, 10, 64)
		d.numbers = append(d.numbers, number)
	}
}

func (d *Day22) SolveSimple() string {
	var total int64 = 0
	for _, number := range d.numbers {
		secret := number
		for range 2000 {
			secret = ((secret << 6) ^ secret) & (2<<23 - 1)
			secret = ((secret >> 5) ^ secret) & (2<<23 - 1)
			secret = ((secret << 11) ^ secret) & (2<<23 - 1)
		}

		total += secret
	}

	return fmt.Sprint(total)
}

func (d *Day22) SolveAdvanced() string {
	var maxIncome int64 = 0
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(19)
	for i := -9; i <= 9; i++ {
		go func(icode int64) {
			for j := -9; j <= 9; j++ {
				jcode := icode*100 + 10 + int64(j)
				for k := -9; k <= 9; k++ {
					kcode := jcode*100 + 10 + int64(k)
					for m := -9; m <= 9; m++ {
						mcode := kcode*100 + 10 + int64(m)
						var total int64 = 0
						for _, number := range d.numbers {
							secret := number

							var code int64 = 0
							var last int64 = secret % 10
							for range 4 {
								secret = ((secret << 6) ^ secret) & (2<<23 - 1)
								secret = ((secret >> 5) ^ secret) & (2<<23 - 1)
								secret = ((secret << 11) ^ secret) & (2<<23 - 1)
								code = code*100 + 10 + (secret%10 - last)
								last = secret % 10
							}

							for range 1996 {
								if code == mcode {
									total += last
									break
								}

								secret = ((secret << 6) ^ secret) & (2<<23 - 1)
								secret = ((secret >> 5) ^ secret) & (2<<23 - 1)
								secret = ((secret << 11) ^ secret) & (2<<23 - 1)

								code = (code%1000000)*100 + 10 + (secret%10 - last)
								last = secret % 10
							}
						}

						lock.Lock()
						if total > maxIncome {
							maxIncome = total
						}
						lock.Unlock()
					}
				}
			}
			wg.Done()
		}(10 + int64(i))
	}

	wg.Wait()

	return fmt.Sprint(maxIncome)
}
