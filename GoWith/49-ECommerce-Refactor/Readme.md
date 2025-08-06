- `go mod init ecommerce`
- Run the file: ` go run main.go`


## To see the fronted

-  cd product-listing-react
-  npm i
-  npm start


## Refactor the codebase GET,POST:


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

	handleCors(w)
	handlePreflightReq(w,r)

	if r.Method != http.MethodGet {
		http.Error(w, "please provide get request", http.StatusBadRequest)
		return
	}

	sendData(w,productList,http.StatusOK)

}

func createProduct(w http.ResponseWriter, r *http.Request) {

	handleCors(w)
	handlePreflightReq(w,r)

	if r.Method != http.MethodPost {
		http.Error(w, "Plz provide Post request", http.StatusBadRequest)
		return
	}

	var newProductData Product

	dcoder := json.NewDecoder(r.Body)
	err := dcoder.Decode(&newProductData)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}

	newProductData.ID = len(productList) + 1
	productList = append(productList, newProductData)

	sendData(w,newProductData,http.StatusCreated)

}
//cors
func handleCors(w http.ResponseWriter)  {
	//Handle the cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Bappa")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

}

func handlePreflightReq(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusCreated);
	}
}

// data interface{} -> generic type
func sendData(w http.ResponseWriter, data interface{},statusCode int)  {

	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

// main
func main() {
	mux := http.NewServeMux() //this is router

	//Routes
	mux.HandleFunc("/hello", hellohandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-products", createProduct)

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