package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Participants struct {
	Participants []student `json:"participants"`
}

type student struct {
	Id        string `json:"id"`
	Code      string `json:"student_code"`
	Nama      string `json:"student_name"`
	Alamat    string `json:"student_address"`
	Pekerjaan string `json:"student_occupation"`
	Alasan    string `json:"joining_reason"`
}

func searchStudent(kode string, arrayOfParticipants []student) (student, error) {
	kode = strings.ToLower(kode)
	var theRightStudent student
	for _, chosenStudent := range arrayOfParticipants {
		if strings.ToLower(chosenStudent.Code) == kode {
			theRightStudent = chosenStudent
			return theRightStudent, nil
		}
	}
	return theRightStudent, errors.New("Data Not Found")

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	jsonByte, err := os.ReadFile("participants.json")
	checkError(err)

	var data Participants
	checkError(json.Unmarshal(jsonByte, &data))

	student, err := searchStudent(os.Args[1], data.Participants)
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