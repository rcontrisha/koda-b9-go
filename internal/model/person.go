package model

import "fmt"

type Person struct {
	Nama    string
	Address string
	Phone   string
}

func NewPerson (nama string, address string, phone string) *Person {
	return &Person{
		Nama: nama,
		Address: address,
		Phone: phone,
	}
}

func (p *Person) Print() (nama string, address string, phone string) {
	nama = p.Nama
	address = p.Address
	phone = p.Phone
	return nama, address, phone
}

func (p *Person) Greet() {
	fmt.Printf("Hello, %s!\n", p.Nama)
}

func (p *Person) ChangeName(newName string) {
	p.Nama = newName
}