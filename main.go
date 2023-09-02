package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Struct participants json
type Participants struct {
	Participants []student `json:"participants"`
}

// struct student
type student struct {
	Id        string `json:"id"`
	Code      string `json:"student_code"`
	Nama      string `json:"student_name"`
	Alamat    string `json:"student_address"`
	Pekerjaan string `json:"student_occupation"`
	Alasan    string `json:"joining_reason"`
}

// get data student by Kode Peserta
func getStudent(code string, arrayOfParticipants []student) (student, error) {
	code = strings.ToLower(code)
	var dataStudent student
	for _, readStudent := range arrayOfParticipants {
		if strings.ToLower(readStudent.Code) == code {
			dataStudent = readStudent
			return dataStudent, nil
		}
	}
	return dataStudent, errors.New("Data Not Found")

}

// error function
func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// function main untuk membaca file json yang akan menampilkan data student
func main() {
	jsonByte, err := os.ReadFile("participants.json")
	errorCheck(err)

	var data Participants
	errorCheck(json.Unmarshal(jsonByte, &data))

	student, err := getStudent(os.Args[1], data.Participants)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("ID\t\t: %s\n", student.Id)
	fmt.Printf("Nama\t\t: %s\n", student.Nama)
	fmt.Printf("Alamat\t\t: %s\n", student.Alamat)
	fmt.Printf("Pekerjaan\t: %s\n", student.Pekerjaan)
	fmt.Printf("Alasan\t\t: %s", student.Alasan)
}