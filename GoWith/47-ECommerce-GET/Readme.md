- `go mod init ecommerce`
- Run the file: ` go run main.go`


## Handle the Get Request:



```Go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struct means object
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product // products er array

func hellohandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "This is hello")
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Orgin","*");
	w.Header().Set("Content-Type","application/json");
	// if r.Method !="GET"{}
	if r.Method != http.MethodGet {
		http.Error(w, "please provide get request", http.StatusBadRequest)
		return
	}

	//if no error then
	encoder := json.NewEncoder(w)
	encoder.Encode(productList) // send fronted to json format

}

func main() {
	mux := http.NewServeMux() //this is router

	mux.HandleFunc("/hello", hellohandler)
	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server is running on port: 1234")

	err := http.ListenAndServe(":1234", mux)
	// if err == nill means no error // if err != nill means error

	//means error exist
	if err != nil {
		fmt.Println("Error prevents to start the server", err)
	}
}

// start first
func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Apple",
		Description: "This is an apple",
		Price:       10,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Red_Apple.jpg/1130px-Red_Apple.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Black",
		Description: "This is an bapple",
		Price:       20,
		ImgUrl:      "https://dk2dv4ezy246u.cloudfront.net/widgets/sSpmjChEZud_large.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Green",
		Description: "This is an capple",
		Price:       30,
		ImgUrl:      "https://static.vecteezy.com/system/resources/thumbnails/012/086/172/small_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "yellow",
		Description: "This is an dapple",
		Price:       40,
		ImgUrl:      "https://t4.ftcdn.net/jpg/07/21/27/41/360_F_721274171_y29SQKuTPhLWn1ovrRG1SqNfY8rs8obD.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)

}




```