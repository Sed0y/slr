package solaris

import (
	"fmt"
)

const (
	api_GetSchema = "/business_api/get_schema/"
	api_GetPeople = "/business_api/get_peoples/"
)

type Solaris struct {
	Proxy     string
	ProxyAuth string

	ServiceURL string
	Token      string

	Schemas SchemaList
}

func (sol *Solaris) Init(
	service_url string,
	proxy string,
	proxyAuth string,
	token string) {

	sol.Schemas.Load()

	sol.ServiceURL = service_url
	sol.Proxy = proxy
	sol.ProxyAuth = proxyAuth
	sol.Token = token

}

func (sol *Solaris) CheckSchema(name string) Schema {

	//fmt.Println("CheckSchema")
	cur := sol.Schemas.GetSchema(name)
	//fmt.Println("cur_sch |", cur)
	if cur.Table == "" {
		//fmt.Println("CheckSchema - empty - get from service |", name)
		cur = sol.GetSchema(name)
		//fmt.Println("cur_sch |", cur)
	}
	return cur

}

func (sol *Solaris) GetSchema(name string) Schema {

	var NewSchema Schema

	fmt.Println("GetSchema")
	NewSchema.Table = name
	NewSchema.GetSchema(sol.ServiceURL, sol.Proxy, sol.ProxyAuth, sol.Token)
	sol.Schemas.Add(NewSchema)
	sol.Schemas.Load()

	return NewSchema

}

/*
func (sol *Solaris) GetPeoplesByFio(Surname string, Name string, Fathername string, Birthdate string) People {

	var NewPeople People

	NewPeople.Surname = Surname
	NewPeople.Name = Name
	NewPeople.FatherName = Fathername
	NewPeople.Birthdate = Birthdate

	NewPeople.RequestByFio(sol.ServiceURL, sol.Proxy, sol.ProxyAuth, sol.Token)

	for index, value := range NewPeople.ParcedResponse {

		//fmt.Println("NewPeople ParcedResponse (befor)|", value)
		//fmt.Println(index, "|", value.Base)

		sc := sol.CheckSchema(value.Base)
		//fmt.Println("after_chaeck |", sc)
		NewPeople.ParcedResponse[index].Base_schema = sc
		//fmt.Println("NewPeople ParcedResponse (after)|", NewPeople.ParcedResponse[index].Base_schema)

	}
	//fmt.Println("NewPeople |", NewPeople)
	return NewPeople

}
*/

func (sol *Solaris) GetPeoplesFullRequest(
	Surname string,
	Name string,
	Fathername string,
	BirthdateDay string,
	BirthdateMonth string,
	BirthdateYear string,
	Inn string) People {

	var NewPeople People

	NewPeople.Surname = Surname
	NewPeople.Name = Name
	NewPeople.FatherName = Fathername
	NewPeople.BirthdateDay = BirthdateDay
	NewPeople.BirthdateMonth = BirthdateMonth
	NewPeople.BirthdateYear = BirthdateYear
	NewPeople.Inn = Inn

	res := NewPeople.FullRequest(sol.ServiceURL, sol.Proxy, sol.ProxyAuth, sol.Token)
	//fmt.Println(NewPeople)

	if res == true {

		for index, value := range NewPeople.ParcedResponse {
			sc := sol.CheckSchema(value.Base)
			NewPeople.ParcedResponse[index].Base_schema = sc
		}

	} else {
		//fmt.Println(NewPeople.Response)
	}

	return NewPeople

}
