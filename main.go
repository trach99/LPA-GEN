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

	// Generate QR codes and add to the PDF
	for i := 1; i < len(records); i += 10 {
		// Get the first QR code and text pair
		record1 := records[i]
		text1 := record1[findColumnIndex(records[0], "Lpa")]
		iccid1 := record1[findColumnIndex(records[0], "ICCID")]

		// Generate QR code image as a byte slice
		qrCode1, err := qr.Encode(text1, qr.Medium, 200)
		if err != nil {
			log.Fatal(err)
		}

		// Create an io.Reader from the QR code image byte slice
		qrCodeReader1 := bytes.NewReader(qrCode1)

		// Get the second QR code and text pair
		if i+1 < len(records) {
			record2 := records[i+1]
			text2 := record2[findColumnIndex(records[0], "Lpa")]
			iccid2 := record2[findColumnIndex(records[0], "ICCID")]

			// Generate QR code image as a byte slice
			qrCode2, err := qr.Encode(text2, qr.Medium, 200)
			if err != nil {
				log.Fatal(err)
			}

			// Create an io.Reader from the QR code image byte slice
			qrCodeReader2 := bytes.NewReader(qrCode2)

			// Get the third QR code and text pair
			if i+2 < len(records) {
				record3 := records[i+2]
				text3 := record3[findColumnIndex(records[0], "Lpa")]
				iccid3 := record3[findColumnIndex(records[0], "ICCID")]

				// Generate QR code image as a byte slice
				qrCode3, err := qr.Encode(text3, qr.Medium, 200)
				if err != nil {
					log.Fatal(err)
				}

				// Create an io.Reader from the QR code image byte slice
				qrCodeReader3 := bytes.NewReader(qrCode3)

				// Get the fourth QR code and text pair
				if i+3 < len(records) {
					record4 := records[i+3]
					text4 := record4[findColumnIndex(records[0], "Lpa")]
					iccid4 := record4[findColumnIndex(records[0], "ICCID")]

					// Generate QR code image as a byte slice
					qrCode4, err := qr.Encode(text4, qr.Medium, 200)
					if err != nil {
						log.Fatal(err)
					}

					// Create an io.Reader from the QR code image byte slice
					qrCodeReader4 := bytes.NewReader(qrCode4)

					// Get the fifth QR code and text pair
					if i+4 < len(records) {
						record5 := records[i+4]
						text5 := record5[findColumnIndex(records[0], "Lpa")]
						iccid5 := record5[findColumnIndex(records[0], "ICCID")]

						// Generate QR code image as a byte slice
						qrCode5, err := qr.Encode(text5, qr.Medium, 200)
						if err != nil {
							log.Fatal(err)
						}
						// Create an io.Reader from the QR code image byte slice
						qrCodeReader5 := bytes.NewReader(qrCode5)

						// Get the sixth QR code and text pair
						if i+5 < len(records) {
							record6 := records[i+5]
							text6 := record6[findColumnIndex(records[0], "Lpa")]
							iccid6 := record6[findColumnIndex(records[0], "ICCID")]

							// Generate QR code image as a byte slice
							qrCode6, err := qr.Encode(text6, qr.Medium, 200)
							if err != nil {
								log.Fatal(err)
							}
							// Create an io.Reader from the QR code image byte slice
							qrCodeReader6 := bytes.NewReader(qrCode6)

							// Get the seventh QR code and text pair
							if i+6 < len(records) {
								record7 := records[i+6]
								text7 := record7[findColumnIndex(records[0], "Lpa")]
								iccid7 := record7[findColumnIndex(records[0], "ICCID")]

								// Generate QR code image as a byte slice
								qrCode7, err := qr.Encode(text7, qr.Medium, 200)
								if err != nil {
									log.Fatal(err)
								}
								// Create an io.Reader from the QR code image byte slice
								qrCodeReader7 := bytes.NewReader(qrCode7)

								// Get the eighth QR code and text pair
								if i+7 < len(records) {
									record8 := records[i+7]
									text8 := record8[findColumnIndex(records[0], "Lpa")]
									iccid8 := record8[findColumnIndex(records[0], "ICCID")]

									// Generate QR code image as a byte slice
									qrCode8, err := qr.Encode(text8, qr.Medium, 200)
									if err != nil {
										log.Fatal(err)
									}
									// Create an io.Reader from the QR code image byte slice
									qrCodeReader8 := bytes.NewReader(qrCode8)

									// Get the ninth QR code and text pair
									if i+8 < len(records) {
										record9 := records[i+8]
										text9 := record9[findColumnIndex(records[0], "Lpa")]
										iccid9 := record9[findColumnIndex(records[0], "ICCID")]

										// Generate QR code image as a byte slice
										qrCode9, err := qr.Encode(text9, qr.Medium, 200)
										if err != nil {
											log.Fatal(err)
										}
										// Create an io.Reader from the QR code image byte slice
										qrCodeReader9 := bytes.NewReader(qrCode9)

										// Get the tenth QR code and text pair
										if i+9 < len(records) {
											record10 := records[i+9]
											text10 := record10[findColumnIndex(records[0], "Lpa")]
											iccid10 := record10[findColumnIndex(records[0], "ICCID")]

											// Generate QR code image as a byte slice
											qrCode10, err := qr.Encode(text10, qr.Medium, 200)
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
											pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the first 'Lpa' and 'ICCID' strings below the first QR code
											pdf.SetFont("Arial", "", 9)
											pdf.Text(10, 75, text1)
											pdf.Text(10, 82, "ICCID: "+iccid1)

											// Add second QR code image to the PDF
											pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
											pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the second 'Lpa' and 'ICCID' strings below the second QR code
											pdf.Text(110, 75, text2)
											pdf.Text(110, 82, "ICCID: "+iccid2)

											// Add third QR code image to the PDF
											pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
											pdf.ImageOptions(text3, 20, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the third 'Lpa' and 'ICCID' strings below the third QR code
											pdf.Text(10, 151, text3)
											pdf.Text(10, 158, "ICCID: "+iccid3)
											// Add fourth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
											pdf.ImageOptions(text4, 120, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the fourth 'Lpa' and 'ICCID' strings below the fourth QR code
											pdf.Text(110, 151, text4)
											pdf.Text(110, 158, "ICCID: "+iccid4)

											// Add fifth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
											pdf.ImageOptions(text5, 20, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the fifth 'Lpa' and 'ICCID' strings below the fifth QR code
											pdf.Text(10, 227, text5)
											pdf.Text(10, 234, "ICCID: "+iccid5)

											// Add sixth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
											pdf.ImageOptions(text6, 120, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the sixth 'Lpa' and 'ICCID' strings below the sixth QR code
											pdf.Text(110, 227, text6)
											pdf.Text(110, 234, "ICCID: "+iccid6)

											// Add seventh QR code image to the PDF
											pdf.RegisterImageOptionsReader(text7, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader7)
											pdf.ImageOptions(text7, 20, 242, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the seventh 'Lpa' and 'ICCID' strings below the seventh QR code
											pdf.Text(10, 303, text7)
											pdf.Text(10, 310, "ICCID: "+iccid7)

											// Add eighth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text8, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader8)
											pdf.ImageOptions(text8, 120, 242, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the eighth 'Lpa' and 'ICCID' strings below the eighth QR code
											pdf.Text(110, 303, text8)
											pdf.Text(110, 310, "ICCID: "+iccid8)

											// Add ninth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text9, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader9)
											pdf.ImageOptions(text9, 20, 318, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the ninth 'Lpa' and 'ICCID' strings below the ninth QR code
											pdf.Text(10, 379, text9)
											pdf.Text(10, 386, "ICCID: "+iccid9)

											// Add tenth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text10, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader10)
											pdf.ImageOptions(text10, 120, 318, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the tenth 'Lpa' and 'ICCID' strings below the tenth QR code
											pdf.Text(110, 379, text10)
											pdf.Text(110, 386, "ICCID: "+iccid10)

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
											pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the first 'Lpa' and 'ICCID' strings below the first QR code
											pdf.SetFont("Arial", "", 9)
											pdf.Text(10, 75, text1)
											pdf.Text(10, 82, "ICCID: "+iccid1)

											// Add second QR code image to the PDF
											pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
											pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the second 'Lpa' and 'ICCID' strings below the second QR code
											pdf.Text(110, 75, text2)
											pdf.Text(110, 82, "ICCID: "+iccid2)

											// Add third QR code image to the PDF
											pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
											pdf.ImageOptions(text3, 20, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the third 'Lpa' and 'ICCID' strings below the third QR code
											pdf.Text(10, 151, text3)
											pdf.Text(10, 158, "ICCID: "+iccid3)

											// Add fourth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
											pdf.ImageOptions(text4, 120, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the fourth 'Lpa' and 'ICCID' strings below the fourth QR code
											pdf.Text(110, 151, text4)
											pdf.Text(110, 158, "ICCID: "+iccid4)

											// Add fifth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
											pdf.ImageOptions(text5, 20, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the fifth 'Lpa' and 'ICCID' strings below the fifth QR code
											pdf.Text(10, 227, text5)
											pdf.Text(10, 234, "ICCID: "+iccid5)

											// Add sixth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
											pdf.ImageOptions(text6, 120, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the sixth 'Lpa' and 'ICCID' strings below the sixth QR code
											pdf.Text(110, 227, text6)
											pdf.Text(110, 234, "ICCID: "+iccid6)

											// Add seventh QR code image to the PDF
											pdf.RegisterImageOptionsReader(text7, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader7)
											pdf.ImageOptions(text7, 20, 242, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the seventh 'Lpa' and 'ICCID' strings below the seventh QR code
											pdf.Text(10, 303, text7)
											pdf.Text(10, 310, "ICCID: "+iccid7)

											// Print the eighth 'Lpa' and 'ICCID' strings below the eighth QR code
											pdf.Text(110, 303, text8)
											pdf.Text(110, 310, "ICCID: "+iccid8)

											// Add ninth QR code image to the PDF
											pdf.RegisterImageOptionsReader(text9, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader9)
											pdf.ImageOptions(text9, 20, 318, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

											// Print the ninth 'Lpa' and 'ICCID' strings below the ninth QR code
											pdf.Text(10, 379, text9)
											pdf.Text(10, 386, "ICCID: "+iccid9)

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
										pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the first 'Lpa' and 'ICCID' strings below the first QR code
										pdf.SetFont("Arial", "", 9)
										pdf.Text(10, 75, text1)
										pdf.Text(10, 82, "ICCID: "+iccid1)

										// Add second QR code image to the PDF
										pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
										pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the second 'Lpa' and 'ICCID' strings below the second QR code
										pdf.Text(110, 75, text2)
										pdf.Text(110, 82, "ICCID: "+iccid2)

										// Add third QR code image to the PDF
										pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
										pdf.ImageOptions(text3, 20, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the third 'Lpa' and 'ICCID' strings below the third QR code
										pdf.Text(10, 151, text3)
										pdf.Text(10, 158, "ICCID: "+iccid3)

										// Add fourth QR code image to the PDF
										pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
										pdf.ImageOptions(text4, 120, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the fourth 'Lpa' and 'ICCID' strings below the fourth QR code
										pdf.Text(110, 151, text4)
										pdf.Text(110, 158, "ICCID: "+iccid4)

										// Add fifth QR code image to the PDF
										pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
										pdf.ImageOptions(text5, 20, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the fifth 'Lpa' and 'ICCID' strings below the fifth QR code
										pdf.Text(10, 227, text5)
										pdf.Text(10, 234, "ICCID: "+iccid5)

										// Add sixth QR code image to the PDF
										pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
										pdf.ImageOptions(text6, 120, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the sixth 'Lpa' and 'ICCID' strings below the sixth QR code
										pdf.Text(110, 227, text6)
										pdf.Text(110, 234, "ICCID: "+iccid6)

										// Add seventh QR code image to the PDF
										pdf.RegisterImageOptionsReader(text7, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader7)
										pdf.ImageOptions(text7, 20, 242, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

										// Print the seventh 'Lpa' and 'ICCID' strings below the seventh QR code
										pdf.Text(10, 303, text7)
										pdf.Text(10, 310, "ICCID: "+iccid7)

										// Print the eighth 'Lpa' and 'ICCID' strings below the eighth QR code
										pdf.Text(110, 303, text8)
										pdf.Text(110, 310, "ICCID: "+iccid8)
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
									pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the first 'Lpa' and 'ICCID' strings below the first QR code
									pdf.SetFont("Arial", "", 9)
									pdf.Text(10, 75, text1)
									pdf.Text(10, 82, "ICCID: "+iccid1)

									// Add second QR code image to the PDF
									pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
									pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the second 'Lpa' and 'ICCID' strings below the second QR code
									pdf.Text(110, 75, text2)
									pdf.Text(110, 82, "ICCID: "+iccid2)

									// Add third QR code image to the PDF
									pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
									pdf.ImageOptions(text3, 20, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the third 'Lpa' and 'ICCID' strings below the third QR code
									pdf.Text(10, 151, text3)
									pdf.Text(10, 158, "ICCID: "+iccid3)

									// Add fourth QR code image to the PDF
									pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
									pdf.ImageOptions(text4, 120, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the fourth 'Lpa' and 'ICCID' strings below the fourth QR code
									pdf.Text(110, 151, text4)
									pdf.Text(110, 158, "ICCID: "+iccid4)

									// Add fifth QR code image to the PDF
									pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
									pdf.ImageOptions(text5, 20, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the fifth 'Lpa' and 'ICCID' strings below the fifth QR code
									pdf.Text(10, 227, text5)
									pdf.Text(10, 234, "ICCID: "+iccid5)

									// Add sixth QR code image to the PDF
									pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
									pdf.ImageOptions(text6, 120, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the sixth 'Lpa' and 'ICCID' strings below the sixth QR code
									pdf.Text(110, 227, text6)
									pdf.Text(110, 234, "ICCID: "+iccid6)

									// Add seventh QR code image to the PDF
									pdf.RegisterImageOptionsReader(text7, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader7)
									pdf.ImageOptions(text7, 20, 242, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

									// Print the seventh 'Lpa' and 'ICCID' strings below the seventh QR code
									pdf.Text(10, 303, text7)
									pdf.Text(10, 310, "ICCID: "+iccid7)
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
								pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the first 'Lpa' and 'ICCID' strings below the first QR code
								pdf.SetFont("Arial", "", 9)
								pdf.Text(10, 75, text1)
								pdf.Text(10, 82, "ICCID: "+iccid1)

								// Add second QR code image to the PDF
								pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
								pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the second 'Lpa' and 'ICCID' strings below the second QR code
								pdf.Text(110, 75, text2)
								pdf.Text(110, 82, "ICCID: "+iccid2)

								// Add third QR code image to the PDF
								pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
								pdf.ImageOptions(text3, 20, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the third 'Lpa' and 'ICCID' strings below the third QR code
								pdf.Text(10, 151, text3)
								pdf.Text(10, 158, "ICCID: "+iccid3)

								// Add fourth QR code image to the PDF
								pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
								pdf.ImageOptions(text4, 120, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the fourth 'Lpa' and 'ICCID' strings below the fourth QR code
								pdf.Text(110, 151, text4)
								pdf.Text(110, 158, "ICCID: "+iccid4)

								// Add fifth QR code image to the PDF
								pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
								pdf.ImageOptions(text5, 20, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the fifth 'Lpa' and 'ICCID' strings below the fifth QR code
								pdf.Text(10, 227, text5)
								pdf.Text(10, 234, "ICCID: "+iccid5)

								// Add sixth QR code image to the PDF
								pdf.RegisterImageOptionsReader(text6, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader6)
								pdf.ImageOptions(text6, 120, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

								// Print the sixth 'Lpa' and 'ICCID' strings below the sixth QR code
								pdf.Text(110, 227, text6)
								pdf.Text(110, 234, "ICCID: "+iccid6)
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
							pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the first 'Lpa' and 'ICCID' strings below the first QR code
							pdf.SetFont("Arial", "", 9)
							pdf.Text(10, 75, text1)
							pdf.Text(10, 82, "ICCID: "+iccid1)

							// Add second QR code image to the PDF
							pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
							pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the second 'Lpa' and 'ICCID' strings below the second QR code
							pdf.Text(110, 75, text2)
							pdf.Text(110, 82, "ICCID: "+iccid2)

							// Add third QR code image to the PDF
							pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
							pdf.ImageOptions(text3, 20, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the third 'Lpa' and 'ICCID' strings below the third QR code
							pdf.Text(10, 151, text3)
							pdf.Text(10, 158, "ICCID: "+iccid3)

							// Add fourth QR code image to the PDF
							pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
							pdf.ImageOptions(text4, 120, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the fourth 'Lpa' and 'ICCID' strings below the fourth QR code
							pdf.Text(110, 151, text4)
							pdf.Text(110, 158, "ICCID: "+iccid4)

							// Add fifth QR code image to the PDF
							pdf.RegisterImageOptionsReader(text5, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader5)
							pdf.ImageOptions(text5, 20, 166, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

							// Print the fifth 'Lpa' and 'ICCID' strings below the fifth QR code
							pdf.Text(10, 227, text5)
							pdf.Text(10, 234, "ICCID: "+iccid5)
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
						pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Print the first 'Lpa' and 'ICCID' strings below the first QR code
						pdf.SetFont("Arial", "", 9)
						pdf.Text(10, 75, text1)
						pdf.Text(10, 82, "ICCID: "+iccid1)

						// Add second QR code image to the PDF
						pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
						pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Print the second 'Lpa' and 'ICCID' strings below the second QR code
						pdf.Text(110, 75, text2)
						pdf.Text(110, 82, "ICCID: "+iccid2)

						// Add third QR code image to the PDF
						pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
						pdf.ImageOptions(text3, 20, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Print the third 'Lpa' and 'ICCID' strings below the third QR code
						pdf.Text(10, 151, text3)
						pdf.Text(10, 158, "ICCID: "+iccid3)

						// Add fourth QR code image to the PDF
						pdf.RegisterImageOptionsReader(text4, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader4)
						pdf.ImageOptions(text4, 120, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

						// Print the fourth 'Lpa' and 'ICCID' strings below the fourth QR code
						pdf.Text(110, 151, text4)
						pdf.Text(110, 158, "ICCID: "+iccid4)

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
					pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

					// Print the first 'Lpa' and 'ICCID' strings below the first QR code
					pdf.SetFont("Arial", "", 9)
					pdf.Text(10, 75, text1)
					pdf.Text(10, 82, "ICCID: "+iccid1)

					// Add second QR code image to the PDF
					pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
					pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

					// Print the second 'Lpa' and 'ICCID' strings below the second QR code
					pdf.Text(110, 75, text2)
					pdf.Text(110, 82, "ICCID: "+iccid2)

					// Add third QR code image to the PDF
					pdf.RegisterImageOptionsReader(text3, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader3)
					pdf.ImageOptions(text3, 20, 90, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

					// Print the third 'Lpa' and 'ICCID' strings below the third QR code
					pdf.Text(10, 151, text3)
					pdf.Text(10, 158, "ICCID: "+iccid3)
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
				pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

				// Print the first 'Lpa' and 'ICCID' strings below the first QR code
				pdf.SetFont("Arial", "", 9)
				pdf.Text(10, 75, text1)
				pdf.Text(10, 82, "ICCID: "+iccid1)

				// Add second QR code image to the PDF
				pdf.RegisterImageOptionsReader(text2, gofpdf.ImageOptions{ImageType: "png"}, qrCodeReader2)
				pdf.ImageOptions(text2, 120, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

				// Print the second 'Lpa' and 'ICCID' strings below the second QR code
				pdf.Text(110, 75, text2)
				pdf.Text(110, 82, "ICCID: "+iccid2)
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
			pdf.ImageOptions(text1, 20, 14, 0, 0, false, gofpdf.ImageOptions{ImageType: "png"}, 0, "")

			// Print the first 'Lpa' and 'ICCID' strings below the first QR code
			pdf.SetFont("Arial", "", 9)
			pdf.Text(10, 75, text1)
			pdf.Text(10, 82, "ICCID: "+iccid1)
		}
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
