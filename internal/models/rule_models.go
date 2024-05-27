package models

type RulePluralStructure struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Rulebase []struct {
		Uid   string          `json:"uid"`
		Name  string          `json:"name"`
		Type  string          `json:"type"`
		From  int             `json:"from"`
		To    int             `json:"to"`
		Rules []RuleStructure `json:"rules"`
	} `json:"rulebase"`
	From  int `json:"from"`
	To    int `json:"to"`
	Total int `json:"total"`
}

type RuleStructure struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Domain struct {
		Uid        string `json:"uid"`
		Name       string `json:"name"`
		DomainType string `json:"domain-type"`
	} `json:"domain"`
	RuleNumber int `json:"rule-number"`
	Track      struct {
		Type struct {
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
		} `json:"type"`
		PerSession            bool   `json:"per-session"`
		PerConnection         bool   `json:"per-connection"`
		Accounting            bool   `json:"accounting"`
		EnableFirewallSession bool   `json:"enable-firewall-session"`
		Alert                 string `json:"alert"`
	} `json:"track"`
	Source []struct {
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
	} `json:"source"`
	SourceNegate bool `json:"source-negate"`
	Destination  []struct {
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
	} `json:"destination"`
	DestinationNegate bool `json:"destination-negate"`
	Service           []struct {
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
	} `json:"service"`
	ServiceNegate   bool   `json:"service-negate"`
	ServiceResource string `json:"service-resource"`
	Vpn             []struct {
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
	} `json:"vpn"`
	Action struct {
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
	} `json:"action"`
	ActionSettings struct {
		EnableIdentityCaptivePortal bool `json:"enable-identity-captive-portal"`
	} `json:"action-settings"`
	Content []struct {
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
	} `json:"content"`
	ContentNegate    bool   `json:"content-negate"`
	ContentDirection string `json:"content-direction"`
	Time             []struct {
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
	} `json:"time"`
	Hits struct {
		Percentage string `json:"percentage"`
		Level      string `json:"level"`
		Value      int    `json:"value"`
	} `json:"hits"`
	CustomFields struct {
		Field1 string `json:"field-1"`
		Field2 string `json:"field-2"`
		Field3 string `json:"field-3"`
	} `json:"custom-fields"`
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
	Comments  string `json:"comments"`
	Enabled   bool   `json:"enabled"`
	InstallOn []struct {
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
	} `json:"install-on"`
}
