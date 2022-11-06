package main

import "fmt"

//Car A Car struct
type Car struct {
	make  string
	model string
	year  int
}

//Product A Product struct
type Product struct {
	car      []Car
	name     string
	price    float64
	quantity int
}

//showProduct method to show product details
func (p Product) showProduct() {
	fmt.Println("Product name:", p.name)
	fmt.Println("Product price:", p.price)
	fmt.Println("Product quantity:", p.quantity)
	fmt.Println("Product car:", p.car)
	fmt.Println("--------------------")
}

//showProductStatus method to show product status
func (p Product) showProductStatus() {
	if p.quantity > 0 {
		fmt.Println("Product is available")
	} else {
		fmt.Println("Product is not available")
	}
}

//Store A Store struct
type Store struct {
	products          []Product
	numbOfProducts    int
	totalSoldProducts int
	totalRevenue      float64
}

//addProduct method to add a product to the store
func (s *Store) addProduct(p Product) {
	s.products = append(s.products, p)
	s.numbOfProducts++
}

//sellProduct method to sell a product from the store
func (s *Store) sellProduct(name string, quantity int) {
	for i, p := range s.products {
		if p.name == name {
			if p.quantity >= quantity {
				s.products[i].quantity -= quantity
				s.totalSoldProducts += quantity
				s.totalRevenue += p.price * float64(quantity)
			} else {
				fmt.Println("Not enough quantity")
			}
		}
	}
}

//showNumOfProducts method to show the number of products in the store
func (s *Store) showNumOfProducts() {
	fmt.Println("Total products:", s.numbOfProducts)
	fmt.Println("--------------------")
}

//showProducts method to show to the products in the store
func (s *Store) showProducts() {
	for _, p := range s.products {
		p.showProduct()
	}
}

//showProductStatus method to show the status of the store
func (s *Store) showProductStatus() {
	for _, p := range s.products {
		p.showProductStatus()
		fmt.Println("--------------------")
	}
}

//showTotalSoldProducts method to show the total number of sold products
func (s *Store) showTotalSoldProducts() {
	fmt.Println("Total sold products:", s.totalSoldProducts)
}

//showTotalRevenue method to show total revenue made from product sales
func (s *Store) showTotalRevenue() {
	fmt.Println("Total revenue:", s.totalRevenue)
}

func main() {
	//Car struct
	car1 := Car{"Toyota", "Corolla", 2019}
	car2 := Car{"Honda", "Civic", 2018}
	car3 := Car{"Ford", "Mustang", 2017}

	//Product struct
	product1 := Product{[]Car{car1, car2}, "Car", 10000, 10}
	product2 := Product{[]Car{car3}, "Car", 20000, 5}

	//Store struct
	store := Store{}
	store.addProduct(product1)
	store.addProduct(product2)

	store.showNumOfProducts()

	store.showProducts()

	store.showProductStatus()

	store.sellProduct("Car", 5)

	store.showProducts()

	store.showProductStatus()

	store.showTotalSoldProducts()

	store.showTotalRevenue()
}
