package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
	qr "github.com/skip2/go-qrcode"
)

func main() {
	// Get CSV file path from user
	var csvFilePath string
	fmt.Print("Enter the path to the CSV file: ")
	fmt.Scanln(&csvFilePath)

	// Read the CSV file
	file, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Generate QR codes and add to the PDF
	for i, record := range records {
		if i == 0 {
			continue // Skip the first record (header row)
		}
		text := record[findColumnIndex(records[0], "Lpa")] // Assuming the header 'Lpa' is in the first row

		// Generate QR code image as a byte slice
		qrCode, err := qr.Encode(text, qr.Medium, 256)
		if err != nil {
			log.Fatal(err)
		}

		// Create an io.Reader from the QR code image byte slice
		qrCodeReader := bytes.NewReader(qrCode)

		// Add QR code image to the PDF as a new page
		pdf.AddPage()
		pdf.RegisterImageOptionsReader(text, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader)
		pdf.ImageOptions(text, 10, 10, 0, 0, true, gofpdf.ImageOptions{ImageType: "png"}, 0, "")
	}

	// Save the PDF file
	outputPath := "output.pdf"
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("QR codes generated and saved to %s\n", outputPath)
}

// Helper function to find the index of a column in the header row
func findColumnIndex(headerRow []string, column string) int {
	for i, value := range headerRow {
		if value == column {
			return i
		}
	}
	return -1
}
