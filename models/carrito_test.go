package models

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"


)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Shopping cart", func() {

	itemA := Item{ID: "NX", Name: "Nintendo Switch", Price: 300.00, Qty: 0}
	itemB := Item{ID: "Dolphin", Name: "Nintendo GameCube", Price: 199.99, Qty: 1}

	Context("Cuando se inicia el carrito", func() {
		cart := Cart{}

		It("tiene 0 elementos unicos", func() {
			Ω(cart.TotalUniqueItems()).Should(BeZero())
		})

		It("tiene 0 elementos", func() {
			Expect(cart.TotalUnits()).Should(BeZero())
		})

		Specify("el total es de 0.00", func() {
			Expect(cart.TotalAmount()).Should(BeZero())
		})
	})

	Context("cuando un nuevo articulo es agregado", func() {
		//se prepara el carrito
		cart := Cart{}
		//cantidad de artículos unicos en el carrito
		originalItemCount := cart.TotalUniqueItems()
		//número total de unidades en todos los artículos en el carro
		originalUnitCount := cart.TotalUnits()
		//total del carrito
		originalAmount := cart.TotalAmount()
		//agregamos al carrito
		cart.AddItem(itemA)

		Context("el carrito de compras ahora", func() {
			It("tiene 1 elemento unico más de lo que tenía antes", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount + 1))
			})

			It("tiene 1 unidad más de lo que tenía antes", func() {
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount + 1))
			})

			Specify("el monto total aumenta por precio del artículo", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount + itemA.Price))
			})
		})
	})

	Context("cuando se agrega un elemento existente", func() {
		//se prepara el carrito
		cart := Cart{}
		// se agrega al carrito
		cart.AddItem(itemA)
		//cantidad de artículos unicos en el carrito
		originalItemCount := cart.TotalUniqueItems()
		//número total de unidades en todos los artículos en el carro
		originalUnitCount := cart.TotalUnits()
		//total del carrito
		originalAmount := cart.TotalAmount()
		//agregamos al carrito otro producto igual
		cart.AddItem(itemA)

		Context("El carrito de compras", func() {
			It("tiene la misma cantidad de elementos únicos que antes", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
			})

			It("tiene 1 unidad más de lo que tenía antes", func() {
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount + 1))
			})

			Specify("el monto total aumenta por precio del artículo", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount + itemA.Price))
			})
		})
	})

	Context("se tiene 0 elementos del ItemA", func() {
		cart := Cart{}

		cart.AddItem(itemB) // solo para imitar la existencia de otros artículos
		cart.AddItem(itemB) // solo para imitar la existencia de otros artículos

		originalItemCount := cart.TotalUniqueItems()
		originalUnitCount := cart.TotalUnits()
		originalAmount := cart.TotalAmount()

		Context("quitando elemnto A", func() {
			cart.RemoveItem(itemA.ID, 1)

			It("no debe cambiar la cantidad de elementos", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
			})
			It("no debería cambiar la cantidad de unidades", func() {
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount))
			})
			It("No deberia cambiar la cantidad", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount))
			})
		})
	})

	Context("cuando se tiene 1 item A", func() {
		cart := Cart{}

		cart.AddItem(itemB) // solo para imitar la existencia de otros artículos
		cart.AddItem(itemB) // solo para imitar la existencia de otros artículos

		cart.AddItem(itemA)

		originalItemCount := cart.TotalUniqueItems()
		originalUnitCount := cart.TotalUnits()
		originalAmount := cart.TotalAmount()

		Context("quitando 1 unidad del item A", func() {
			cart.RemoveItem(itemA.ID, 1)

			It("debería reducir la cantidad de elementos en 1", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount - 1))
			})

			It("debería reducir el número de unidades por 1", func() {
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount - 1))
			})

			It("debería reducir la cantidad por precio de artículo", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - itemA.Price))
			})
		})
	})

	Context("se tiene 2 unidades del elemento A", func() {
		Context("eliminando 1 unidad del elemento A", func() {
			cart := Cart{}

			cart.AddItem(itemB) // solo para imitar la existencia de otros artículos
			cart.AddItem(itemB) // solo para imitar la existencia de otros artículos
			//Restablece el carro con 2 unidades del elemento A
			cart.AddItem(itemA)
			cart.AddItem(itemA)

			originalItemCount := cart.TotalUniqueItems()
			originalUnitCount := cart.TotalUnits()
			originalAmount := cart.TotalAmount()

			cart.RemoveItem(itemA.ID, 1)

			It("no debe reducir la cantidad de elementos", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount))
			})

			It("debería reducir el número de unidades por 1", func() {
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount - 1))
			})

			It("debe reducir la cantidad por el precio del artículo", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - itemA.Price))
			})
		})

		Context("eliminando 2 unidades del elemento A", func() {
			cart := Cart{}

			cart.AddItem(itemB) // solo para imitar la existencia de otros artículos
			cart.AddItem(itemB) // solo para imitar la existencia de otros artículos
			//Restablece el carro con 2 unidades del elemento A
			cart.AddItem(itemA)
			cart.AddItem(itemA)

			originalItemCount := cart.TotalUniqueItems()
			originalUnitCount := cart.TotalUnits()
			originalAmount := cart.TotalAmount()

			cart.RemoveItem(itemA.ID, 2)

			It("debería reducir la cantidad de elementos en 1", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(originalItemCount - 1))
			})

			It("debería reducir el número de unidades por 2", func() {
				Expect(cart.TotalUnits()).Should(Equal(originalUnitCount - 2))
			})

			It("debe reducir la cantidad en dos veces el precio del artículo", func() {
				Expect(cart.TotalAmount()).Should(Equal(originalAmount - 2*itemA.Price))
			})
		})

	})

	Context("Agregando 3 elementos diferentes al carrito y en total son 10 productos", func() {
		cart := Cart{}

		cart.AddItem(itemA)
		cart.AddItem(itemA)
		cart.AddItem(itemA)
		cart.AddItem(itemA)

		cart.AddItem(itemB)
		cart.AddItem(itemB)
		cart.AddItem(itemB)
		cart.AddItem(itemB)
		cart.AddItem(itemB)

		cart.AddItem(Item{"3ds", "Nintendo 3ds", 250.00, 1})

		TotalElementosUnicos := cart.TotalUniqueItems()
		TotalElementos := cart.TotalUnits()
		MontoTotal := cart.TotalAmount()

		Context("En el Carrito ... ", func() {
			cart.RemoveItem("3ds", 1)
			It("Eliminamos el ultimo elemento agregado", func() {
				Expect(cart.TotalUniqueItems()).Should(Equal(TotalElementosUnicos - 1))
			})

			It("Comprobamos que existan 9 elementos", func() {
				Expect(cart.TotalUnits()).Should(Equal(TotalElementos - 1))
			})

			It("el monto total se reduce por el precio del articulo", func() {
				Expect(cart.TotalAmount()).Should(Equal(MontoTotal - 250.00))
			})
		})


		Context("En el Carrito ...", func() {
			cart := Cart{}

			cart.AddItem(itemA)
			cart.AddItem(itemA)
			 // === || ======
			cart.AddItem(itemB)
			cart.AddItem(itemB)
			// limpiamos
			cart.EmptyCart()

			Context("Vaciamos todo el carrito", func() {
				It("El total de elementos debe ser igual a 0", func() {
					Expect(cart.TotalUnits()).Should(BeZero())
				})

				It("El total de elementos unicos debe ser 0", func() {
					Expect(cart.TotalUniqueItems()).Should(BeZero())
				})

				It("El Monto total debe ser igual a 0.00", func() {
					Expect(cart.TotalAmount()).Should(BeZero())
				})
			})
		})
	})
})
