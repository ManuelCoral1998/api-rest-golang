package test

import (
	"fmt"
	"model"
	"strings"
	"testing"
)

func TestCreateStruct(t *testing.T) {

	servers := make([]model.Server, 2)

	servers[0] = model.Server{"192.168.0.1", "A", "Colombia", "Mine"}
	servers[1] = model.Server{"192.168.0.2", "B", "Colombia", "Mine"}

	domain := model.Domain{servers, false, "B", "B", "--", "midireccion", true}

	if strings.Compare(domain.Servers[0].Address, "192.168.0.1") != 0 {
		t.Errorf("The struct server is not working")
	}

}

func TestGetInformationCountryServer(t *testing.T) {

	var logicModel model.LogicModel
	domain := logicModel.GetInformationFromServers("truora.com")

	fmt.Println("SSL: " + domain.SslGrade)
	if strings.Compare(domain.Servers[0].Country, "US") != 0 && strings.Compare(domain.Servers[1].Country, "US") != 0 {
		t.Errorf("The method getInformationFromServer is not bringing the correct ssl grade of the domain")
	}

}

func TestGetInformationOwnerServer(t *testing.T) {

	var logicModel model.LogicModel
	domain := logicModel.GetInformationFromServers("truora.com")

	if strings.Compare(domain.Servers[0].Owner, "Amazon Technologies Inc.") != 0 && strings.Compare(domain.Servers[1].Owner, "Amazon Technologies Inc.") != 0 {
		t.Errorf("The method getInformationFromServer is not bringing the correct server's owner of the domain")
	}

}

func TestGetInformationFromDomain(t *testing.T) {

	var logicModel model.LogicModel
	domain := logicModel.GetInformationFromServers("truora.com")

	domain = logicModel.GetInformationFromDomain(domain, "truora.com")

	if strings.Compare(domain.Title, "Truora | Validación de Antecedentes e Identidad Digital") != 0 {
		t.Errorf("The method GetInformationFromDomain is not bringing the correct domain's title")
	}

}
