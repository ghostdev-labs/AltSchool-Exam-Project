package main

import (
	"fmt"
)

//declare a car struct
type car struct {
	make string
	model string
	year int
}

//declare a product struct and promote the car struct
type product struct {
	car []car
	name string
	price int
	quantity int
}

//declare a method for the product struct 
func (p product) showProduct() {
	fmt.Println(p)
}

//declare a method for the product struct to display the status of a product if it is still in stock or not.
func (p product) showProductStatus() {
	if p.quantity > 0 {
		fmt.Println("In Stock")
	} else {
		fmt.Println("Out of Stock")
	}
}

//declare a store struct with the following attributes:
//1. Show the total number of products in the store for saleppl
//2. Show all products in the store
//3. Add a product to the store
//4. Sell a product from the store
//5. Show list of products sold with the total amount of money made
type store struct {
	products []product
	soldProducts []product
	totalProducts int
	totalSoldProducts int
	totalMoneyMade int
}

// //declare a method for the store struct to show all products in the store
// func (s store) showAllProducts() {
// 	for _, p := range s.products {
// 		fmt.Println(p)
// 	}
// }

func (s store) display(){
	fmt.Println("Number of Products: ", s.totalProducts)
	fmt.Println("List of Products: ", s.products)
	fmt.Println("Sold Products + (total amount): ", s.soldProducts, s.totalMoneyMade)
}

//declare a method for the store struct to add a product to the store
func (s store) addProduct(p product) {
	s.products = append(s.products, p)
	s.totalProducts++
}

//declare a method for the store struct to sell a product from the store
func (s store) sellProduct(p product) {
	s.soldProducts = append(s.soldProducts, p)
	s.totalSoldProducts++
	s.totalMoneyMade += p.price
}

//declare a method for the store struct to show all sold products in the store
func (s store) showAllSoldProducts() {
	for _, p := range s.soldProducts {
		fmt.Println(p)
	}
}

//declare a method for the store struct to show the total money made by the store
func (s store) showTotalMoneyMade() {
	fmt.Println(s.totalMoneyMade)
}


func main() {
	c1 := car{
		make: "Ford",
		model: "Mustang",
		year: 2007,
	}
	c2 := car {
		make: "Dodge",
		model: "Charger",
		year: 2018,
	}
	c3 := car{
		make: "Mercedes Benz",
		model: "C300",
		year: 2019,
	}

	p1 := product{
		car: []car{c1},
		price: 10000.00,
		quantity: 1,
	}
	p2 := product{
		car: []car{c2},
		price: 20000.00,
		quantity: 1,
	}
	p3 := product{
		car: []car{c3},
		price: 10000.00,
		quantity: 1,
	}

	store := store{
		totalProducts: 3,
		products: []product{p1,p2,p3},
		soldProducts: []product{p1,p2},
		totalMoneyMade: p1.price + p2.price,
	}

	
	store.display()
	p1.showProductStatus()
	p2.showProductStatus()
	p3.showProductStatus()
}
