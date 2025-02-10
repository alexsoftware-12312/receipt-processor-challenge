package model

import "regexp"

type GetPointsResponse struct {
	Points int `json:"points"`
}

type ProcessReceiptResponse struct {
	Id string `json:"id"`
}

type ProcessReceiptRequest struct {
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Items        []ReceiptItem `json:"items"`
	Total        string        `json:"total"`
}

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func Validate(r ProcessReceiptRequest) bool {
	res := true

	if r.Retailer == "" || r.PurchaseDate == "" || r.PurchaseTime == "" || r.Items == nil || r.Total == "" {
		res = false
	}

	check, err := regexp.MatchString("^[\\w\\s\\-&]+$", r.Retailer)
	if err != nil{
		res = false
	} else if !check {
		res = false
	}

	check, err = regexp.MatchString("^\\d{4}-\\d{2}-\\d{2}$", r.PurchaseDate)
	if err != nil{
		res = false
	} else if !check {
		res = false
	}

	check, err = regexp.MatchString("^\\d{2}:\\d{2}$", r.PurchaseTime)
	if err != nil{
		res = false
	} else if !check {
		res = false
	}

	for _,i := range r.Items {
		
		check, err = regexp.MatchString("^[\\w\\s\\-]+$", i.ShortDescription)
		if err != nil{
			res = false
		} else if !check {
			res = false
		}

		check, err = regexp.MatchString("^\\d+\\.\\d{2}$", i.Price)
		if err != nil{
			res = false
		} else if !check {
			res = false
		}
	}

	check, err = regexp.MatchString("^\\d+\\.\\d{2}$", r.Total)
	if err != nil{
		res = false
	} else if !check {
		res = false
	}



	return res
}