package main

import "fmt"

type contactInfo struct{
	email string
	zip int
}
type person struct{
	firstName string
	lastName string
	contactInfo
}

func main(){
	jim := person{
		firstName : "Jim",
		lastName : "Neutron",
		contactInfo : contactInfo{
			email : "jim@yahoo.com",
			zip : 25000,
		},

	}

	jim.updateName("Jimmy")
	jim.print()
	

}

func (p person) print(){
	fmt.Printf("%+v" ,p)
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}
