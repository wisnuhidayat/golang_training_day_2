package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

type customer struct {
	firstName, lastName string
	phoneNo             string
	email               string
	shippingAddress     shippingAddress
	defaultAddress      defaultAddress
}

type shippingAddress struct {
	address    string
	postalCode int
}

type defaultAddress struct {
	address    string
	postalCode int
}

type item struct {
	id    int
	nam   string
	stock int
	price int
}

type cart struct {
	customer   customer
	item       item
	orderedQty int
}

func main() {
	init()
}

func int() {
	fmt.Println("================================================")
	fmt.Println("||   SELAMAT DATANG DI APLIKASI GO-COMMERCE   ||")
	fmt.Println("================================================")
	fmt.Println("|| MENU                                       ||")
	fmt.Println("|| 1. Customer                                ||")
	fmt.Println("|| 2. Produk                                  ||")
	fmt.Println("|| 3. Belanja                                 ||")
	fmt.Println("|| 4. Exit                                    ||")
	fmt.Println("================================================")
	fmt.Println("> Masukkan Menu : ")
	readerMenu := bufio.NewReader(os.Stdin)
	menu, _ := readerMenu.ReadString('\n')
	if menu == "1" {

	} else if menu == "2" {

	} else if menu == "3" {

	} else if menu == "4" {
		fmt.Println("Terima Kasih Telah Mengunjungi Kami")
	}
}

func fnCustomer() {
	fmt.Print("> First Name : ")
	readerFirstName := bufio.NewReader(os.Stdin)
	firstName, _ := readerFirstName.ReadString('\n')

	fmt.Print("> Last Name : ")
	readerLastName := bufio.NewReader(os.Stdin)
	lastName, _ := readerLastName.ReadString('\n')

	fmt.Print("> Phone Number : ")
	readerPhoneNumber := bufio.NewReader(os.Stdin)
	phoneNumber, _ := readerPhoneNumber.ReadString('\n')

	fmt.Print("> Email : ")
	readerEmail := bufio.NewReader(os.Stdin)
	email, _ := readerEmail.ReadString('\n')
}

func fnProduk() {

}

func fnBelanja() {

}
