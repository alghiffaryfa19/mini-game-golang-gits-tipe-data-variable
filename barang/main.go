package main

import (
	"fmt"
	"strconv"
)


func main() {

	var barang string
    var harga int
    var qty int

	fmt.Println("Nama Barang : ")
	fmt.Scan(&barang)

    fmt.Println("Harga : ")
	fmt.Scan(&harga)

    fmt.Println("Jumlah : ")
	fmt.Scan(&qty)

    total := harga*qty

	jumlah := strconv.Itoa(qty)
	satuan := strconv.Itoa(harga)
	x := strconv.Itoa(total)

	

	v := fmt.Sprintf("Barang %s berjumlah %s memiliki harga satuan %s dengan harga total %s",barang,jumlah,satuan,x)

    fmt.Println(v)
}

