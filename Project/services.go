/*
	DONE :
	Redirect Ketika sudah Upload Gambar
	Tampilkan Hasil Deteksi
	Deteksi Gambar dan menghasilkan Lokasi File Hasil Deteksi

	NOTE:
	YOLOv5s
	command python
	python detect.py --source ../go/uploaded-images/107u.jpg --weights ../best.pt --img 500 --project ../go/result-images
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SERVICES
func handleImageServices(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("Image Upload Error", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	filename := file.Filename
	image := fmt.Sprintf("%s", filename)

	err = c.SaveFile(file, fmt.Sprintf("./uploaded-images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// Deteksi Gambar Menggunakan YOLOv5
	cmm := "python ../yolov5/detect.py --source uploaded-images/" + file.Filename + " --weights ../best.pt --img 500 --project ../go/result-images --name hasildeteksi --exist-ok"
	// python ../yolov5/detect.py --source uploaded-images/107u.jpg --weights ../best.pt --img 500 --project ../go/result-images --name hasildeteksi --exist-ok
	cmd := exec.Command("bash", "-c", cmm)
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cmm)

	imgUrl := fmt.Sprintf("http://localhost:9000/result/%s", image)

	// Create meta data for response
	data := map[string]interface{}{
		"imageUrl": imgUrl,
		"header":   file.Header,
		"size":     file.Size,
	}

	// Return JSON
	return c.JSON(fiber.Map{"status": 200, "message": "Image Upload Success", "data": data})
}

var port = ":9000"

func main() {
	// Main Function
	fmt.Println("Berjalan di localhost", port)

	// FIBER
	app := fiber.New()

	// Communication using HTTP and Allow Cors
	app.Use(cors.New())

	// Akses Folder
	app.Static("/result", "result-images/hasildeteksi")

	// Upload Image Services Route
	app.Post("/image-services", handleImageServices)

	log.Fatal(app.Listen(port))

}
