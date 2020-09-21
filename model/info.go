package model

type Info struct {
	Tags []string `json:"tags"`
	Group string `json:"group"`
	Title string `json:"title"`
	LinkType string `json:"link_type"`
	LinkUrl string `json:"link_url"`
}

func GetGroupTypes() []string {
	a := []string{"Tax", "Benefit", "Immigration", "HouseAndMortgage", "News"}
	return a
}


