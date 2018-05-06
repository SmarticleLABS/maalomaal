package tests

import (
	"fmt"
	//"log"
	//"io/ioutil"
	//"encoding/json"
	"testing"

	"github.com/SmarticleLABS/maalomaal/wrapper_pkg"
)

func TestGetCompanyList(t *testing.T) {
	companies, err := wrapper_pkg.FetchCompaniesList()
	if nil != err {
		t.Errorf("User Authentication returned http error: %d.", err)
	}
	fmt.Printf("companies List is = %q\n", companies)
}
