package main
import "fmt"
type pessoa struct{
	nome string
	sobrenome string
	idade uint16
	altura uint16

}
type aluno struct{
	pessoa //herança de structs pessoa para aluno
	curso string
	faculdade string
}
func main(){
	novoaluno:=pessoa{"joao","paiva",21,190}
	fmt.Println(novoaluno)
	novoestudante:=aluno{novoaluno,"SI" ,"UFPB"}
	fmt.Println(novoestudante)
	//acessar um campo especifico
	fmt.Println(novoaluno.sobrenome)
	fmt.Println(novoestudante.altura)
	fmt.Println(novoestudante.curso,novoestudante.faculdade)

}