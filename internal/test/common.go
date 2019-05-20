package test

import (
	"github.com/ivost/nixug/internal/models"
	"log"
	"os"
)

// appendToFile If the file doesn't exist, create it, append to the file
func AppendToFile(fileName string, data string) error {
	//log.Print("AppendToFile")
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if check(err) {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(data))
	//log.Printf("written %v bytes", n)
	_, _ = f.Write([]byte("\n"))
	return err
}

func NewTestGroups() []models.Group {
	// must be sorted by name
	g := []models.Group{
		{Name: "adm", GID: 4, Members: []string{"syslog", "foo"}},
		{Name: "log", GID: 42, Members: []string{"foo", "bar"}},
		{Name: "root", GID: 0},
	}
	return g
}

func check(err error) bool {
	if err == nil {
		return false
	}
	log.Print(err.Error())
	return true
}
