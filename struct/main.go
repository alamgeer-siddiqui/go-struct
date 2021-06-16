package main

import "fmt"

type contactInfo struct{
	email string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo//struct data type 
}

func main() {
	// alam := person{firstName: "Alamgeer", lastName: "siddiqui"}
	// fmt.Println(alam)

	// var alam person
	// alam.firstName="ALAMGEER"
	// alam.lastName="SIDDIQUI"
	// fmt.Println(alam)
	// fmt.Printf("%+v",alam)

	alam:=person{
		firstName: "alam",
		lastName: "siddiqui",
		contactInfo: contactInfo{
			email: "alamgeer.siddiqui@qsstechnosoft.com",
			pinCode: 250002,
		},
	
	}
	// alamPointer:=&alam
	alam.updateName("alamgeer")
	alam.print()


}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName=newFirstName
}

func (p person) print(){
	fmt.Printf("%+v",p)
}
