package main

import "fmt"

func main() {

	fmt.Println("Struct in golang")
	// no inheritance in golang ; no super ot partner

	Anupam := User{"Anupam", "Anupam@go.dev", true, 18}

	fmt.Println(Anupam)
	fmt.Printf("hitesh details are:%v\n", Anupam)
	fmt.Printf("Name is %v and email is %v,\n", Anupam.Name, Anupam.Email)
	Anupam.GetStatus()
	Anupam.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)

}

func (u User) NewMail() {

	u.Email = "test@go.dev"
	fmt.Println("new email:", u.Email)
	// this wont change the original email .
}
