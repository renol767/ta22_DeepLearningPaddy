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
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func showData(c *fiber.Ctx) error {
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	var imgd []bson.M

	if err := cursor.All(c.Context(), &imgd); err != nil {
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	return c.JSON(fiber.Map{"status": 200, "message": "Success Get Data", "data": imgd})
}

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
	dataToStore := bson.D{
		{"imgUrlData", imgUrl},
		{"timestamp", time.Now().Format(time.RFC3339)},
	}

	ins, errsa := collection.InsertOne(context.Background(), dataToStore)
	if errsa != nil {
		fmt.Println(errsa)
	}
	fmt.Println("Success Insert Data", ins.InsertedID)

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

	// MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo Connect ERROR : ", err)
		os.Exit(1)
	}

	collection = client.Database("paddy").Collection("imageData")

	// Main Function
	fmt.Println("Berjalan di localhost", port)

	// FIBER
	app := fiber.New()

	// Communication using HTTP and Allow Cors
	app.Use(cors.New())

	// Akses Folder
	app.Static("/result", "result-images/hasildeteksi")

	// Get List Deteksi
	app.Get("/imageDetectionData", showData)

	// Upload Image Services Route
	app.Post("/image-services", handleImageServices)

	log.Fatal(app.Listen(port))

}
