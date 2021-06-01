package covid

type Filter struct {
	ResponseType     *string         `json:"responseType"`
	Sexo             *[]*WhereClause `json:"sexo"`
	Edad             *[]*WhereClause `json:"edad"`
	Origen           *[]*WhereClause `json:"origen"`
	Defunto          *[]*WhereClause `json:"defunto"`
	EstadoResidencia *[]*WhereClause `json:"estadoResidencia"`
}

type WhereClause struct {
	Eq  *interface{} `json:"eq"`
	Ne  *interface{} `json:"ne"`
	Gte *interface{} `json:"gte"`
	Lte *interface{} `json:"lte"`
}
