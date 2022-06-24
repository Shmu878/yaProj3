package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nodA := 0
	nodB := 0
	var numbers = make(map[int][]int)

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	size, err := strconv.Atoi(in.Text())
	if err != nil {
		fmt.Println("Ошибка конвертации.", err)
	}
	for i := 0; i < size; i++ {
		var nn []int
		in.Scan()
		st := strings.Split(in.Text(), " ")
		for s := 0; s < len(st); s++ {
			n, er := strconv.Atoi(st[s])
			if er != nil {
				fmt.Println("Ошибка конвертации.", er)
			}
			nn = append(nn, n)
		}
		numbers[i] = nn
	}

	for i := 0; i < size; i++ {
		n := numbers[i]

		for z := 1; z <= n[0]; z++ {
			if prime(z) == true {
				nd := nod(n[0], n[1]*z)
				if nodA < nd {
					nodA = nd
				}
			}
		}
		for z := 1; z <= n[1]; z++ {
			if prime(z) == true {
				nd := nod(n[0]*z, n[1])
				if nodB < nd {
					nodB = nd
				}
			}
		}
		if nodA > nodB {
			fmt.Println(nodA)
		} else {
			fmt.Println(nodB)
		}
	}
}

func nod(a int, b int) int {
	var c int
	if a < b {
		for c = a; b%c != 0 && a%c != 0; c-- {
		}
	} else {
		for c = b; b%c != 0 && a%c != 0; c-- {
		}
	}
	return c
}

func prime(a int) bool {
	if a == 1 {
		return false
	}
	for i := 2; a%i == 0; i++ {
		if i != a {
			return false
		}
	}
	return true
}
