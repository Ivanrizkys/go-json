package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logEncode(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logEncode("Hai Sayang")              // "Hai Sayang"
	logEncode(1)                         // 1
	logEncode(false)                     // false
	logEncode([]string{"Hai", "Sayang"}) // ["Hai","Sayang"]
}

type family struct {
	Anak           string
	Istri          string
	JumlahKeluarga int
}

type Person struct {
	// ** agar hasil key dari json yang sudah di encode tidak huruf besar
	Nama string `json:"nama"`
	// ** hasil key dari jsson encode huruf besar
	Status string
	Age    int
	Maried bool
	Family []family
}

func TestStructObject(t *testing.T) {
	person := Person{
		Nama:   "Ivan Rizky Saputra",
		Status: "Kuliah",
		Age:    19,
		Maried: false,
	}
	bytes, _ := json.Marshal(person)
	fmt.Println(string(bytes)) // {"nama":"Ivan Rizky Saputra","Status":"Kuliah","Age":19,"Maried":false}
}

func TestDecode(t *testing.T) {
	// * data json
	data := `{"nama":"Ivan Rizky Saputra","Status":"Kuliah","Age":19,"Maried":false}`
	// * embed json to byte
	dataBytes := []byte(data)
	person := Person{}

	err := json.Unmarshal(dataBytes, &person)
	if err != nil {
		panic(err)
	}

	fmt.Println(person)
	fmt.Println(person.Nama)
}

func TestJsonFromSlice(t *testing.T) {
	orang1 := Person{
		Nama:   "Ivan Rizky Saputra",
		Status: "Kepala Keluarga",
		Age:    20,
		Maried: true,
		Family: []family{
			{
				Anak:           "Sabila",
				Istri:          "Fera Anissa",
				JumlahKeluarga: 3,
			},
			{
				Anak:           "Rifki Adi",
				Istri:          "Yolanda Putri",
				JumlahKeluarga: 3,
			},
		},
	}
	bytes, _ := json.Marshal(orang1)
	fmt.Println(string(bytes))
}

func TestJsonToSlice(t *testing.T) {
	arrObj := `{"nama":"Ivan Rizky Saputra","Status":"Kepala Keluarga","Age":20,"Maried":true,"Family":[{"Anak":"Sabila","Istri":"Fera Anissa","JumlahKeluarga":3},{"Anak":"Rifki Adi","Istri":"Yolanda Putri","JumlahKeluarga":3}]}`
	dataBytes := []byte(arrObj)
	data := Person{}
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data.Nama)
	fmt.Println(data.Family)
	fmt.Println(data.Family[0])
}
