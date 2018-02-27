package models

import (
	"encoding/json"
)

// Cart represents the state of a buyer's shopping cart
type Cart struct {
	items map[string]Item
}

// Item represents any item available for sale
type Item struct {
	ID    string
	Name  string
	Price float64
	Qty   int
}

func (c *Cart) init() {
	if c.items == nil {
		c.items = map[string]Item{}
	}
}

// AddItem agrega un articulo al carrito
func (c *Cart) AddItem(i Item) {
	c.init()
	if existingItem, ok := c.items[i.ID]; ok {
		existingItem.Qty++
		c.items[i.ID] = existingItem
	} else {
		i.Qty = 1
		c.items[i.ID] = i
	}
}

// RemoveItem elimina n cantidad de elementos con la identificación del carrito
func (c *Cart) RemoveItem(id string, n int) {
	c.init()
	if existingItem, ok := c.items[id]; ok {
		if existingItem.Qty <= n {
			delete(c.items, id)
		} else {
			existingItem.Qty -= n
			c.items[id] = existingItem
		}
	}

}

// TotalAmount devuelve el monto total del carrito
func (c *Cart) TotalAmount() float64 {
	c.init()
	totalAmount := 0.0
	for _, i := range c.items {
		totalAmount += i.Price * float64(i.Qty)
	}
	return totalAmount
}


// TotalUnits devuelve el número total de unidades en todos los artículos en el carro
func (c *Cart) TotalUnits() int {
	c.init()
	totalUnits := 0
	for _, i := range c.items {
		totalUnits += i.Qty
	}
	return totalUnits
}

// TotalUniqueItems devuelve la cantidad de artículos unicos en el carro
func (c *Cart) TotalUniqueItems() int {
	return len(c.items)
}

//Metodo para convertir una nota a formato JSON
func (c *Cart) ToJSON() (string) {
	b, err := json.Marshal(c.items)
	if err != nil {
		panic(err)
	}
	return string(b)
}

//vaciar el carrito de compras
func (c *Cart)EmptyCart(){
	c.init()
	c.items = make(map[string]Item)
}