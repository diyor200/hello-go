package main

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func encoding_func() {
	message := "hello,go (*w3hu%#"
	demoBase64(message)
	demoHex(message)
	demoJSON()
	demoXML()
	demoCSV()
}

func demoCSV() {
	print("---Demo encoding csv ---\n")
	type Employee struct {
		Name    string
		Email   string
		Country string
	}

	// read csv file
	print("File o'qilmoqda...\n")
	file, err := os.Open("employee.csv")
	if err != nil {
		fmt.Println("Fileni ochishda xatolik yuz berdi!", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	var emp Employee
	var employees []Employee

	for _, item := range csvData {
		emp.Name = item[0]
		emp.Email = item[1]
		emp.Country = item[2]

		employees = append(employees, emp)
		fmt.Printf("name: %s, email: %s, country: %s\n", item[0], item[1], item[2])
	}

	print("Showing all employee strcts ...\n")
	fmt.Println(employees)

	// write data to csv file
	print("writing data to csv file\n")

	csvFile, err := os.Create("demo.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = ','

	for _, itemEmp := range employees {
		records := []string{
			itemEmp.Name,
			itemEmp.Email,
			itemEmp.Country,
		}

		err := writer.Write(records)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
	print("Done!")
}

func demoBase64(message string) {

	fmt.Println("---- Demo encoding base64 -----")

	fmt.Println("Plain text: ")
	fmt.Println(message)

	encoding := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println("base64 msg: ")
	fmt.Println(encoding)

	decoding, _ := base64.StdEncoding.DecodeString(encoding)
	fmt.Println("decoding base64 msg: ")
	fmt.Println(string(decoding))
}

func demoHex(message string) {
	fmt.Println("----Demo encoding Hex----")

	fmt.Println("plain text: ")
	fmt.Println(message)

	encoding := hex.EncodeToString([]byte(message))
	fmt.Println("Hex message: ", encoding)

	decoding, _ := hex.DecodeString(encoding)
	fmt.Println("decoding Hex message: ", string(decoding))
}

func demoJSON() {
	type Employee struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	print("---demo json encoding---\n")

	emp := &Employee{
		Id:    "14",
		Name:  "Diyorbek",
		Email: "test@gmail.com",
	}
	fmt.Println(emp)

	print(">>> struct to json ...\n")
	b, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// json to struct
	var newEmp Employee
	fmt.Println(&newEmp)
	str := `{"Id":"154","Name":"Diyorbek","Email":"email@gmail.com"}`
	print("<<< jsont to struct ...\n")
	json.Unmarshal([]byte(str), &newEmp)
	fmt.Printf("Id: %s\n", newEmp.Id)
	fmt.Printf("Name: %s\n", newEmp.Name)
	fmt.Printf("Email: %s\n", newEmp.Email)
}

func demoXML() {
	print("----Demo XML encoding----\n")

	type EmployeeCountry struct {
		CountryCode string `xml:"code,attr"` // xml attribute: code
		CountryName string `xml:",chardata"` // XML value
	}

	type Employee struct {
		XMLName xml.Name        `xml:"employee"`
		Id      string          `xml:"id"`
		Name    string          `xml:"name"`
		Email   string          `xml:"email"`
		Country EmployeeCountry `xml:"country"`
	}

	// struct to xml
	print(">>> struct to XML ...\n")
	emp := &Employee{
		Id:    "15",
		Name:  "Diyorbek",
		Email: "test2gmail.com",
		Country: EmployeeCountry{
			CountryCode: "uz",
			CountryName: "Uzbekistan",
		},
	}

	b, err := xml.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}
	print(string(b))

	// xml string to struct
	print("\n<<< xml to string ...\n")
	var newEmp Employee
	xmlStr := `
	<employee>
		<id>150</id>
		<name>Diyorbek</name>
		<email>test@gmail.com</email>
		<country code="uz">Uzbekistan
		</country>
	</employee>
	`
	xml.Unmarshal([]byte(xmlStr), &newEmp)
	fmt.Printf("Id: %s\n", newEmp.Id)
	fmt.Printf("Name: %s\n", newEmp.Name)
	fmt.Printf("Email: %s\n", newEmp.Email)
	fmt.Printf("CountryCode: %s\n", newEmp.Country.CountryCode)
	fmt.Printf("CountryName: %s\n", newEmp.Country.CountryName)
}
