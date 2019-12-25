package main

import "fmt"

type user struct {
	name  string
	email string
}

// notify implements a method with a VALUE receiver
func (u user) notify() {
	fmt.Printf("Send email to %s<%s>\n", u.name, u.email)
}

// changeEmail implements a method with a POINTER receiver
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	/*
	   VALUES of type user can be used to call methods
	   declared with both value and pointer receivers
	*/
	aakash := user{"aakash", "aakash@email.com"}
	aakash.notify()
	fmt.Printf("%+v\n", aakash.email)
	fmt.Println("Email after invoking change email")
	aakash.changeEmail("aakash@123mail.com")
	fmt.Printf("%+v\n", aakash.email)

	/*
	   POINTERS of type user can also be used to call methods
	   declated with both value and pointer receivers
	*/
	mike := &user{"mike", "mike@email.com"}
	mike.notify()
	/*
	   In the call above, notify() uses a VALUE receiver,
	   yet somehow we seem to be able to invoke the method
	   using a POINTER (???)
	   The go runtime internally creates a VALUE from the object
	   being pointed to , and passes that a receiver to notify()
	*/
	fmt.Printf("%+v\n", mike.email)
	fmt.Println("Email after invoking change email")
	mike.changeEmail("mike@123mail.com")
	fmt.Printf("%+v\n", mike.email)
}
