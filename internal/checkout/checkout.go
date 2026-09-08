package checkout

import "fmt"

type Checkout interface {
	Payment(items []uint32) error
}

func HandlePayment(channel Checkout, items []uint32) error {
	content := channel.Payment(items)
	// fmt.Println(content)
	if content != nil {
		fmt.Println(content.Error())
	}
	return content
}

type Bank struct {
	Name string
}

func (b *Bank) Payment(items []uint32) error {
	var payment uint32 = 0

	for _, i := range items {
		if i <= 0 {
			return fmt.Errorf("Harga Tidak Boleh/Kurang dari 0")
		}
		payment += i
	}

	fmt.Printf("Pembayaran sebesar %d berhasil dilakukan via %s\n", payment, b.Name)

	return nil
}

type Online struct {
	Name string
}

func (o *Online) Payment(items []uint32) error {
	var payment uint32 = 0

	for _, i := range items {
		if i <= 0 {
			return fmt.Errorf("Harga Tidak Boleh/Kurang dari 0")
		}
		payment += i
	}

	fmt.Printf("Pembayaran sebesar %d berhasil dilakukan via %s\n", payment, o.Name)

	return nil
}

type Fiktif struct {
	Name      string
	Subtotals []uint32
}

func (f *Fiktif) Payment(items []uint32) error {
	for _, i := range items {
		if i <= 0 {
			return fmt.Errorf("Harga Tidak Boleh/Kurang dari 0")
		}
		f.Subtotals = append(f.Subtotals, i)
	}

	return nil
}

func (f *Fiktif) CalcTotal(subtotals []uint32) (total uint32) {
	total = 0
	for _, i := range subtotals {
		total += i
	}
	return total
}
