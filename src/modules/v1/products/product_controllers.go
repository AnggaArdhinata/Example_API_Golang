package products

import (
	"encoding/json"
	"net/http"
	"restapiexample/src/databases/models"
	"restapiexample/src/interfaces"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	ProductService interfaces.ProductService
}

func NewController(productService interfaces.ProductService) *ProductController {
	return &ProductController{productService}
}

func (c *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := c.ProductService.GetAll()
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *ProductController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idToInt, _ := strconv.Atoi(id)
	idToUInt := uint(idToInt)

	response, err := c.ProductService.GetByID(idToUInt)

	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *ProductController) GetByName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	response, err := c.ProductService.GetByName(name)

	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *ProductController) Add(w http.ResponseWriter, r *http.Request) {
	var data models.Product

	json.NewDecoder(r.Body).Decode(&data)
	response, err := c.ProductService.Add(&data)
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product
	idToInt, _ := strconv.Atoi(id)
	productID := uint(idToInt)
	product.ID = uint(idToInt)

	json.NewDecoder(r.Body).Decode(&product)
	response, err := c.ProductService.Update(productID, &product)
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}

func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product
	idToInt, _ := strconv.Atoi(id)
	product.ID = uint(idToInt)

	json.NewDecoder(r.Body).Decode(&product)
	response, err := c.ProductService.Delete(&product)
	if err != nil {
		response.Send(w)
	} else {
		response.Send(w)
	}
}
