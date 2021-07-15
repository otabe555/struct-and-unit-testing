package main

import (
	"fmt"
	"time"
)

func main() {
	p := CreateContactManager()
	p.Create(&Contact{ID: 123, Name: "Joe", Phone: "+999123452211", Firstname: "John", Lastname: "Doe", Email: "johndoe@mail.com", CreatedAt: time.Now().Format(time.UnixDate)})
	p.Create(&Contact{ID: 143, Name: "Jee", Phone: "+999123452211", Firstname: "Ian", Lastname: "Doe", Email: "jeendoe@mail.com", CreatedAt: time.Now().Format(time.UnixDate)})
	p.Create(&Contact{ID: 1123, Name: "Jof", Phone: "+999123452211", Firstname: "Pak", Lastname: "Doe", Email: "pakdoe@mail.com", CreatedAt: time.Now().Format(time.UnixDate)})
	fmt.Println("the deleted element: =  ", *p.Delete(123))
	fmt.Println("id = 143", *p.Get(143))
	fmt.Println(*p.GetAll())
}
