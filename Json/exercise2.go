package main
import (
	"encoding/json"
	"fmt"
)
type Company struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipcode"`
}
func main() {
	company := Company{
		Name: "Tesla",
		Address: Address{
			Street:  "1 Tesla Rd",
			City:    "Austin",
			State:   "TX",
			ZipCode: "78725",
		},
	}
	jsonData, err := json.Marshal(company)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
	fmt.Println("_____-------_____")
	var newCompany Company
	err = json.Unmarshal([]byte(jsonData), &newCompany)
	if err != nil {
		fmt.Println("Error Unma", err)
		return
	}
	fmt.Printf("%+v\n", newCompany)

}
