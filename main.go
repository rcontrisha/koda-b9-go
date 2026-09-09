package main

import "rcontrisha/koda-b9-go/internal/service"

// "bufio"
// "fmt"
// "os"
// "strconv"
// "strings"

// "rcontrisha/koda-b9-go/internal/model"
// "rcontrisha/koda-b9-go/internal/service"
// "rcontrisha/koda-b9-go/internal/checkout"

func main() {
	// fmt.Println("Hello World")
	// perimeter, area := countRectangle(10, 5)
	// fmt.Printf("Perimeter: %d\n", perimeter)
	// fmt.Printf("Area: %d", area)
	// fmt.Println(countRectangle(10,5))

	// err := generateWindow(5)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// injectToSlice()

	// myData := model.Biodata {
	// 	Nama: "Ridho Contrisha",
	// 	Foto: "",
	// 	Email: "rcontrisha@gmail.com",
	// 	Umur: 23,
	// 	Telepon: "081315468293",
	// 	Pernikahan: false,
	// 	Pendidikan: []model.Pendidikan{
	// 		{
	// 			Nama: "Sarjana",
	// 			Jurusan: "Sistem Informasi",
	// 		},
	// 		{
	// 			Nama: "SMA",
	// 			Jurusan: "MIPA",
	// 		},
	// 	},
	// }
	// fmt.Println(myData)

	// scanner := bufio.NewScanner(os.Stdin)

	// for {
	// 	fmt.Println("\n=== MENU UTAMA KODA B9 ===")
	// 	fmt.Println("1. Hitung Rectangle (Perimeter & Area)")
	// 	fmt.Println("2. Generate Window Pattern")
	// 	fmt.Println("3. Inject to Slice")
	// 	fmt.Println("4. Lihat Biodata")
	// 	fmt.Println("5. Lihat Isi File")
	// 	fmt.Println("0. Keluar")
	// 	fmt.Print("Pilih menu (0-4): ")

	// 	if !scanner.Scan() {
	// 		break
	// 	}
	// 	choice := strings.TrimSpace(scanner.Text())

	// 	switch choice {
	// 	case "1":
	// 		fmt.Print("Masukkan Width (angka): ")
	// 		scanner.Scan()
	// 		w, _ := strconv.Atoi(scanner.Text())

	// 		fmt.Print("Masukkan Height (angka): ")
	// 		scanner.Scan()
	// 		h, _ := strconv.Atoi(scanner.Text())

	// 		perimeter, area := service.CountRectangle(uint8(w), uint8(h))
	// 		fmt.Printf("Perimeter: %d\nArea: %d", perimeter, area)
	// 	case "2":
	// 		fmt.Print("Masukkan Ukuran Jendela (angka): ")
	// 		scanner.Scan()
	// 		s, _ := strconv.Atoi(scanner.Text())

	// 		err := service.GenerateWindow(s)
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}
	// 	case "3":
	// 		fmt.Println("Simulating Inject Element into Exact Position in a Slice")
	// 		service.InjectToSlice()
	// 	case "4":
	// 		myData := model.Biodata{
	// 			Nama:       "Ridho Contrisha",
	// 			Foto:       "",
	// 			Email:      "rcontrisha@gmail.com",
	// 			Umur:       23,
	// 			Telepon:    "081315468293",
	// 			Pernikahan: false,
	// 			Pendidikan: []model.Pendidikan{
	// 				{
	// 					Nama:    "Sarjana",
	// 					Jurusan: "Sistem Informasi",
	// 				},
	// 				{
	// 					Nama:    "SMA",
	// 					Jurusan: "MIPA",
	// 				},
	// 			},
	// 		}
	// 		fmt.Printf("=== BIODATA ===\n"+
	// 			"Nama       : %s\n"+
	// 			"Email      : %s\n"+
	// 			"Umur       : %d\n"+
	// 			"Telepon    : %s\n"+
	// 			"Pendidikan : 1. %s (%s), 2. %s (%s)\n",
	// 			myData.Nama, myData.Email, myData.Umur, myData.Telepon,
	// 			myData.Pendidikan[0].Nama, myData.Pendidikan[0].Jurusan,
	// 			myData.Pendidikan[1].Nama, myData.Pendidikan[1].Jurusan,
	// 		)
	// 	case "5":
	// 		fmt.Print("Input File Path (Pastikan path mengarah ke file, bukan directory):\n")
	// 		scanner.Scan()
	// 		path := scanner.Text()
	// 		err := service.FileReader(path)
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}
	// 	case "0":
	// 		fmt.Println("Keluar dari program.")
	// 		return

	// 	default:
	// 		fmt.Println("Menu tidak valid, silakan pilih lagi.")
	// 	}
	// }

	// Create new instance
	// person1 := model.NewPerson("Ucup", "Bogor", "081375757283")
	// // Getter method to print person data
	// nama, address, phone := person1.Print()
	// fmt.Printf("Name: %s\nAddress: %s\nPhone: %s\n", nama, address, phone)
	// // Getter method to print "Hello, {person_name}"
	// person1.Greet()
	// // Setter method to change person's name
	// person1.ChangeName("Gipen")
	// person1.Greet()

	// bankData := checkout.Bank{
	// 	Name: "Bank",
	// }

	// onlineData := checkout.Online{
	// 	Name: "Online",
	// }

	// fiktifData := checkout.Fiktif{
	// 	Name: "Fiktif",
	// 	Subtotals: []uint32{},
	// }

	// checkout.HandlePayment(&bankData, []uint32{50000, 30000, 40000})
	// checkout.HandlePayment(&onlineData, []uint32{150000, 30000, 0})

	// checkout.HandlePayment(&fiktifData, []uint32{50000, 0})
	// checkout.HandlePayment(&fiktifData, []uint32{150000, 20000, 30000})
	// fmt.Println("Total", fiktifData.CalcTotal(fiktifData.Subtotals))

	// service.Rutinitas()

	service.Run()
}
