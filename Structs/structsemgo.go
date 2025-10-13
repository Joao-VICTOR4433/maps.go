package main

import "fmt"


type usuario struct{
	nome string
	idade uint8
}
type endereco struct{
	logradouro string
	numero uint64
}

type cliente struct{
	nomedocliente string
	cpf uint64
	idadedocliente uint8
	sexo string
}
type valorproduto struct{
	nomedoproduto string
	valor float32
}
func main(){
	fmt.Println(".")
	//se nao passar nada joga o valor vaziou ou zero
	var u usuario
	u.nome="joao"
	u.idade=12
	fmt.Println(u)
	//declarar o structs primeiro jeito

	usuario2:=usuario{"joao",21}
	fmt.Println(usuario2)
	//criar um structs com inferencia mais rapido

	usuario3:=usuario{nome:"davi"}
	fmt.Println(usuario3)
	//se quiser passar so um dado sem passar o resto

	novoendereco:=endereco{"Rua aurora",928}
	fmt.Println(novoendereco)

	novocliente:=cliente{"joao",0344444444,15,"masculino"}
	fmt.Println(novocliente)

	novoproduto:=valorproduto{"Prato de cristal",39.90}
	fmt.Println(novoproduto)


}
