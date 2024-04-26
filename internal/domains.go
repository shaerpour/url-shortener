package internal

type Domain struct {
	ID     string `json: "id"`
	Domain string `json: "domain"`
}

var DOMAIN_LIST = []Domain{}
