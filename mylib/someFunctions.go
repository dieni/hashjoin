package mylib

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"sync"
	"time"
	"math/rand"
	"hash/fnv"
)

type JoinElement struct {
	StudentId  int
	CourseId   string
	CourseName string
}

var tablesPath = "src/github.com/dieni/hashjoin/tables/"

func SaveCourseTable(content []Course, filename string) {
	encContent, err := json.Marshal(content)
	check(err)

	err = ioutil.WriteFile(path.Join(tablesPath, filename), encContent, 0664)
	check(err)
}

func ReadCourseTable(filename string) []Course {
	encContent, err := ioutil.ReadFile(path.Join(tablesPath, filename))
	check(err)
	content := make([]Course, 0)
	err = json.Unmarshal(encContent, &content)
	return content
}

func SaveStudentTable(content []Student, filename string) {
	encContent, err := json.Marshal(content)
	check(err)

	err = ioutil.WriteFile(path.Join(tablesPath, filename), encContent, 0664)
	check(err)
}

func ReadStudentTable(filename string) []Student {
	encContent, err := ioutil.ReadFile(path.Join(tablesPath, filename))
	check(err)
	content := make([]Student, 0)
	err = json.Unmarshal(encContent, &content)
	return content
}

//Generates a random number within a given range
func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Contains method for slice of type "Course"
func containsCourse(c []string, e string) bool {
	for _, a := range c {
		if a == e {
			return true
		}
	}
	return false
}

//Contains method for slice of type "Student"
func containsStudent(s []Student, e int) bool {
	for _, a := range s {
		if a.Id == e {
			return true
		}
	}
	return false
}

func GenerateHashvalue(value string) int {
	h := fnv.New32()
	h.Write([]byte(value))
	return int(h.Sum32())
}

func Merge(cs []<-chan JoinElement) <-chan JoinElement {
	var wg sync.WaitGroup
	out := make(chan JoinElement)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan JoinElement) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
