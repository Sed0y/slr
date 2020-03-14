package controllers

import (
	//"dpm/models"
	builder "dpm/models/builder"
	detail "dpm/models/builder/detail"
	dictionary "dpm/models/dictionary"
	entities "dpm/models/entities"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type EmployerPreContractDataController struct {
	beego.Controller
}

func (em *EmployerPreContractDataController) Post() {

	var peri dictionary.Period
	var pc builder.PreContractBuilder
	var empl entities.Employer
	var inst dictionary.TypeOfInsurance

	EmployerID, _ := strconv.Atoi(em.Ctx.Request.FormValue("id"))
	fmt.Println("begin is - " + em.Ctx.Request.FormValue("begin"))
	fmt.Println("end is - " + em.Ctx.Request.FormValue("end"))

	peri.Begin, _ = time.Parse("2006-01-02", em.Ctx.Request.FormValue("begin"))
	peri.End, _ = time.Parse("2006-01-02", em.Ctx.Request.FormValue("end"))

	fmt.Println(peri)
	//	InsType := em.Ctx.Request.Form["instype"]

	_, empl = App.Emloyers.GetById(EmployerID)
	inst.Set("КАСКО")
	//peri.Begin = time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	//peri.End = time.Date(2018, time.February, 1, 0, 0, 0, 0, time.UTC)

	var Res builder.
	//res := pc.GetEmployersResult(&inst, &peri)
	res := pc.GetEmployerResult(&empl, &inst, &peri)

	Res.PreContract.KASKO = append(Res.PreContract.KASKO, res)

	em.Ctx.ResponseWriter.Write([]byte(Res.ToJSON()))
}

type EmployerRegulirationDataController struct {
	beego.Controller
}

func (em *EmployerRegulirationDataController) Post() {

	var peri dictionary.Period
	var r builder.RegulationBuilder
	var empl entities.Employer
	var inst dictionary.TypeOfInsurance

	EmployerID, _ := strconv.Atoi(em.Ctx.Request.FormValue("id"))

	_, empl = App.Emloyers.GetById(EmployerID)
	peri.Begin, _ = time.Parse("2006-01-02", em.Ctx.Request.FormValue("begin"))
	peri.End, _ = time.Parse("2006-01-02", em.Ctx.Request.FormValue("end"))

	var Res builder.Results
	//res := pc.GetEmployersResult(&inst, &peri)

	inst.Set("КАСКО")
	res := r.GetEmployerResult(&empl, &inst, &peri)
	Res.Regulation.KASKO = append(Res.Regulation.KASKO, res)

	inst.Set("ОСАГО")
	res = r.GetEmployerResult(&empl, &inst, &peri)
	Res.Regulation.OSAGO = append(Res.Regulation.OSAGO, res)

	inst.Set("ГО и ЗК")
	res = r.GetEmployerResult(&empl, &inst, &peri)
	Res.Regulation.GO = append(Res.Regulation.GO, res)

	em.Ctx.ResponseWriter.Write([]byte(Res.ToJSON()))

}

type EmployerDataController struct {
	beego.Controller
}

func (em *EmployerDataController) Post() {

	var peri dictionary.Period
	var empl entities.Employer
	var inst dictionary.TypeOfInsurance

	EmployerID, _ := strconv.Atoi(em.Ctx.Request.FormValue("id"))
	peri.Begin, _ = time.Parse("2006-01-02", em.Ctx.Request.FormValue("begin"))
	peri.End, _ = time.Parse("2006-01-02", em.Ctx.Request.FormValue("end"))

	_, empl = App.Emloyers.GetById(EmployerID)

	var r builder.RegulationBuilder
	var pc builder.PreContractBuilder

	var Res builder.Results
	//res := pc.GetEmployersResult(&inst, &peri)

	inst.Set("КАСКО")
	res := r.GetEmployerResult(&empl, &inst, &peri)
	Res.Regulation.KASKO = append(Res.Regulation.KASKO, res)

	res = pc.GetEmployerResult(&empl, &inst, &peri)
	Res.PreContract.KASKO = append(Res.PreContract.KASKO, res)

	inst.Set("ОСАГО")
	res = r.GetEmployerResult(&empl, &inst, &peri)
	Res.Regulation.OSAGO = append(Res.Regulation.OSAGO, res)

	inst.Set("ГО и ЗК")
	res = r.GetEmployerResult(&empl, &inst, &peri)
	Res.Regulation.GO = append(Res.Regulation.GO, res)

	em.Ctx.ResponseWriter.Write([]byte(Res.ToJSON()))

}

type EmployerDataDetailController struct {
	beego.Controller
}

func (em *EmployerDataDetailController) Post() {

	var pc_details detail.PreContractDetails
	var reg_details detail.RegulationDetails

	var Res detail.DetailResults

	var peri dictionary.Period
	var empl entities.Employer
	var inst dictionary.TypeOfInsurance
	var job dictionary.TypeOfJob
	var restype dictionary.TypeOfResult

	EmployerID, _ := strconv.Atoi(em.Ctx.Request.FormValue("id"))
	peri.Begin, _ = time.Parse("2006-01-02", em.Ctx.Request.FormValue("begin"))
	peri.End, _ = time.Parse("2006-01-02", em.Ctx.Request.FormValue("end"))
	inst.Set(em.Ctx.Request.FormValue("instype"))
	job.Name = em.Ctx.Request.FormValue("job")
	restype.Name = em.Ctx.Request.FormValue("restype")

	_, empl = App.Emloyers.GetById(EmployerID)

	fmt.Println(empl)
	fmt.Println(peri)
	fmt.Println(em.Ctx.Request.FormValue("instype"), " - ", inst)
	fmt.Println(job)
	fmt.Println(restype)
	fmt.Println(job.IsPreContract())

	if job.IsPreContract() {
		Res = pc_details.GetEmployerDetails(&restype, &empl, &inst, &peri)
	}

	if job.IsRegulation() {
		Res = reg_details.GetEmployerDetails(&restype, &empl, &inst, &peri)
	}

	em.Ctx.ResponseWriter.Write([]byte(Res.ToJSON()))

}
