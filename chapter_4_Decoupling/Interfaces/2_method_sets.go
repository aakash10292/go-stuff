// Sample program to show how to understand method sets.
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

func main() {

	// Create a value of type User and send a notification.
	u := user{"Bill", "bill@email.com"}

	// Values of type user do not implement the interface because pointer
	// receivers don't belong to the method set of a value.

	// Uncomment the below line of code
	// sendNotification(u)
	// You should see an error as follows:
	// ./example1.go:36: cannot use u (type user) as type notifier in argument to sendNotification:
	//   user does not implement notifier (notify method has pointer receiver)

	sendNotification(&u)
	/*
	   Why does passing the ADDRESS work but not passing the VALUE?

	   If we are working with a VALUE, then only those methods using VALUE receivers
	   belong to the method set for this value
	   If we are working with a POINTER, then all methods (both VALUE and POINTER receivers)
	   are available in the method set for this pointer

	   WHY THESE RULES??
	   For integrity reasons :
	   * Not every value we work with have addresses. Which means, we cannot be using
	   POINTER semantics for something that has no address a.k.a cannot be shared
	   (3_address_of_value.go for an example). Conversely, if we work with POINTERS,
	   then we already have a pointer (an address). So, we can work with both types of
	   receivers.
	   * An implication of these rules are - If you've chosen POINTER receivers, then you
	   can ONLY SHARE data. If we've chose VALUE receivers, the we can share or pass around
	   copies of data. (recollect the scenario of unmarshalling or decoding where the method
	   is implemented with VALUE receivers but we pass a pointer during method invocation
	   (In this case go under-the-hood, takes the value being pointed to and passes that
	   to the method)
	*/
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
