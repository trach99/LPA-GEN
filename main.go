package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time" //TEMP

	"github.com/jung-kurt/gofpdf"
	qr "github.com/skip2/go-qrcode"
)

func main() {

	// Get CSV file path from user
	var csvFilePath string
	fmt.Print("Enter the path to the CSV file: ")
	fmt.Scanln(&csvFilePath)

	//RUNTIME measure TEMP
	start := time.Now()

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
	//pdf.SetMargins(10, 10, 10)

	// Load watermark image
	watermarkPath := "Workz_Logo_2022-Blue-short.png"
	pdf.RegisterImageOptions(watermarkPath, gofpdf.ImageOptions{ImageType: "png"})

	// Generate QR codes and add to the PDF
	for i, record := range records {
		if i == 0 {
			continue // Skip the first record (header row)
		}
		text := record[findColumnIndex(records[0], "Lpa")]    // Assuming the header 'Lpa' is in the first row
		iccid := record[findColumnIndex(records[0], "ICCID")] // Assuming the header 'ICCID' is in the first row

		// Generate QR code image as a byte slice
		qrCode, err := qr.Encode(text, qr.Medium, 256)
		if err != nil {
			log.Fatal(err)
		}

		// Create an io.Reader from the QR code image byte slice
		qrCodeReader := bytes.NewReader(qrCode)

		// Add page and watermark image to the PDF
		pdf.AddPage()

		// Add watermark image
		pdf.SetAlpha(1, "Multiply")
		pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

		// Reset alpha settings
		pdf.SetAlpha(1.0, "")

		//Rest PDF margins
		pdf.SetMargins(10, 13, 10)

		// Add QR code image to the PDF as a new page
		pdf.RegisterImageOptionsReader(text, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader)
		pdf.ImageOptions(text, 20, 10, 0, 0, true, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

		// Print the 'Lpa' and 'ICCID' strings below the QR code
		pdf.SetFont("Arial", "", 10)
		pdf.Text(10, 85, text)
		pdf.Text(10, 92, "ICCID: "+iccid)
	}

	// Save the PDF file
	outputPath := "output.pdf"
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("QR codes generated and saved to %s\n", outputPath)

	// calculate to exe time TEMP
	elapsed := time.Since(start)
	fmt.Printf("TEST-BENCHMARKS took %s", elapsed)
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
