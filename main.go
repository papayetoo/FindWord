package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	word := os.Args[1]

	// 입력 파일 읽기 1
	// dat, err := ioutil.ReadFile("./lorem.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Print(string(dat))

	// 입력 파일 읽기 2
	fileName := os.Args[2]
	f, err := os.Open(fileName)
	checkError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0
	count := 0
	lines := map[int]int{}
	for scanner.Scan() {
		s := scanner.Text()
		words := strings.Split(s, " ")
		for _, w := range words {
			fmt.Println(w)
			if w == word {
				count++
			}
		}
		lines[i] = count
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for k, v := range lines {
		fmt.Printf("find %s at line %d is %d\n", word, k, v)
	}
}
