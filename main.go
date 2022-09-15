package main

import 
(
	"bufio"
//	"strconv"
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"openlife/models"
)

func main() {

	//loading the temporary test data
	file, err := os.Open("./raw-data/wfo-json.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//initialize data container for the time being
	var dataset [20]models.WFOPlant

	//this here can fail if there exist lines with more than around 65k characters. then we'd need to adjust the buffer size manually.
	scanner := bufio.NewScanner(file)

	var line = 0;
	for scanner.Scan() {
		line++
		if line > 10000 && line <= 10020 {
			//fmt.Printf(scanner.Text())
			err := json.Unmarshal([]byte(scanner.Text()), &dataset[line - 10001])
			fmt.Printf("Unmarshaled: " + dataset[line - 10001].TaxonID + "\n")

			if err != nil {
				log.Fatal(err)
			}
		}
	}



	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")

	//using gin to handle request
	router := gin.Default()

	router.GET(
		"/plants",
		func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, dataset)
		})

	router.GET(
		"/helloworld",
		func(c *gin.Context) {
			fmt.Fprintf(os.Stdout, "Hello, World! from gin")
		})

	router.Run("localhost:8080")

	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}
}
