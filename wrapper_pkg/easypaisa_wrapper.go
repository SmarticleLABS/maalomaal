package wrapper_pkg

import (
	"fmt"
	"log"

	"github.com/SmarticleLABS/maalomaal/easypaisa_lib/easypaisa_pkg"
)

type CompaniesList struct {
	Companies   []*Company `json:"companies" form:"companies"`     //TODO: Write general description for this field
	Description string     `json:"description" form:"description"` //TODO: Write general description for this field
	//ResultCode      string          `json:"resultCode" form:"resultCode"` //TODO: Write general description for this field
}
type Company struct {
	Name            string `json:"company_name" form:"company_name"`         //TODO: Write general description for this field
	ServiceProvider string `json:"service_provider" form:"service_provider"` //TODO: Write general description for this field
	SubType         string `json:"sub_type" form:"sub_type"`                 //TODO: Write general description for this field
	Type            string `json:"type" form:"type"`                         //TODO: Write general description for this field
}

func FetchCompaniesList() (*CompaniesList, error) {
	ep := &easypaisa_pkg.EASYPAISA_IMPL{}
	compListResp, err := ep.GetCompanyList()
	if nil != err {
		log.Println(err)
		return nil, err
	}

	companies := make([]*Company, 0)
	for _, comp := range compListResp.Companies {
		company := &Company{Name: comp.CompanyName, ServiceProvider: comp.ServiceProvider, SubType: comp.SubType, Type: comp.Type}
		companies = append(companies, company)
	}

	fmt.Printf("Company List is:%+v\n", compListResp)
	companiesList := &CompaniesList{Companies: companies, Description: compListResp.Description}

	return companiesList, nil
}
