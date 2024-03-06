package main

var _ User = &user{}

type user struct {
	FIO, Address, Phone string
}

func (u *user) ChangeFIO(newFIO string) {
	u.FIO = newFIO
}

func (u *user) ChangeAddress(newAddress string) {
	u.Address = newAddress
}

type User interface {
	ChangeFIO(newFIO string)
	ChangeAddress(newAddress string)
}

func newUser(fio, address, phone string) User {
	u := user{
		FIO:     fio,
		Address: address,
		Phone:   phone,
	}
	return &u
}

func main() {

}
