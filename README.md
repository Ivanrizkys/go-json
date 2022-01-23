# Encode & Decode

Untuk men encode suatu tipe data menjadi json maka kita membutuhkan function **json.Marshal(interface{})**. Function ini memerima seluruh tipe data dan akan mengembalikan byte dan error. Sedangkan untuk men decode sebuah json maka kita harus membuatya dulu ke dalam bytes lalu memasukkanya ke dalam function **json.Unmarshal(bytes, variabelUntukMenyimpan)**. Unmarshal mengembalikan error.

```go
// encode function
func logEncode(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

// ** encode
func TestEncode(t *testing.T) {
	logEncode("Hai Sayang")              // "Hai Sayang"
	logEncode(1)                         // 1
	logEncode(false)                     // false
	logEncode([]string{"Hai", "Sayang"}) // ["Hai","Sayang"]

// decode
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
```

# Object

Untuk membuat json object, kita bisa menggunakan struct

```go
type Person struct {
	// ** agar hasil key dari json yang sudah di encode tidak huruf besar
	Nama string `json:"nama"`
	// ** hasil key dari jsson encode huruf besar
	Status string
	Age    int
	Maried bool
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
```

