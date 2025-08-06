- `go mod init ecommerce`
- Run the file: ` go run main.go`


## To see the fronted

-  cd product-listing-react
-  cd client
-  npm i
-  npm start


## Handle the POST Request:

```Go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// if r.Method !="GET"{}
	if r.Method != http.MethodGet {
		http.Error(w, "please provide get request", http.StatusBadRequest)
		return
	}

	//if no error then
	//the response send using encoder
	encoder := json.NewEncoder(w)
	encoder.Encode(productList) // send fronted to json format

}

func createProduct(w http.ResponseWriter, r *http.Request) {

	//Handle the cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	//handle options; so that no error comes 
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusCreated);
		return;
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Plz provide Post request", http.StatusBadRequest)
		return
	}

	// fmt.Println(r.Body)
	//1. take the data{title, description, price,imagUrl} from  request.Body
	//2. Create an instance of Struct with the body information
	//3. Append the instance into product list array / slice

	var newProductData Product

	dcoder := json.NewDecoder(r.Body)
	err := dcoder.Decode(&newProductData)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide valid json", http.StatusBadRequest)
		return
	}

	//set new id . As we can not get id from fronted from
	newProductData.ID = len(productList) + 1
	productList = append(productList, newProductData)
	w.WriteHeader(http.StatusCreated)

	//response send using encoder
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

// this handle form data from postman
func createProduct2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Please provide POST request", http.StatusBadRequest)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Get values from form
	title := r.FormValue("title")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	imgUrl := r.FormValue("imageUrl")

	// Convert price string to float64
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	// Create the product
	newProductData := Product{
		ID:          len(productList) + 1,
		Title:       title,
		Description: description,
		Price:       price,
		ImgUrl:      imgUrl,
	}

	// Append to the product list
	productList = append(productList, newProductData)

	// Return the updated product list as response
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

// main
func main() {
	mux := http.NewServeMux() //this is router

	//Routes
	mux.HandleFunc("/hello", hellohandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-products", createProduct)
	mux.HandleFunc("/create-product2", createProduct2) //handle formdata from postman

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