package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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
		cmd := exec.Command("clear") //MacOs example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

type customer struct {
	name            string
	email           string
	phone           string
	shippingAddress shippingAddress
	defaultAddress  defaultAddress
}

var customers = make(map[string]customer, 5)

type shippingAddress struct {
	streetName string
	city       string
	postalCode string
}

var shippingAddresses = make(map[string]shippingAddress, 5)

type defaultAddress struct {
	streetName string
	city       string
	postalCode string
}

var defaultAddresses = make(map[string]defaultAddress, 5)

type item struct {
	id    string
	name  string
	stock int
	price int
}

var items = make(map[string]item, 5)

type cart struct {
	customer   customer
	orderedQty int
	item       item
}

var carts = make(map[string]cart, 5)

func main() {
	login()
}

func login() {
	menuScanner := bufio.NewScanner(os.Stdin)
	var pilihanLogin string
menuLogin:
	for {
		clearScreen()
		fmt.Println("SELAMAT DATANG DI GOMART")
		fmt.Println("LOGIN SEBAGAI : ")
		fmt.Println("1. ADMINISTRATOR")
		fmt.Println("2. USER")
		fmt.Println("0. KELUAR")
		fmt.Printf("> ")
		menuScanner.Scan()
		pilihanLogin = menuScanner.Text()
		switch pilihanLogin {
		case "1":
			menuAdmin(menuScanner)
		case "2":
			menuUser(menuScanner)
		case "0":
			fmt.Println("TERIMA KASIH TELAH MENGUNJUNGI KAMI")
			break menuLogin
		}
	}
}

func menuAdmin(menuScanner *bufio.Scanner) {
	adminScanner := bufio.NewScanner(os.Stdin)
	var pilihanAdmin string
menuAdmin:
	for {
		clearScreen()
		fmt.Println("MENU ADMIN")
		fmt.Println("1. LIHAT DAFTAR PRODUK")
		fmt.Println("2. TAMBAH PRODUK")
		fmt.Println("3. EDIT PRODUK")
		fmt.Println("4. HAPUS PRODUK")
		fmt.Println("0. KELUAR")
		fmt.Printf("> ")
		menuScanner.Scan()
		pilihanAdmin = menuScanner.Text()
		switch pilihanAdmin {
		case "1":
			daftarProduk()
		case "2":
			tambahProduk(adminScanner)
		case "3":
			editProduk(adminScanner)
		case "4":
			hapusProduk(adminScanner)
		case "0":
			fmt.Println("0")
			break menuAdmin
		}
	}
}

func daftarProduk() {
	clearScreen()
	fmt.Println("DAFTAR PRODUK")
	fmt.Println()
	for _, v := range items {
		fmt.Println("KODE PRODUK : " + v.id)
		fmt.Println("NAMA PRODUK : " + v.name)
		fmt.Println("JUMLAH BARANG : " + strconv.Itoa(v.stock))
		fmt.Println("HARGA SATUAN : " + strconv.Itoa(v.price))
		fmt.Println()
	}

	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func tambahProduk(adminScanner *bufio.Scanner) {
	clearScreen()
	var item item
	fmt.Printf("TAMBAH PRODUK\n\n")
	fmt.Printf("KODE PRODUK : ")
	adminScanner.Scan()
	item.id = adminScanner.Text()
	fmt.Printf("NAMA PRODUK : ")
	adminScanner.Scan()
	item.name = adminScanner.Text()
	fmt.Printf("JUMLAH BARANG : ")
	adminScanner.Scan()
	item.stock, _ = strconv.Atoi(adminScanner.Text())
	fmt.Printf("HARGA SATUAN : ")
	adminScanner.Scan()
	item.price, _ = strconv.Atoi(adminScanner.Text())

	items[item.id] = item

	fmt.Println()
	fmt.Println("DATA BERHASIL DISIMPAN")
	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func editProduk(adminScanner *bufio.Scanner) {
	clearScreen()
	fmt.Printf("EDIT PRODUK\n\n")
	fmt.Printf("MASUKKAN KODE PRODUK : ")
	adminScanner.Scan()
	kodeProduk := adminScanner.Text()
	if produk, found := items[kodeProduk]; found {
		var item item

		fmt.Println("NAMA PRODUK : " + produk.name)
		fmt.Println("JUMLAH BARANG : " + strconv.Itoa(produk.stock))
		fmt.Println("HARGA SATUAN : " + strconv.Itoa(produk.price))
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println()
		fmt.Println("INPUT DATA PRODUK")
		fmt.Printf("NAMA PRODUK : ")
		adminScanner.Scan()
		item.name = adminScanner.Text()
		fmt.Printf("JUMLAH BARANG : ")
		adminScanner.Scan()
		item.stock, _ = strconv.Atoi(adminScanner.Text())
		fmt.Printf("HARGA SATUAN : ")
		adminScanner.Scan()
		item.price, _ = strconv.Atoi(adminScanner.Text())
		item.id = produk.id

		items[item.id] = item

		fmt.Println()
		fmt.Println("PRODUK BERHASIL DIEDIT")
	} else {
		fmt.Println()
		fmt.Println("KODE PRODUK TIDAK DITEMUKAN")
	}

	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func hapusProduk(adminScanner *bufio.Scanner) {
	clearScreen()
	fmt.Printf("HAPUS PRODUK\n\n")
	fmt.Printf("MASUKKAN KODE PRODUK : ")
	adminScanner.Scan()
	kodeProduk := adminScanner.Text()
	if produk, found := items[kodeProduk]; found {
		fmt.Println("NAMA PRODUK : " + produk.name)
		fmt.Println("JUMLAH BARANG : " + strconv.Itoa(produk.stock))
		fmt.Println("HARGA SATUAN : " + strconv.Itoa(produk.price))
		fmt.Println()
		fmt.Println("ANDA YAKIN UNTUK MENGHAPUS PRODUK DENGAN KODE " + produk.id + " ? Y / N")
		adminScanner.Scan()
		if adminScanner.Text() == "Y" || adminScanner.Text() == "y" {
			delete(items, produk.id)

			fmt.Println()
			fmt.Println("PRODUK BERHASIL DIHAPUS")
		}
	} else {
		fmt.Println()
		fmt.Println("KODE PRODUK TIDAK DITEMUKAN")
	}

	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func menuUser(menuScanner *bufio.Scanner) {
	userScanner := bufio.NewScanner(os.Stdin)
	var pilihanAdmin string
menuAdmin:
	for {
		clearScreen()
		fmt.Println("MENU USER")
		fmt.Println("1. LIHAT DATA USER")
		fmt.Println("2. TAMBAH DATA USER")
		fmt.Println("3. EDIT DATA USER")
		fmt.Println("4. HAPUS DATA USER")
		fmt.Println("5. TAMBAH KE KERANJANG")
		fmt.Println("6. LIHAT DATA KERANJANG")
		fmt.Println("7. HAPUS DATA KERANJANG")
		fmt.Println("8. PEMBAYARAN")
		fmt.Println("0. KELUAR")
		fmt.Printf("> ")
		menuScanner.Scan()
		pilihanAdmin = menuScanner.Text()
		switch pilihanAdmin {
		case "1":
			lihatDataUser(userScanner)
		case "2":
			tambahDataUser(userScanner)
		case "3":
			editDataUser(userScanner)
		case "4":
			hapusDataUser(userScanner)
		case "5":
			tambahKeranjang(userScanner)
		case "6":
			lihatKeranjang(userScanner)
		case "7":
			hapusKeranjang(userScanner)
		case "8":
			pembayaran(userScanner)
		case "0":
			fmt.Println("0")
			break menuAdmin
		}
	}
}

func lihatDataUser(userScanner *bufio.Scanner) {
	clearScreen()
	fmt.Printf("LIHAT DATA USER\n\n")
	fmt.Printf("MASUKKAN NAMA USER : ")
	userScanner.Scan()
	namaUser := userScanner.Text()
	if user, found := customers[namaUser]; found {
		fmt.Println("EMAIL : " + user.email)
		fmt.Println("PHONE : " + user.phone)
		fmt.Println("ALAMAT ASAL : " + user.defaultAddress.streetName + " " + user.defaultAddress.city + " " + user.defaultAddress.postalCode)
		fmt.Println("ALAMAT PENGIRIMAN : " + user.shippingAddress.streetName + " " + user.shippingAddress.city + " " + user.shippingAddress.postalCode)
	} else {
		fmt.Println()
		fmt.Println("USER TIDAK DITEMUKAN")
	}

	fmt.Println()
	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func tambahDataUser(userScanner *bufio.Scanner) {
	clearScreen()
	var customer customer
	fmt.Printf("TAMBAH USER\n\n")
	fmt.Printf("NAMA USER : ")
	userScanner.Scan()
	customer.name = userScanner.Text()
	fmt.Printf("EMAIL : ")
	userScanner.Scan()
	customer.email = userScanner.Text()
	fmt.Printf("PHONE : ")
	userScanner.Scan()
	customer.phone = userScanner.Text()
	fmt.Printf("ALAMAT ASAL (NAMA JALAN;KOTA;KODEPOS) : ")
	userScanner.Scan()
	alamatAsal := strings.Split(userScanner.Text(), ";")
	customer.defaultAddress.streetName = alamatAsal[0]
	customer.defaultAddress.city = alamatAsal[1]
	customer.defaultAddress.postalCode = alamatAsal[2]
	fmt.Printf("ALAMAT PENGIRIMAN (NAMA JALAN;KOTA;KODEPOS) : ")
	userScanner.Scan()
	alamatPengiriman := strings.Split(userScanner.Text(), ";")
	customer.shippingAddress.streetName = alamatPengiriman[0]
	customer.shippingAddress.city = alamatPengiriman[1]
	customer.shippingAddress.postalCode = alamatPengiriman[2]

	customers[customer.name] = customer

	fmt.Println()
	fmt.Println("USER BERHASIL DISIMPAN")
	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func editDataUser(userScanner *bufio.Scanner) {
	clearScreen()
	fmt.Printf("EDIT USER\n\n")
	fmt.Printf("MASUKKAN NAMA USER : ")
	userScanner.Scan()
	namaUser := userScanner.Text()
	if user, found := customers[namaUser]; found {
		var customer customer

		fmt.Println("EMAIL : " + user.email)
		fmt.Println("PHONE : " + user.phone)
		fmt.Println("ALAMAT ASAL : " + user.defaultAddress.streetName + " " + user.defaultAddress.city + " " + user.defaultAddress.postalCode)
		fmt.Println("ALAMAT PENGIRIMAN : " + user.shippingAddress.streetName + " " + user.shippingAddress.city + " " + user.shippingAddress.postalCode)
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println()
		fmt.Println("INPUT DATA USER")
		fmt.Printf("EMAIL : ")
		userScanner.Scan()
		customer.email = userScanner.Text()
		fmt.Printf("PHONE : ")
		userScanner.Scan()
		customer.phone = userScanner.Text()
		fmt.Printf("ALAMAT ASAL (NAMA JALAN;KOTA;KODEPOS) : ")
		userScanner.Scan()
		alamatAsal := strings.Split(userScanner.Text(), ";")
		customer.defaultAddress.streetName = alamatAsal[0]
		customer.defaultAddress.city = alamatAsal[1]
		customer.defaultAddress.postalCode = alamatAsal[2]
		fmt.Printf("ALAMAT PENGIRIMAN (NAMA JALAN;KOTA;KODEPOS) : ")
		userScanner.Scan()
		alamatPengiriman := strings.Split(userScanner.Text(), ";")
		customer.shippingAddress.streetName = alamatPengiriman[0]
		customer.shippingAddress.city = alamatPengiriman[1]
		customer.shippingAddress.postalCode = alamatPengiriman[2]
		customer.name = user.name

		customers[customer.name] = customer

		fmt.Println()
		fmt.Println("USER BERHASIL DIEDIT")
	} else {
		fmt.Println()
		fmt.Println("USER TIDAK DITEMUKAN")
	}

	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func hapusDataUser(userScanner *bufio.Scanner) {
	clearScreen()
	fmt.Printf("MASUKKAN NAMA USER\n\n")
	userScanner.Scan()
	namaUser := userScanner.Text()
	if user, found := customers[namaUser]; found {
		fmt.Println("EMAIL : " + user.email)
		fmt.Println("PHONE : " + user.phone)
		fmt.Println("ALAMAT ASAL : " + user.defaultAddress.streetName + " " + user.defaultAddress.city + " " + user.defaultAddress.postalCode)
		fmt.Println("ALAMAT PENGIRIMAN : " + user.shippingAddress.streetName + " " + user.shippingAddress.city + " " + user.shippingAddress.postalCode)
		fmt.Println()
		fmt.Println("ANDA YAKIN UNTUK MENGHAPUS USER DENGAN NAMA " + user.name + " ? Y / N")
		userScanner.Scan()
		if userScanner.Text() == "Y" || userScanner.Text() == "y" {
			delete(customers, user.name)

			fmt.Println()
			fmt.Println("USER BERHASIL DIHAPUS")
		}
	} else {
		fmt.Println()
		fmt.Println("USER TIDAK DITEMUKAN")
	}

	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func tambahKeranjang(userScanner *bufio.Scanner) {
	clearScreen()
	fmt.Printf("TAMBAH KERANJANG\n\n")
	fmt.Printf("MASUKKAN NAMA USER : ")
	userScanner.Scan()
	namaUser := userScanner.Text()
	if user, found := customers[namaUser]; found {
		fmt.Println("EMAIL : " + user.email)
		fmt.Println("PHONE : " + user.phone)
		fmt.Println("ALAMAT ASAL : " + user.defaultAddress.streetName + " " + user.defaultAddress.city + " " + user.defaultAddress.postalCode)
		fmt.Println("ALAMAT PENGIRIMAN : " + user.shippingAddress.streetName + " " + user.shippingAddress.city + " " + user.shippingAddress.postalCode)
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println()
		fmt.Printf("MASUKKAN KODE PRODUK : ")
		userScanner.Scan()
		kodeProduk := userScanner.Text()
		if produk, found := items[kodeProduk]; found {
			var cart cart

			jumlahBarang := produk.stock

			fmt.Println("NAMA PRODUK : " + produk.name)
			fmt.Println("JUMLAH BARANG : " + strconv.Itoa(produk.stock))
			fmt.Println("HARGA SATUAN : " + strconv.Itoa(produk.price))
			fmt.Printf("MASUKKAN JUMLAH YANG INGIN DIBELI : ")
			userScanner.Scan()
			orderedQty, _ := strconv.Atoi(userScanner.Text())
			if orderedQty <= jumlahBarang {
				cart.customer = user
				cart.orderedQty = orderedQty
				cart.item = produk
				carts[user.name] = cart

				produk.stock = produk.stock - orderedQty
				items[produk.id] = produk

				fmt.Println()
				fmt.Println("BERHASIL MENAMBAH KERANJANG")
			} else {
				fmt.Println()
				fmt.Println("JUMLAH PRODUK YANG INGIN DIBELI LEBIH DARI STOK YANG TERSEDIA")
			}
		} else {
			fmt.Println()
			fmt.Println("PRODUK TIDAK DITEMUKAN")
		}
	} else {
		fmt.Println()
		fmt.Println("USER TIDAK DITEMUKAN")
	}

	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func lihatKeranjang(userScanner *bufio.Scanner) {
	clearScreen()
	fmt.Printf("LIHAT DATA KERANJANG\n\n")
	fmt.Printf("MASUKKAN NAMA USER : ")
	userScanner.Scan()
	namaUser := userScanner.Text()
	if keranjang, found := carts[namaUser]; found {
		var hargaSatuan int
		var totalHarga int

		hargaSatuan = keranjang.item.price
		totalHarga = keranjang.orderedQty * hargaSatuan

		fmt.Println("KODE PRODUK : " + keranjang.item.id)
		fmt.Println("NAMA PRODUK : " + keranjang.item.name)
		fmt.Println("JUMLAH : " + strconv.Itoa(keranjang.orderedQty))
		fmt.Println("HARGA SATUAN : " + strconv.Itoa(hargaSatuan))
		fmt.Println("TOTAL HARGA : " + strconv.Itoa(totalHarga))
		fmt.Println("EMAIL : " + keranjang.customer.email)
		fmt.Println("PHONE : " + keranjang.customer.phone)
		fmt.Println("ALAMAT ASAL : " + keranjang.customer.defaultAddress.streetName + " " + keranjang.customer.defaultAddress.city + " " + keranjang.customer.defaultAddress.postalCode)
		fmt.Println("ALAMAT PENGIRIMAN : " + keranjang.customer.shippingAddress.streetName + " " + keranjang.customer.shippingAddress.city + " " + keranjang.customer.shippingAddress.postalCode)
	} else {
		fmt.Println()
		fmt.Println("USER TIDAK DITEMUKAN")
	}

	fmt.Println()
	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func hapusKeranjang(userScanner *bufio.Scanner) {
	clearScreen()
	fmt.Printf("HAPUS DATA KERANJANG\n\n")
	fmt.Printf("MASUKKAN NAMA USER : ")
	userScanner.Scan()
	namaUser := userScanner.Text()
	if keranjang, found := carts[namaUser]; found {
		var hargaSatuan int
		var totalHarga int

		hargaSatuan = keranjang.item.price
		totalHarga = keranjang.orderedQty * hargaSatuan

		fmt.Println("KODE PRODUK : " + keranjang.item.id)
		fmt.Println("NAMA PRODUK : " + keranjang.item.name)
		fmt.Println("JUMLAH : " + strconv.Itoa(keranjang.orderedQty))
		fmt.Println("HARGA SATUAN : " + strconv.Itoa(hargaSatuan))
		fmt.Println("TOTAL HARGA : " + strconv.Itoa(totalHarga))
		fmt.Println("EMAIL : " + keranjang.customer.email)
		fmt.Println("PHONE : " + keranjang.customer.phone)
		fmt.Println("ALAMAT ASAL : " + keranjang.customer.defaultAddress.streetName + " " + keranjang.customer.defaultAddress.city + " " + keranjang.customer.defaultAddress.postalCode)
		fmt.Println("ALAMAT PENGIRIMAN : " + keranjang.customer.shippingAddress.streetName + " " + keranjang.customer.shippingAddress.city + " " + keranjang.customer.shippingAddress.postalCode)

		fmt.Println()
		fmt.Println("ANDA YAKIN UNTUK MENGHAPUS DATA KERANJANG? Y / N")
		userScanner.Scan()
		if userScanner.Text() == "Y" || userScanner.Text() == "y" {
			var item item

			item = items[keranjang.item.id]
			item.stock = item.stock + keranjang.orderedQty
			items[keranjang.item.id] = item

			delete(carts, namaUser)

			fmt.Println()
			fmt.Println("DATA KERANJANG BERHASIL DIHAPUS")
		}
	} else {
		fmt.Println()
		fmt.Println("DATA KERANJANG TIDAK DITEMUKAN")
	}

	fmt.Println()
	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func pembayaran(userScanner *bufio.Scanner) {
	clearScreen()
	clearScreen()
	fmt.Printf("PEMBAYARAN\n\n")
	fmt.Printf("MASUKKAN NAMA USER : ")
	userScanner.Scan()
	namaUser := userScanner.Text()
	if keranjang, found := carts[namaUser]; found {
		var hargaSatuan int
		var totalHarga int

		hargaSatuan = keranjang.item.price
		totalHarga = keranjang.orderedQty * hargaSatuan

		fmt.Println("KODE PRODUK : " + keranjang.item.id)
		fmt.Println("NAMA PRODUK : " + keranjang.item.name)
		fmt.Println("JUMLAH : " + strconv.Itoa(keranjang.orderedQty))
		fmt.Println("HARGA SATUAN : " + strconv.Itoa(hargaSatuan))
		fmt.Println("TOTAL HARGA : " + strconv.Itoa(totalHarga))
		fmt.Println("EMAIL : " + keranjang.customer.email)
		fmt.Println("PHONE : " + keranjang.customer.phone)
		fmt.Println("ALAMAT ASAL : " + keranjang.customer.defaultAddress.streetName + " " + keranjang.customer.defaultAddress.city + " " + keranjang.customer.defaultAddress.postalCode)
		fmt.Println("ALAMAT PENGIRIMAN : " + keranjang.customer.shippingAddress.streetName + " " + keranjang.customer.shippingAddress.city + " " + keranjang.customer.shippingAddress.postalCode)

		fmt.Println()
		fmt.Println("ANDA YAKIN UNTUK MELAKUKAN PEMBAYARAN ? Y / N")
		userScanner.Scan()
		if userScanner.Text() == "Y" || userScanner.Text() == "y" {
			fmt.Println()
			fmt.Println("PEMBAYARAN BERHASIL DILAKUKAN")
		}
	} else {
		fmt.Println()
		fmt.Println("DATA KERANJANG TIDAK DITEMUKAN")
	}

	fmt.Println()
	fmt.Println("TEKAN SEMBARANG TOMBOL UNTUK KEMBALI")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
