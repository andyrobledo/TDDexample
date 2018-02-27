package main

import (
	"fmt"
	"github.com/andyrobledo/TDDexample/models"
)

func main(){
	fmt.Println("\n-- Iniciando Carrito de compras --")
	//iniciamos carrito
	carrito := models.Cart{}

	// capturamos algunos productos
	Nswitch := models.Item{"NX", "Nintendo Switch", 300.00, 0}
	Ncubo := models.Item{"Dolphin", "Nintendo GameCube", 80.00, 0}
	Nwii := models.Item{"Revo", "Nintendo Wii", 100.00, 0}
	NwiiU := models.Item{"Cemu", "Nintendo Wii U", 200.00, 2}

	// tambien se puede asi, es mas claro
	N3ds := models.Item{}
	N3ds.ID = "3ds"
	N3ds.Name = "Nintendo 3DS"
	N3ds.Price = 150.00
	N3ds.Qty = 0

	// los agregamos al carrito de compras
	fmt.Println("\nAgregamos Productos al carrito")
	carrito.AddItem(Nswitch)
	carrito.AddItem(Ncubo)
	carrito.AddItem(Nwii)
	carrito.AddItem(NwiiU)
	carrito.AddItem(NwiiU)
	carrito.AddItem(N3ds)

	fmt.Println("\nTotal de articulos: ", carrito.TotalUnits())
	fmt.Println("Total de articulos diferentes: ", carrito.TotalUniqueItems())
	fmt.Println("Total: $", carrito.TotalAmount())

	fmt.Println("\nAhora eliminamos al nintendo switch: ")
	carrito.RemoveItem(Nswitch.ID, 1)
	fmt.Println("\nTotal de articulos: ", carrito.TotalUnits())
	fmt.Println("Total de Unicos: ", carrito.TotalUniqueItems())
	fmt.Println("Total: $", carrito.TotalAmount())

	fmt.Println("\nEliminamos dos nintendos wii u")
	carrito.RemoveItem(NwiiU.ID, 2)
	fmt.Println("\nTotal de articulos: ", carrito.TotalUnits())
	fmt.Println("Total de Unicos: ", carrito.TotalUniqueItems())
	fmt.Println("Total: $", carrito.TotalAmount())


	//fmt.Println("\n==Vaciar Carrito, No queremos nada")
	//carrito.EmptyCart()
	//fmt.Println("\nTotal de articulos: ", carrito.TotalUnits())

	//captura elemento
	newItem := models.Item{}
	fmt.Println("Ingrese ID")
	fmt.Scanf("%s", &newItem.ID)
	fmt.Println("Ingrese Nombre:")
	fmt.Scanf("%s", &newItem.Name)
	fmt.Println("Ingrese Precio")
	fmt.Scanf("%f", &newItem.Price)

	carrito.AddItem(newItem)


	fmt.Println("==== JSON ====")
	fmt.Println(carrito.ToJSON())
}
