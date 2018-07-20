package main

import (
	"fmt"
	"net/http"
)

//Student - Student struct
type Student struct {
	studentName string
	rollNumber  int32
	class       string
	section     string
}

//DisplayStudent - ValueReceiver method of the Student Struct
func (myStudent Student) DisplayStudent() string {
	return "The Name of Student is " + myStudent.studentName
}

//EditStudent - PointerReceiver method of the Student Struct
func (myStudent *Student) EditStudent(student Student) {
	student = Student{studentName: "K.VijayaLalitha",
		rollNumber: 22,
		class:      "XII",
		section:    "C"}
	*myStudent = student
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", About)
	http.HandleFunc("/student", DisplayStudent)
	http.ListenAndServe(":8005", nil)
}

//IndexHandler - To display content when accessing Go Web Server
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "First Program based on Go's Web Server!\n")
	fmt.Fprintln(w, "Enjoying Learning Go!")
}

//About - To display About page content
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Displaying About Page")
}

//DisplayStudent - To display a Student Information
func DisplayStudent(w http.ResponseWriter, r *http.Request) {
	myStudent := Student{studentName: "C.R.Kasinath",
		rollNumber: 44,
		class:      "XII",
		section:    "A"}
	fmt.Fprintln(w, "The Name of Student is "+myStudent.studentName)
	fmt.Fprintln(w, "The Class of the Student "+myStudent.studentName+" is "+myStudent.class)
	fmt.Fprintln(w, "The Section of the Student "+myStudent.studentName+" is "+myStudent.section+"\n")
	fmt.Fprintln(w, "Calling to the Student Struct's DisplayStudent() ValueReceiver method:\n"+myStudent.DisplayStudent())
	myStudent.EditStudent(myStudent) //Making a call to PointerReceiver method to change the myStudent object forever
	fmt.Fprintln(w, "Calling to the Student Struct's DisplayStudent() ValueReceiver method to check myStudent Object:\n"+myStudent.DisplayStudent())
}
