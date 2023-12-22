package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount int

type Store struct {
	entrySeparator string
}

func (s *Store) SaveToFile(file string, j Journal) {
	ioutil.WriteFile(file, []byte(j.String(s.entrySeparator)), 0x644)
}

type Journal struct {
	entries []string
}

func (j *Journal) String(sep string) string {
	return strings.Join(j.entries, sep)
}

func (j *Journal) AddEntry(e string) int {
	entryCount++
	fmt.Println("Adding entry", e)
	j.entries = append(j.entries, e)
	return entryCount
}

func (j *Journal) RemoveEntry() {
	// TBD
}

func main() {

	fmt.Println("SRP main...")

	j := Journal{}
	j.AddEntry("i am sick")
	j.AddEntry("i ate a bug")

	fmt.Println(j.String("\n"))

	// to save
	s := Store{entrySeparator: "\n"}
	s.SaveToFile("srp.txt", j)
}
