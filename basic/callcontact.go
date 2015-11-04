package main

import "contact"
import "fmt"

func main() {
	// contact.Save("Hello","1234","kenshero@mail.com")
	
	c := contact.Contact{
			Tel  	: "123456",
			Email	: "mailmail@.com",
	}

	contact.Add("hello",c)
	fmt.Println(contact.Get("hello"))
}