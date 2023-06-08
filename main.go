package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	type Product struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		Price        string `json:"price"`
		Stock        int64  `json:"stock"`
		Code         string `json:"code"`
		Published    bool   `json:"published"`
		CreationDate string `json:"creation_date"`
	}

	content, err := ioutil.ReadFile("./products.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var Stock []Product
	err = json.Unmarshal(content, &Stock)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Println(Stock)

	// func openFile(path string) {

	// 	f, err := os.Open(path)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	defer f.Close()

	// 	r := csv.NewReader(f)

	// 	var productList []Product

	// 	for {
	// 		record, err := r.Read()

	// 		if err == io.EOF {

	// 			fmt.Printf("Se ha llegado al final del archivo: %s \n", err)
	// 			break
	// 		}

	// 		id, err := strconv.Atoi(record[0])

	// 		price, err := strconv.ParseFloat(record[5], 64)

	// 		if err != nil {
	// 			return nil, err

	// 		}

	// 		ticketsList = append(ticketsList, Ticket{
	// 			Id:          id,
	// 			Name:        record[1],
	// 			Email:       record[2],
	// 			Destination: record[3],
	// 			Departure:   record[4],
	// 			Price:       price,
	// 		})
	// 	}

	// 	return ticketsList, nil
	// }

	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"products": Stock,
		})
	})

	router.Run()

}
