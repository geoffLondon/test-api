package model

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	Nationality string `json:"nationality"`
	Investment  string `json:"investment"`
	Fund        Fund   `json:"fund"`
}

type Fund struct {
	Commodity string `json:"commodity,omitempty"`
	Equities  string `json:"equities,omitempty"`
	Hedge     string `json:"hedge,omitempty"`
}
