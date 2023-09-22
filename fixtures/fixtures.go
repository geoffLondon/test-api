package fixtures

import (
	"net/url"

	"api-with-interfaces/model"
)

func ValidFormValues() url.Values {
	form := url.Values{}
	form.Add("name", "Molly")
	form.Add("age", "30")
	form.Add("nationality", "UK")
	form.Add("investment", "5000")
	form.Add("fund1", "equities")

	return form
}

func ValidCustomer() model.Customer {
	return model.Customer{
		ID:          "mocked-uuid",
		Name:        "Molly",
		Age:         "30",
		Nationality: "UK",
		Investment:  "5000",
		Fund: model.Fund{
			Equities: "equities",
		},
	}
}

func ValidCustomers() []model.Customer {
	return []model.Customer{
		{
			ID:          "mocked-uuid",
			Name:        "Molly",
			Age:         "30",
			Nationality: "UK",
			Investment:  "5000",
			Fund: model.Fund{
				Equities: "equities",
			},
		},
		{
			ID:          "mocked-uuid-2",
			Name:        "Geoff",
			Age:         "50",
			Nationality: "UK",
			Investment:  "10000",
			Fund: model.Fund{
				Equities: "equities",
			},
		},
	}
}
