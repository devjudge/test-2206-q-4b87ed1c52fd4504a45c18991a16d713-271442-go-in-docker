package service

import (
	"devjudge/go-in-docker/model"
	"fmt"
)

type LeadService struct {
}

var mapResp = make(map[int]model.Lead)
var maintainEmail = make(map[string]int)
var maintainMobile = make(map[string]int)

// var count int

func (LeadService) GetLeadByIdService(id int) *model.Lead {
	// resp := model.Lead{
	// 	Id:              id + count,
	// 	First_name:      "Hari",
	// 	Last_name:       "Nagula",
	// 	Mobile:          "9700704767",
	// 	Email:           "hariprasadnagula94@gmail.com",
	// 	Location_string: "Hyd",
	// }
	// mapResp[id+count] = resp
	// maintainEmail[resp.Email] = id + count
	// count++
	res, ok := mapResp[id]

	if ok {
		return &res
	}
	return nil
}

func (LeadService) CreateLead(lead model.Lead) (*model.Lead, error) {
	idCount := len(mapResp)
	_, ok := mapResp[lead.Id]
	if ok && idCount > 0 && lead.Id != 0 {
		return nil, fmt.Errorf("Id already Exists")
	}
	_, ok = maintainEmail[lead.Email]
	if ok {
		return nil, fmt.Errorf("Email already Exists")
	}
	_, ok = maintainMobile[lead.Mobile]
	if ok {
		return nil, fmt.Errorf("Mobile already Exists")
	}
	idIncr := lead.Id
	if idCount > 0 && lead.Id == 0 {
		idIncr = idCount
	}
	lead.Id = idIncr
	mapResp[idIncr] = lead
	maintainEmail[lead.Email] = idCount
	maintainMobile[lead.Mobile] = idCount
	return &lead, nil
}
