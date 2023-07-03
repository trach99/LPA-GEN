package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time" //TEMP

	"github.com/cheggaaa/pb/v3" //PROGRESS BAR
	"github.com/jung-kurt/gofpdf"
	qr "github.com/skip2/go-qrcode"
)

func main() {

	// Get CSV file path from user
	var csvFilePath string
	fmt.Print("Enter the path to the CSV file: ")
	fmt.Scanln(&csvFilePath)

	// Get Customer input from user
	var customer string
	fmt.Print("Enter the Customer name: ")
	fmt.Scanln(&customer)

	// Get Profile Type input from user
	var profileType string
	fmt.Print("Enter the Profile Type: ")
	fmt.Scanln(&profileType)

	//RUNTIME measure TEMP
	start := time.Now()

	// Read the CSV file
	file, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the CSV records
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

	pagelength := (len(records) / 10)

	//Check if page nos. is not a multiple of 10
	if (pagelength % 10) != 0 {
		pagelength += 1
	}

	pagelength += 1 //Due to Cover Page

	//--------------------------------------------------------------COVER PAGE--------------------------------------------------------------------//
	// Add the cover page with the table
	pdf.AddPage()

	// Add watermark image
	pdf.SetAlpha(1, "Multiply")
	pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

	// Reset alpha settings
	pdf.SetAlpha(1.0, "")

	// Set font and font size for the cover page
	pdf.SetFont("Arial", "B", 16)

	pdf.SetY(30) // Since the header initially was aligned on the top

	// Add header text
	headerText := customer + " eSIM RSP - QR Codes"
	pdf.CellFormat(0, 20, headerText, "", 1, "C", false, 0, "")

	pdf.SetLeftMargin(26) //Since the table initially was getting left aligned

	// Set font size for the table
	//pdf.SetFont("Arial", "B", 12)
	pdf.SetFont("Arial", "", 12)

	// Add additional table content
	pdf.CellFormat(80, 10, "Customer", "1", 0, "L", false, 0, "")
	pdf.CellFormat(80, 10, customer, "1", 1, "L", false, 0, "")

	// Add additional table content
	pdf.CellFormat(80, 10, "Profile Type", "1", 0, "L", false, 0, "")
	pdf.CellFormat(80, 10, profileType, "1", 1, "L", false, 0, "")

	// Fetch the first and last iccid from the CSV records
	firstICCID := records[1][findColumnIndex(records[0], "iccid")]
	lastICCID := records[len(records)-1][findColumnIndex(records[0], "iccid")]

	// Add additional table content
	pdf.CellFormat(80, 10, "Start ICCID", "1", 0, "L", false, 0, "")
	pdf.CellFormat(80, 10, firstICCID, "1", 1, "L", false, 0, "")
	pdf.CellFormat(80, 10, "End ICCID", "1", 0, "L", false, 0, "")
	pdf.CellFormat(80, 10, lastICCID, "1", 1, "L", false, 0, "")

	// Calculate the quantity of records in the CSV
	quantity := len(records) - 1                    // Subtract 1 to exclude the header row
	quantityStr := formatNumberWithCommas(quantity) //commas for the 1000's :-)
	pdf.CellFormat(80, 10, "Quantity", "1", 0, "L", false, 0, "")
	pdf.CellFormat(80, 10, fmt.Sprint(quantityStr), "1", 1, "L", false, 0, "")

	// Get the current date
	currentDate := time.Now().Format("02-Jan-2006")
	pdf.CellFormat(80, 10, "Release Date", "1", 0, "L", false, 0, "")
	pdf.CellFormat(80, 10, currentDate, "1", 1, "L", false, 0, "")

	// Add page number footer text
	pdf.SetFont("Arial", "I", 8)
	pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))

	//--------------------------------------------------------------COVER PAGE--------------------------------------------------------------------//

	// Create a progress bar
	bar := pb.StartNew(pagelength - 1) //need 100% :-)
	bar.SetWidth(80)

	// Generate QR codes and add to the PDF
	for i := 1; i < len(records); i += 10 {
		// Get the first QR code and text pair
		record1 := records[i]
		text1 := record1[findColumnIndex(records[0], "lpa_string")]
		iccid1 := record1[findColumnIndex(records[0], "iccid")]

		// Generate QR code image as a byte slice
		qrCode1, err := qr.Encode(text1, qr.Medium, 150)
		if err != nil {
			log.Fatal(err)
		}

		// Create an io.Reader from the QR code image byte slice
		qrCodeReader1 := bytes.NewReader(qrCode1)

		// Get the second QR code and text pair
		if i+1 < len(records) {
			record2 := records[i+1]
			text2 := record2[findColumnIndex(records[0], "lpa_string")]
			iccid2 := record2[findColumnIndex(records[0], "iccid")]

			// Generate QR code image as a byte slice
			qrCode2, err := qr.Encode(text2, qr.Medium, 150)
			if err != nil {
				log.Fatal(err)
			}

			// Create an io.Reader from the QR code image byte slice
			qrCodeReader2 := bytes.NewReader(qrCode2)

			// Get the third QR code and text pair
			if i+2 < len(records) {
				record3 := records[i+2]
				text3 := record3[findColumnIndex(records[0], "lpa_string")]
				iccid3 := record3[findColumnIndex(records[0], "iccid")]

				// Generate QR code image as a byte slice
				qrCode3, err := qr.Encode(text3, qr.Medium, 150)
				if err != nil {
					log.Fatal(err)
				}

				// Create an io.Reader from the QR code image byte slice
				qrCodeReader3 := bytes.NewReader(qrCode3)

				// Get the fourth QR code and text pair
				if i+3 < len(records) {
					record4 := records[i+3]
					text4 := record4[findColumnIndex(records[0], "lpa_string")]
					iccid4 := record4[findColumnIndex(records[0], "iccid")]

					// Generate QR code image as a byte slice
					qrCode4, err := qr.Encode(text4, qr.Medium, 150)
					if err != nil {
						log.Fatal(err)
					}

					// Create an io.Reader from the QR code image byte slice
					qrCodeReader4 := bytes.NewReader(qrCode4)

					// Get the fifth QR code and text pair
					if i+4 < len(records) {
						record5 := records[i+4]
						text5 := record5[findColumnIndex(records[0], "lpa_string")]
						iccid5 := record5[findColumnIndex(records[0], "iccid")]

						// Generate QR code image as a byte slice
						qrCode5, err := qr.Encode(text5, qr.Medium, 150)
						if err != nil {
							log.Fatal(err)
						}
						// Create an io.Reader from the QR code image byte slice
						qrCodeReader5 := bytes.NewReader(qrCode5)

						// Get the sixth QR code and text pair
						if i+5 < len(records) {
							record6 := records[i+5]
							text6 := record6[findColumnIndex(records[0], "lpa_string")]
							iccid6 := record6[findColumnIndex(records[0], "iccid")]

							// Generate QR code image as a byte slice
							qrCode6, err := qr.Encode(text6, qr.Medium, 150)
							if err != nil {
								log.Fatal(err)
							}
							// Create an io.Reader from the QR code image byte slice
							qrCodeReader6 := bytes.NewReader(qrCode6)

							// Get the seventh QR code and text pair
							if i+6 < len(records) {
								record7 := records[i+6]
								text7 := record7[findColumnIndex(records[0], "lpa_string")]
								iccid7 := record7[findColumnIndex(records[0], "iccid")]

								// Generate QR code image as a byte slice
								qrCode7, err := qr.Encode(text7, qr.Medium, 150)
								if err != nil {
									log.Fatal(err)
								}
								// Create an io.Reader from the QR code image byte slice
								qrCodeReader7 := bytes.NewReader(qrCode7)

								// Get the eighth QR code and text pair
								if i+7 < len(records) {
									record8 := records[i+7]
									text8 := record8[findColumnIndex(records[0], "lpa_string")]
									iccid8 := record8[findColumnIndex(records[0], "iccid")]

									// Generate QR code image as a byte slice
									qrCode8, err := qr.Encode(text8, qr.Medium, 150)
									if err != nil {
										log.Fatal(err)
									}
									// Create an io.Reader from the QR code image byte slice
									qrCodeReader8 := bytes.NewReader(qrCode8)

									// Get the ninth QR code and text pair
									if i+8 < len(records) {
										record9 := records[i+8]
										text9 := record9[findColumnIndex(records[0], "lpa_string")]
										iccid9 := record9[findColumnIndex(records[0], "iccid")]

										// Generate QR code image as a byte slice
										qrCode9, err := qr.Encode(text9, qr.Medium, 150)
										if err != nil {
											log.Fatal(err)
										}
										// Create an io.Reader from the QR code image byte slice
										qrCodeReader9 := bytes.NewReader(qrCode9)

										// Get the tenth QR code and text pair
										if i+9 < len(records) {
											record10 := records[i+9]
											text10 := record10[findColumnIndex(records[0], "lpa_string")]
											iccid10 := record10[findColumnIndex(records[0], "iccid")]

											// Generate QR code image as a byte slice
											qrCode10, err := qr.Encode(text10, qr.Medium, 150)
											if err != nil {
												log.Fatal(err)
											}
											// Create an io.Reader from the QR code image byte slice
											qrCodeReader10 := bytes.NewReader(qrCode10)

											// Add page and watermark image to the PDF
											pdf.AddPage()

											// Add watermark image
											pdf.SetAlpha(1, "Multiply")
											pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Reset alpha settings
											pdf.SetAlpha(1.0, "")

											// Reset PDF margins
											pdf.SetMargins(10, 14, 10)

											// Add first QR code image to the PDF
											pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
											pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the first 'lpa_string' and 'iccid' strings below the first QR code
											pdf.SetFont("Arial", "", 8)
											pdf.Text(20, 55, text1)
											pdf.Text(20, 62, "iccid: "+iccid1)

											// Add second QR code image to the PDF
											pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
											pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the second 'lpa_string' and 'iccid' strings below the second QR code
											pdf.Text(115, 55, text2)
											pdf.Text(115, 62, "iccid: "+iccid2)

											// Add third QR code image to the PDF
											pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
											pdf.ImageOptions(text3, 30, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the third 'lpa_string' and 'iccid' strings below the third QR code
											pdf.Text(20, 111, text3)
											pdf.Text(20, 118, "iccid: "+iccid3)
											// Add fourth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
											pdf.ImageOptions(text4, 125, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the fourth 'lpa_string' and 'iccid' strings below the fourth QR code
											pdf.Text(115, 111, text4)
											pdf.Text(115, 118, "iccid: "+iccid4)

											// Add fifth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
											pdf.ImageOptions(text5, 30, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the fifth 'lpa_string' and 'iccid' strings below the fifth QR code
											pdf.Text(20, 167, text5)
											pdf.Text(20, 174, "iccid: "+iccid5)

											// Add sixth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
											pdf.ImageOptions(text6, 125, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the sixth 'lpa_string' and 'iccid' strings below the sixth QR code
											pdf.Text(115, 167, text6)
											pdf.Text(115, 174, "iccid: "+iccid6)

											// Add seventh QR code image to the PDF
											pdf.RegisterImageOptionsReader(text7, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader7)
											pdf.ImageOptions(text7, 30, 182, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the seventh 'lpa_string' and 'iccid' strings below the seventh QR code
											pdf.Text(20, 223, text7)
											pdf.Text(20, 230, "iccid: "+iccid7)

											// Add eighth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text8, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader8)
											pdf.ImageOptions(text8, 125, 182, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the eighth 'lpa_string' and 'iccid' strings below the eighth QR code
											pdf.Text(115, 223, text8)
											pdf.Text(115, 230, "iccid: "+iccid8)

											// Add ninth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text9, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader9)
											pdf.ImageOptions(text9, 30, 238, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the ninth 'lpa_string' and 'iccid' strings below the ninth QR code
											pdf.Text(20, 279, text9)
											pdf.Text(20, 286, "iccid: "+iccid9)

											// Add tenth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text10, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader10)
											pdf.ImageOptions(text10, 125, 238, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the tenth 'lpa_string' and 'iccid' strings below the tenth QR code
											pdf.Text(115, 279, text10)
											pdf.Text(115, 286, "iccid: "+iccid10)

											// Add page number footer text
											pdf.SetFont("Arial", "I", 8)
											pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))

										} else {
											// Only nine records left, add them to the PDF
											pdf.AddPage()

											// Add watermark image
											pdf.SetAlpha(1, "Multiply")
											pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Reset alpha settings
											pdf.SetAlpha(1.0, "")

											// Reset PDF margins
											pdf.SetMargins(10, 14, 10)

											// Add first QR code image to the PDF
											pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
											pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the first 'lpa_string' and 'iccid' strings below the first QR code
											pdf.SetFont("Arial", "", 8)
											pdf.Text(20, 55, text1)
											pdf.Text(20, 62, "iccid: "+iccid1)

											// Add second QR code image to the PDF
											pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
											pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the second 'lpa_string' and 'iccid' strings below the second QR code
											pdf.Text(115, 55, text2)
											pdf.Text(115, 62, "iccid: "+iccid2)

											// Add third QR code image to the PDF
											pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
											pdf.ImageOptions(text3, 30, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the third 'lpa_string' and 'iccid' strings below the third QR code
											pdf.Text(20, 111, text3)
											pdf.Text(20, 118, "iccid: "+iccid3)
											// Add fourth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
											pdf.ImageOptions(text4, 125, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the fourth 'lpa_string' and 'iccid' strings below the fourth QR code
											pdf.Text(115, 111, text4)
											pdf.Text(115, 118, "iccid: "+iccid4)

											// Add fifth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
											pdf.ImageOptions(text5, 30, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the fifth 'lpa_string' and 'iccid' strings below the fifth QR code
											pdf.Text(20, 167, text5)
											pdf.Text(20, 174, "iccid: "+iccid5)

											// Add sixth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
											pdf.ImageOptions(text6, 125, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the sixth 'lpa_string' and 'iccid' strings below the sixth QR code
											pdf.Text(115, 167, text6)
											pdf.Text(115, 174, "iccid: "+iccid6)

											// Add seventh QR code image to the PDF
											pdf.RegisterImageOptionsReader(text7, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader7)
											pdf.ImageOptions(text7, 30, 182, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the seventh 'lpa_string' and 'iccid' strings below the seventh QR code
											pdf.Text(20, 223, text7)
											pdf.Text(20, 230, "iccid: "+iccid7)

											// Add eighth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text8, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader8)
											pdf.ImageOptions(text8, 125, 182, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the eighth 'lpa_string' and 'iccid' strings below the eighth QR code
											pdf.Text(115, 223, text8)
											pdf.Text(115, 230, "iccid: "+iccid8)

											// Add ninth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text9, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader9)
											pdf.ImageOptions(text9, 30, 238, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the ninth 'lpa_string' and 'iccid' strings below the ninth QR code
											pdf.Text(20, 279, text9)
											pdf.Text(20, 286, "iccid: "+iccid9)

											// Add page number footer text
											pdf.SetFont("Arial", "I", 8)
											pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))

										}
									} else {
										// Only eight records left, add them to the PDF
										pdf.AddPage()

										// Add watermark image
										pdf.SetAlpha(1, "Multiply")
										pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Reset alpha settings
										pdf.SetAlpha(1.0, "")

										// Reset PDF margins
										pdf.SetMargins(10, 14, 10)

										// Add first QR code image to the PDF
										pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
										pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the first 'lpa_string' and 'iccid' strings below the first QR code
										pdf.SetFont("Arial", "", 8)
										pdf.Text(20, 55, text1)
										pdf.Text(20, 62, "iccid: "+iccid1)

										// Add second QR code image to the PDF
										pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
										pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the second 'lpa_string' and 'iccid' strings below the second QR code
										pdf.Text(115, 55, text2)
										pdf.Text(115, 62, "iccid: "+iccid2)

										// Add third QR code image to the PDF
										pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
										pdf.ImageOptions(text3, 30, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the third 'lpa_string' and 'iccid' strings below the third QR code
										pdf.Text(20, 111, text3)
										pdf.Text(20, 118, "iccid: "+iccid3)
										// Add fourth QR code image to the PDF
										pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
										pdf.ImageOptions(text4, 125, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the fourth 'lpa_string' and 'iccid' strings below the fourth QR code
										pdf.Text(115, 111, text4)
										pdf.Text(115, 118, "iccid: "+iccid4)

										// Add fifth QR code image to the PDF
										pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
										pdf.ImageOptions(text5, 30, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the fifth 'lpa_string' and 'iccid' strings below the fifth QR code
										pdf.Text(20, 167, text5)
										pdf.Text(20, 174, "iccid: "+iccid5)

										// Add sixth QR code image to the PDF
										pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
										pdf.ImageOptions(text6, 125, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the sixth 'lpa_string' and 'iccid' strings below the sixth QR code
										pdf.Text(115, 167, text6)
										pdf.Text(115, 174, "iccid: "+iccid6)

										// Add seventh QR code image to the PDF
										pdf.RegisterImageOptionsReader(text7, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader7)
										pdf.ImageOptions(text7, 30, 182, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the seventh 'lpa_string' and 'iccid' strings below the seventh QR code
										pdf.Text(20, 223, text7)
										pdf.Text(20, 230, "iccid: "+iccid7)

										// Add eighth QR code image to the PDF
										pdf.RegisterImageOptionsReader(text8, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader8)
										pdf.ImageOptions(text8, 125, 182, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the eighth 'lpa_string' and 'iccid' strings below the eighth QR code
										pdf.Text(115, 223, text8)
										pdf.Text(115, 230, "iccid: "+iccid8)

										// Add page number footer text
										pdf.SetFont("Arial", "I", 8)
										pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))
									}
								} else {
									// Only seven records left, add them to the PDF
									pdf.AddPage()

									// Add watermark image
									pdf.SetAlpha(1, "Multiply")
									pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Reset alpha settings
									pdf.SetAlpha(1.0, "")

									// Reset PDF margins
									pdf.SetMargins(10, 14, 10)

									// Add first QR code image to the PDF
									pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
									pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the first 'lpa_string' and 'iccid' strings below the first QR code
									pdf.SetFont("Arial", "", 8)
									pdf.Text(20, 55, text1)
									pdf.Text(20, 62, "iccid: "+iccid1)

									// Add second QR code image to the PDF
									pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
									pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the second 'lpa_string' and 'iccid' strings below the second QR code
									pdf.Text(115, 55, text2)
									pdf.Text(115, 62, "iccid: "+iccid2)

									// Add third QR code image to the PDF
									pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
									pdf.ImageOptions(text3, 30, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the third 'lpa_string' and 'iccid' strings below the third QR code
									pdf.Text(20, 111, text3)
									pdf.Text(20, 118, "iccid: "+iccid3)
									// Add fourth QR code image to the PDF
									pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
									pdf.ImageOptions(text4, 125, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the fourth 'lpa_string' and 'iccid' strings below the fourth QR code
									pdf.Text(115, 111, text4)
									pdf.Text(115, 118, "iccid: "+iccid4)

									// Add fifth QR code image to the PDF
									pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
									pdf.ImageOptions(text5, 30, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the fifth 'lpa_string' and 'iccid' strings below the fifth QR code
									pdf.Text(20, 167, text5)
									pdf.Text(20, 174, "iccid: "+iccid5)

									// Add sixth QR code image to the PDF
									pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
									pdf.ImageOptions(text6, 125, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the sixth 'lpa_string' and 'iccid' strings below the sixth QR code
									pdf.Text(115, 167, text6)
									pdf.Text(115, 174, "iccid: "+iccid6)

									// Add seventh QR code image to the PDF
									pdf.RegisterImageOptionsReader(text7, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader7)
									pdf.ImageOptions(text7, 30, 182, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the seventh 'lpa_string' and 'iccid' strings below the seventh QR code
									pdf.Text(20, 223, text7)
									pdf.Text(20, 230, "iccid: "+iccid7)

									// Add page number footer text
									pdf.SetFont("Arial", "I", 8)
									pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))
								}
							} else {
								// Only six records left, add them to the PDF
								pdf.AddPage()

								// Add watermark image
								pdf.SetAlpha(1, "Multiply")
								pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Reset alpha settings
								pdf.SetAlpha(1.0, "")

								// Reset PDF margins
								pdf.SetMargins(10, 14, 10)

								// Add first QR code image to the PDF
								pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
								pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the first 'lpa_string' and 'iccid' strings below the first QR code
								pdf.SetFont("Arial", "", 8)
								pdf.Text(20, 55, text1)
								pdf.Text(20, 62, "iccid: "+iccid1)

								// Add second QR code image to the PDF
								pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
								pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the second 'lpa_string' and 'iccid' strings below the second QR code
								pdf.Text(115, 55, text2)
								pdf.Text(115, 62, "iccid: "+iccid2)

								// Add third QR code image to the PDF
								pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
								pdf.ImageOptions(text3, 30, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the third 'lpa_string' and 'iccid' strings below the third QR code
								pdf.Text(20, 111, text3)
								pdf.Text(20, 118, "iccid: "+iccid3)
								// Add fourth QR code image to the PDF
								pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
								pdf.ImageOptions(text4, 125, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the fourth 'lpa_string' and 'iccid' strings below the fourth QR code
								pdf.Text(115, 111, text4)
								pdf.Text(115, 118, "iccid: "+iccid4)

								// Add fifth QR code image to the PDF
								pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
								pdf.ImageOptions(text5, 30, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the fifth 'lpa_string' and 'iccid' strings below the fifth QR code
								pdf.Text(20, 167, text5)
								pdf.Text(20, 174, "iccid: "+iccid5)

								// Add sixth QR code image to the PDF
								pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
								pdf.ImageOptions(text6, 125, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the sixth 'lpa_string' and 'iccid' strings below the sixth QR code
								pdf.Text(115, 167, text6)
								pdf.Text(115, 174, "iccid: "+iccid6)

								// Add page number footer text
								pdf.SetFont("Arial", "I", 8)
								pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))
							}
						} else {
							// Only five records left, add them to the PDF
							pdf.AddPage()

							// Add watermark image
							pdf.SetAlpha(1, "Multiply")
							pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Reset alpha settings
							pdf.SetAlpha(1.0, "")

							// Reset PDF margins
							pdf.SetMargins(10, 14, 10)

							// Add first QR code image to the PDF
							pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
							pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the first 'lpa_string' and 'iccid' strings below the first QR code
							pdf.SetFont("Arial", "", 8)
							pdf.Text(20, 55, text1)
							pdf.Text(20, 62, "iccid: "+iccid1)

							// Add second QR code image to the PDF
							pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
							pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the second 'lpa_string' and 'iccid' strings below the second QR code
							pdf.Text(115, 55, text2)
							pdf.Text(115, 62, "iccid: "+iccid2)

							// Add third QR code image to the PDF
							pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
							pdf.ImageOptions(text3, 30, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the third 'lpa_string' and 'iccid' strings below the third QR code
							pdf.Text(20, 111, text3)
							pdf.Text(20, 118, "iccid: "+iccid3)
							// Add fourth QR code image to the PDF
							pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
							pdf.ImageOptions(text4, 125, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the fourth 'lpa_string' and 'iccid' strings below the fourth QR code
							pdf.Text(115, 111, text4)
							pdf.Text(115, 118, "iccid: "+iccid4)

							// Add fifth QR code image to the PDF
							pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
							pdf.ImageOptions(text5, 30, 126, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the fifth 'lpa_string' and 'iccid' strings below the fifth QR code
							pdf.Text(20, 167, text5)
							pdf.Text(20, 174, "iccid: "+iccid5)

							// Add page number footer text
							pdf.SetFont("Arial", "I", 8)
							pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))
						}
					} else {
						// Only four records left, add them to the PDF
						pdf.AddPage()

						// Add watermark image
						pdf.SetAlpha(1, "Multiply")
						pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Reset alpha settings
						pdf.SetAlpha(1.0, "")

						// Reset PDF margins
						pdf.SetMargins(10, 14, 10)

						// Add first QR code image to the PDF
						pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
						pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Print the first 'lpa_string' and 'iccid' strings below the first QR code
						pdf.SetFont("Arial", "", 8)
						pdf.Text(20, 55, text1)
						pdf.Text(20, 62, "iccid: "+iccid1)

						// Add second QR code image to the PDF
						pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
						pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Print the second 'lpa_string' and 'iccid' strings below the second QR code
						pdf.Text(115, 55, text2)
						pdf.Text(115, 62, "iccid: "+iccid2)

						// Add third QR code image to the PDF
						pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
						pdf.ImageOptions(text3, 30, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Print the third 'lpa_string' and 'iccid' strings below the third QR code
						pdf.Text(20, 111, text3)
						pdf.Text(20, 118, "iccid: "+iccid3)
						// Add fourth QR code image to the PDF
						pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
						pdf.ImageOptions(text4, 125, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Print the fourth 'lpa_string' and 'iccid' strings below the fourth QR code
						pdf.Text(115, 111, text4)
						pdf.Text(115, 118, "iccid: "+iccid4)

						// Add page number footer text
						pdf.SetFont("Arial", "I", 8)
						pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))

					}
				} else {
					// Only three records left, add them to the PDF
					pdf.AddPage()

					// Add watermark image
					pdf.SetAlpha(1, "Multiply")
					pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

					// Reset alpha settings
					pdf.SetAlpha(1.0, "")

					// Reset PDF margins
					pdf.SetMargins(10, 14, 10)

					// Add first QR code image to the PDF
					pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
					pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

					// Print the first 'lpa_string' and 'iccid' strings below the first QR code
					pdf.SetFont("Arial", "", 8)
					pdf.Text(20, 55, text1)
					pdf.Text(20, 62, "iccid: "+iccid1)

					// Add second QR code image to the PDF
					pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
					pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

					// Print the second 'lpa_string' and 'iccid' strings below the second QR code
					pdf.Text(115, 55, text2)
					pdf.Text(115, 62, "iccid: "+iccid2)

					// Add third QR code image to the PDF
					pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
					pdf.ImageOptions(text3, 30, 70, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

					// Print the third 'lpa_string' and 'iccid' strings below the third QR code
					pdf.Text(20, 111, text3)
					pdf.Text(20, 118, "iccid: "+iccid3)

					// Add page number footer text
					pdf.SetFont("Arial", "I", 8)
					pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))
				}
			} else {
				// Only two records left, add them to the PDF
				pdf.AddPage()

				// Add watermark image
				pdf.SetAlpha(1, "Multiply")
				pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

				// Reset alpha settings
				pdf.SetAlpha(1.0, "")

				// Reset PDF margins
				pdf.SetMargins(10, 14, 10)

				// Add first QR code image to the PDF
				pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
				pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

				// Print the first 'lpa_string' and 'iccid' strings below the first QR code
				pdf.SetFont("Arial", "", 8)
				pdf.Text(20, 55, text1)
				pdf.Text(20, 62, "iccid: "+iccid1)

				// Add second QR code image to the PDF
				pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
				pdf.ImageOptions(text2, 125, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

				// Print the second 'lpa_string' and 'iccid' strings below the second QR code
				pdf.Text(115, 55, text2)
				pdf.Text(115, 62, "iccid: "+iccid2)

				// Add page number footer text
				pdf.SetFont("Arial", "I", 8)
				pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))
			}
		} else {
			// Only one record left, add it to the PDF
			pdf.AddPage()

			// Add watermark image
			pdf.SetAlpha(1, "Multiply")
			pdf.ImageOptions(watermarkPath, 8, 0, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

			// Reset alpha settings
			pdf.SetAlpha(1.0, "")

			// Reset PDF margins
			pdf.SetMargins(10, 14, 10)

			// Add first QR code image to the PDF
			pdf.RegisterImageOptionsReader(text1, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader1)
			pdf.ImageOptions(text1, 30, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

			// Print the first 'lpa_string' and 'iccid' strings below the first QR code
			pdf.SetFont("Arial", "", 8)
			pdf.Text(20, 55, text1)
			pdf.Text(20, 62, "iccid: "+iccid1)

			// Add page number footer text
			pdf.SetFont("Arial", "I", 8)
			pdf.Text(90, 295, fmt.Sprintf("Page %d of %d", pdf.PageNo(), pagelength))
		}

		// Increment the progress bar
		bar.Increment()

	}

	// Finish the progress bar
	bar.Finish()

	// Save the PDF document
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outputPath := filepath.Join(dir, "output.pdf")
	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nQR codes generated and saved to %s\n", outputPath)

	// calculate time TEMP
	elapsed := time.Since(start)
	fmt.Printf("Total Time Taken: %s\n", elapsed)

	// Wait for user input before exiting..this is just so that the app may not crash before displaying the Total time taken and pdf save location
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

// Helper function to format a number with commas
func formatNumberWithCommas(num int) string {
	numberString := strconv.Itoa(num)
	n := len(numberString)
	if n <= 3 {
		return numberString
	}
	var result []byte
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, numberString[i])
	}
	return string(result)
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
