package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func getpath(file string) string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	filepath := dir + "\\" + file

	return filepath

}

func writefile(letterid string) {
	file := getpath("known.txt")
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		// handle the error here
		log.Fatal(err)
	}

	f.WriteString(letterid + "\n")

	defer f.Close()

}

func readfile(letterid string) bool {
	f := getpath("known.txt")
	file, err := os.Open(f)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		//fmt.Printf("%T\n", scanner.Text()) //так можно узнать тип переменной
		if scanner.Text() == letterid {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return false

}

func initfile() {
	f := getpath("known.txt")
	if _, err := os.Stat(f); os.IsNotExist(err) {
		file, err := os.Create(f)
		if err != nil {
			// handle the error here
			log.Fatal(err)
		}
		defer file.Close()

		file.WriteString("Known letters list. \n")
	}

}
