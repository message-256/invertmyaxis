package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type question struct {
	a, b int
	opp string
}
func add(x,y int) int {
	return x + y;
}
func mul(x,y int) int {
	return x * y;
}
func sub(x,y int) int {
	return x - y;
}
func mod(x,y int) int {
	return x % y;
}
func and(x,y int) int {
	return x & y;
}
func xor(x,y int) int {
	return x ^ y;
}
func or(x,y int) int {
	return x | y;
}

func main() {
	var previous []question
	var ans int
	var current question
	var input string
	var err error
	var contflag bool
	var modes = map[string] func (int,int) int{
			"+":add,
			"-":sub,
			"^":xor,
			"*":mul,
			"&":and,
			"%":mod,
	}
	var mode string = "*"
	var opprolled int;
	for input != "q" {
		ans = 0
		contflag = false
		if len(mode) > 1 {
			opprolled = rand.Intn(len(mode)-1)
		} else {
			opprolled = 0
		}
		current.a = rand.Intn(11) + 1
		current.b = rand.Intn(11) + 1
		//now its expected
		if current.a < current.b {
			current.a, current.b = current.b, current.a
		}
		for i := range previous {
			if current.b == previous[i].b && current.a == previous[i].a && current.opp == previous[i].opp {
				contflag = true
				break
			}
		}
		if contflag {
			fmt.Println(current.a, current.opp, current.b, "not unique")
			continue
		}
		previous = append(previous, current)
		for {
			fmt.Println(current.a, string(mode[opprolled]), current.b, "=")
			fmt.Scanf("%s", &input)
			if input == "q" {
				os.Exit(0)
			} else if input == "mode"{
				var ok bool
				for {
					ok = true;
					for i := range modes {
						fmt.Printf("%s",i)
					}
					fmt.Printf("mode(current:%s):",mode)
					fmt.Scanf("%s",&input)
					if input == "help"{
						fmt.Println("choose mode eg: *&+-")
						fmt.Println("then all modes chosen will be added to the deck")
						continue
					} else if input == "quit" || input == "q"  {
						os.Exit(0);
					} else if input == "keep" {
						break;
					}
					for i := range input {
						_,ok = modes[string(input[i])]
						if !ok {
							break;
						}
					}
					if !ok {
						continue
					} else {
						mode = input;
						break
					}
					
				}
				continue
			} else if input == "reset" {
				fmt.Println("clearing deck")
				previous = previous[:0]
				
			} else if input == "continue" {
				previous = previous[:len(previous)-1]
				break
			} else if ans, err = strconv.Atoi(input); err != nil {
				fmt.Println("not a number :(real error)", err)
				continue
			}
			if ans != modes[string(mode[opprolled])](current.b,current.a){
				fmt.Println(ans,"is wrong")
				continue
			} else {
				break
			}
		}
	}

}
