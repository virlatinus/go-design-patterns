package patterns

import (
	"fmt"
	"github.com/pterm/pterm"
	"log"
	"slices"
	"strings"
)

var entryCount = 0

type SingleResponsibility struct{}

type Journal struct {
	entries []string
}

func (j *Journal) addEntry(message string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, message)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) removeEntry(index int) error {
	if index <= 0 || index > entryCount {
		return fmt.Errorf("index out of bounds")
	}
	j.entries = slices.Delete(j.entries, index-1, index)
	return nil
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (srp *SingleResponsibility) Run() {
	pterm.DefaultBox.Println("SOLID Single Responsibility Principle")

	j := new(Journal)
	j.addEntry("First entry")
	j.addEntry("Second entry")
	j.addEntry("Third entry")

	fmt.Println("Initial data:")
	fmt.Println(j)
	fmt.Println()

	fmt.Println("Removing the second entry:")
	if err := j.removeEntry(2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(j)
}
