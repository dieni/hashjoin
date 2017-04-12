package main

import (
	"fmt"
	"github.com/dieni/hashjoin/mylib"
	"html/template"
	"net/http"
	"path"
	"runtime"
	"time"
)

//################################# GLOBAL VARIABLES and CONSTANTS ##############################################
var tplPath = "src/github.com/dieni/hashjoin/templates/"
var tpl *template.Template
var dataTables DataT

//################################################################################################################

//############################################### STRUCTS ########################################################
type DataT struct {
	CourseT  []mylib.Course
	StudentT []mylib.Student
	JoinedT  []mylib.JoinElement
}

type Result struct {
	Info string
	Time int
}

//################################################################################################################

//############################################## FUNCTIONS #######################################################

//This function loads and joins a range of partitions
func joinRangeofPartitions(courseHashTable [][]mylib.Course, studentHashTable [][]mylib.Student) <-chan mylib.JoinElement {
	out := make(chan mylib.JoinElement)

	go func() {

		//Join partitions in memory
		for hi, _ := range studentHashTable {
			for _, s := range studentHashTable[hi] { //For each student in partition
				for _, c := range courseHashTable[hi] { //compare with each course in same partition
					if s.CourseId == c.Id { //if the courseid is the same -> join them
						out <- mylib.JoinElement{StudentId: s.Id, CourseId: s.CourseId, CourseName: c.Name}
						break
					}
				}
			}
		}

		close(out)
	}()

	return out
}

func partitioning(courses []mylib.Course, students []mylib.Student) ([][]mylib.Course, [][]mylib.Student) {
	hashTableSize := len(courses) * 2
	courseHashTable := make([][]mylib.Course, hashTableSize)
	studentHashTable := make([][]mylib.Student, hashTableSize)

	//partitioning Courses
	for _, c := range courses {
		hv := mylib.GenerateHashvalue(c.Id) % len(courseHashTable)
		courseHashTable[hv] = append(courseHashTable[hv], c)
	}

	//partitioning Students
	for _, s := range students {
		hv := mylib.GenerateHashvalue(s.CourseId) % len(studentHashTable)
		studentHashTable[hv] = append(studentHashTable[hv], s)
	}

	return courseHashTable, studentHashTable
}

func Join(courseHashTable [][]mylib.Course, studentHashTable [][]mylib.Student, cores int) {
	hashTableSize := len(courseHashTable)
	dataTables.JoinedT = make([]mylib.JoinElement, 0) //Output Table

	runtime.GOMAXPROCS(cores)
	joinChannels := make([]<-chan mylib.JoinElement, cores)
	for i := 0; i < cores; i++ {
		joinChannels[i] = joinRangeofPartitions(courseHashTable[i*hashTableSize/cores:(i+1)*hashTableSize/cores], studentHashTable[i*hashTableSize/cores:(i+1)*hashTableSize/cores])
	}

	for n := range mylib.Merge(joinChannels) {
		dataTables.JoinedT = append(dataTables.JoinedT, n)
	}

}

func createTables() {
	mylib.SaveCourseTable(mylib.CreateCourseTable(), "courseTable.json")
	mylib.SaveStudentTable(mylib.CreateStudentTable(1000000), "studentTable1000000.json")
}

func loadTables() {
	dataTables.CourseT = mylib.ReadCourseTable("courseTable.json")
	dataTables.StudentT = mylib.ReadStudentTable("studentTable1000000.json")
}

func testHashjoin(courseHashTable [][]mylib.Course, studentHashTable [][]mylib.Student) {
		var i time.Duration = 1
		var rounds time.Duration = 3
		var mean1 time.Duration = 0
		var mean2 time.Duration = 0
		var mean4 time.Duration = 0
		var mean8 time.Duration = 0
		var mean16 time.Duration = 0
		var mean32 time.Duration = 0

	for ; i <= rounds; i++ {
		start := time.Now()
		Join(courseHashTable, studentHashTable, 1)
		elapsed := time.Since(start)
		mean1 += elapsed
		fmt.Println("Time for joining tables 1 core: %s", elapsed)

		start = time.Now()
		Join(courseHashTable, studentHashTable, 2)
		elapsed = time.Since(start)
		mean2 += elapsed
		fmt.Println("Time for joining tables 2 cores: %s", elapsed)

		start = time.Now()
		Join(courseHashTable, studentHashTable, 4)
		elapsed = time.Since(start)
		mean4 += elapsed
		fmt.Println("Time for joining tables 4 cores: %s", elapsed)

		start = time.Now()
		Join(courseHashTable, studentHashTable, 8)
		elapsed = time.Since(start)
		mean8 += elapsed
		fmt.Println("Time for joining tables 8 cores: %s", elapsed)

		start = time.Now()
		Join(courseHashTable, studentHashTable, 16)
		elapsed = time.Since(start)
		mean16 += elapsed
		fmt.Println("Time for joining tables 16 cores: %s", elapsed)
		
		start = time.Now()
		Join(courseHashTable, studentHashTable, 32)
		elapsed = time.Since(start)
		mean32 += elapsed
		fmt.Println("Time for joining tables 32 cores: %s", elapsed)
	}
	
	fmt.Println("Mean 1: ", mean1/rounds)
	fmt.Println("Mean 2: ", mean2/rounds)
	fmt.Println("Mean 4: ", mean4/rounds)
	fmt.Println("Mean 8: ", mean8/rounds)
	fmt.Println("Mean 16: ", mean16/rounds)
	fmt.Println("Mean 32: ", mean32/rounds)

}

//This funktion will be executed at program start
func init() {
	tpl = template.Must(template.ParseGlob(path.Join(tplPath, "*.html")))

	fmt.Println(runtime.NumCPU())

	//createTables()
	loadTables()
	courseHashTable, studentHashTable := partitioning(dataTables.CourseT, dataTables.StudentT)
	testHashjoin(courseHashTable, studentHashTable)
}

//################################################################################################################

//############################################### MAIN ###########################################################
func main() {
	http.HandleFunc("/", index)             //welcome page
	http.HandleFunc("/tables.html", tables) //generate and show tables and if hash join was executet also the joined table
	http.HandleFunc("/join.html", join)     //does the magic and schows some statistics

	//Create a Filesystem so that the paths in the html files work
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(path.Join(tplPath, "/assets/")))))

	//specifying that it should listen on port 8080. Blocks until the program terminates.
	http.ListenAndServe(":8080", nil)
}

//################################################################################################################

//####################################### HANDLERS ###############################################################
//An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.
//An http.Request is a data structure that represents the client HTTP request.
func index(w http.ResponseWriter, r *http.Request) {
	//Tell the Browser which Content he gets servered
	w.Header().Set("Content-Type", "text/html")

	//Render the template
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func tables(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(w, "tables.html", &dataTables)
}

func join(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tpl.ExecuteTemplate(w, "join.html", nil)
}

//###################################################################################################################
