//maps é tipo um structs é como um dicionário em python

package main
import "fmt"
func main(){
	usuario:=map[string]string{
		"nome":"joao",
		"sobrenome":"nogueira",
	}
	usuario2:=map[string]map[string]string{
		"nome":{
			"primeiro":"joao",
			"ultimo":"pedro",
		},
		"curso":{
			"curso":"SI",
			"campus":"IV",

		},
		 

	}
	fmt.Println(usuario)
	//deletar chave
	delete(usuario,"nome")
	fmt.Println(usuario2)
	//deletou a chave "nome"
	fmt.Println(usuario)
	//pra adicionar
	usuario2["Cidade"]=map[string]string{
		"Rio Tinto":"Campus IV",
	}
	fmt.Println(usuario2)


}



