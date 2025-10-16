package main

import "fmt"

func main() {
	//revisao slice e arrayinterno
	slice := []int{10, 11, 12, 13, 14, 15}
	slice = append(slice, 18)
	fmt.Println(slice)
	
	//arrayinternos
	fmt.Println("x----------------------x--------------x------x")
	slice2:=make([]int8,10,11)
	//make([]tipo do slice,tamanho dele,o maximo que aguenta)
	fmt.Println(slice2)
	fmt.Println("O tamanho do slice é: ",len(slice2))//tamanho
	fmt.Println("A capacidade maxima é :",cap(slice2))//capacidade maxima
	slice2 = append(slice2,5)
	slice2=append(slice2,4)
	slice2=append(slice2, 9)
	fmt.Println(slice2)
	//Quando voce estora o limite maximo ele cria um novo slice com o dobro do tamanho
	fmt.Println("O tamanho do slice é: ",len(slice2))//tamanho
	fmt.Println("A capacidade maxima é :",cap(slice2))//capacidade maxima
	slice3:=make([]int8,8)
	//quando voce nao  passa o utilmo parametro ele entende que o limite é igual o numero de elementos
	fmt.Println(slice3)



}