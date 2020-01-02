package main

type info struct {
	CompanyName string
	Address     string
	Jobs        []job
}

type job struct {
	Name string
	Link string
}

func main() {
	content, err := request("https://www.ejobs.ro/company/ejobs-group/266860")
	if err != nil {
		panic(err)
	}
	extract(*content)

}
