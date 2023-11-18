package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// inisial pertanyaan dan jawaban
	Question := []string{
		"1. Apakah kamu manusia?",
		"2. Berapa hasil dari 22/7 * (14&14)/2",
		"3. Jika Tinggi tiang bendera adalah 6m dan jarak antara kaki kita dengan tiang adalah 4 meter,tentukan panjang kemiringan sisi antara tiang bendera dengan kita",
	}

	Option := [][]string{
		{"A. Iya", "B. Mungkin", "C. Aku adalah titisan ultraman"},
		{"A. 308", "B. 264", "C. 148,5"},
		{"A. 14,23", "B. 7,21", "C. Tidak Tahu"},
	}

	correctAnswer := []string{"a", "a", "b"}

	// jumlah pertanyaan
	totalQuestions := len(Question)
	skor := 0

	// Perulangan pertanyaan
	for i := 0; i < totalQuestions; i++ {
		// menampilkan pertanyaan dan pilihan jawaban
		fmt.Println(Question[i])
		for _, pilihanJawaban := range Option[i] {
			fmt.Println(pilihanJawaban)
		}

		// Masukkan jawabn mu
		fmt.Print("Jawaban: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		Answer := strings.ToLower(scanner.Text())

		// memeriksa jawaban yang ditulis
		if Answer == correctAnswer[i] {
			fmt.Println("Benar!")
			skor++
		} else {
			fmt.Println("Salah. Jawaban yang benar adalah:", strings.ToUpper(correctAnswer[i]))
		}
		fmt.Println()
	}

	// skor total
	fmt.Println("Skor akhir:", skor, "/", totalQuestions)
}
