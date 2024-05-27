package models

type NetworkGroupPluralStructure struct {
	Objects []struct {
		Uid    string `json:"uid"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Domain struct {
			Uid        string `json:"uid"`
			Name       string `json:"name"`
			DomainType string `json:"domain-type"`
		} `json:"domain"`
		Icon  string `json:"icon"`
		Color string `json:"color"`
	} `json:"objects"`
	From  int `json:"from"`
	To    int `json:"to"`
	Total int `json:"total"`
}

type NetworkGroupStructure struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Domain struct {
		Uid        string `json:"uid"`
		Name       string `json:"name"`
		DomainType string `json:"domain-type"`
	} `json:"domain"`
	Members []struct {
		Uid    string `json:"uid"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Domain struct {
			Uid        string `json:"uid"`
			Name       string `json:"name"`
			DomainType string `json:"domain-type"`
		} `json:"domain"`
		Subnet4     string `json:"subnet4"`
		MaskLength4 int    `json:"mask-length4"`
		SubnetMask  string `json:"subnet-mask"`
		Icon        string `json:"icon"`
		Color       string `json:"color"`
	} `json:"members"`
	Groups   []interface{} `json:"groups"`
	Comments string        `json:"comments"`
	Color    string        `json:"color"`
	Icon     string        `json:"icon"`
	Tags     []interface{} `json:"tags"`
	MetaInfo struct {
		Lock            string `json:"lock"`
		ValidationState string `json:"validation-state"`
		LastModifyTime  struct {
			Posix   int64  `json:"posix"`
			Iso8601 string `json:"iso-8601"`
		} `json:"last-modify-time"`
		LastModifier string `json:"last-modifier"`
		CreationTime struct {
			Posix   int64  `json:"posix"`
			Iso8601 string `json:"iso-8601"`
		} `json:"creation-time"`
		Creator string `json:"creator"`
	} `json:"meta-info"`
	ReadOnly bool `json:"read-only"`
}
