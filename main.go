//primero voy a dar entender que es lo que vamos a hacer en el programa.
//tendra los siguentes atributos o variables titular y cantidad puede tener decimales
//el titular sera obligatorio y la cantidad es opcional. crea dos constructores que cumpla lo anterior
//crea sus metodos get,set y toString

//Tendrá dos metodos especiales:
//*Ingresar(double): se ingresa una cantidad a la cuenta, si la cantidad introducida es negativa,no se hara nada.
//*retirar(double): se retira una cantidad a la cuenta, si restando la cantidad acual a la que nos pasan es negativa.la cantidad de la cuenta pasa a ser 0.
//vamos a realizar una cuenta de banco el cual va a tener el nombre de alex y las personas que ingresen o manden dinero valla aumentando la cuenta con un metodo y si retira dinero y queda negativo debe de convertirse en 0

package main

import "fmt"

var dep, ret, CT float32
var selecion int

func main() {
	fmt.Println("Que desea realizar?")
	fmt.Println("")
	var opcion = [3]string{"Depositar", "Consultar", "Retirar"}
	for contador := 0; contador < len(opcion); contador++ {
		//bloque de codigo
		fmt.Println(contador, "-", opcion[contador])
		fmt.Println("")
	}
	fmt.Scan(&selecion)
	for {
		fmt.Println("")
		Opcion()
		break
	}
}

func Opcion() {
	if selecion == 0 {
		fmt.Println("Depositar")
		fmt.Println("")
		fmt.Println("Cuanto es lo que desea depositar")
		fmt.Println("")
		fmt.Scan(&dep)
		CT = CT + dep
		fmt.Println("Gracias por el deposito de: ", dep)
		fmt.Println("")
	}
	if selecion == 2 {
		fmt.Println("Retirar")
		fmt.Println("")
		fmt.Println("Cuanto es lo que desea retirar")
		fmt.Println("")
		fmt.Scan(&ret)
		if ret > CT {
			fmt.Println("El retiro no se pued realizar")
			fmt.Println("")
		} else {
			fmt.Println("el retiro de ", ret, " a sido exitoso")
			fmt.Println("")
			CT = CT - ret
		}
	}
	if selecion == 1 {
		if CT <= 0 {
			CT = 0
		} else {
			fmt.Println("La cantidad que tiene en un cuenta es de:", CT)
			fmt.Println("")
		}
	}
	fmt.Println("")
	main()
}
