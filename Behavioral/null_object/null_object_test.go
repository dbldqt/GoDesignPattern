package null_object

import (
	"fmt"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	customer1 := NewCustomer("Rob")
	customer2 := NewCustomer("Bob")
	customer3 := NewCustomer("Julie")
	customer4 := NewCustomer("Laura")
	customer5 := NewCustomer("test")

	fmt.Println("customers:")
	fmt.Println(customer1.GetName())
	fmt.Println(customer2.GetName())
	fmt.Println(customer3.GetName())
	fmt.Println(customer4.GetName())
	fmt.Println(customer5.GetName())
}
