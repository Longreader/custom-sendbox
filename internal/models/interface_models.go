package models

type InterfacePluralStructure struct {
	Objects []struct {
		Uid    string `json:"uid"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Domain struct {
			Uid        string `json:"uid"`
			Name       string `json:"name"`
			DomainType string `json:"domain-type"`
		} `json:"domain"`
		Policy struct {
			ClusterMembersAccessPolicyRevision []struct {
				Name       string `json:"name"`
				Uid        string `json:"uid"`
				PolicyName string `json:"policy-name"`
				Revision   struct {
					Uid    string `json:"uid"`
					Name   string `json:"name"`
					Type   string `json:"type"`
					Domain struct {
						Uid        string `json:"uid"`
						Name       string `json:"name"`
						DomainType string `json:"domain-type"`
					} `json:"domain"`
					State         string `json:"state"`
					UserName      string `json:"user-name"`
					Description   string `json:"description"`
					LastLoginTime struct {
						Posix   int64  `json:"posix"`
						Iso8601 string `json:"iso-8601"`
					} `json:"last-login-time"`
					PublishTime struct {
						Posix   int64  `json:"posix"`
						Iso8601 string `json:"iso-8601"`
					} `json:"publish-time"`
					ExpiredSession bool          `json:"expired-session"`
					Application    string        `json:"application"`
					Changes        int           `json:"changes"`
					InWork         bool          `json:"in-work"`
					IpAddress      string        `json:"ip-address"`
					Locks          int           `json:"locks"`
					ConnectionMode string        `json:"connection-mode"`
					SessionTimeout int           `json:"session-timeout"`
					Comments       string        `json:"comments"`
					Color          string        `json:"color"`
					Icon           string        `json:"icon"`
					Tags           []interface{} `json:"tags"`
					MetaInfo       struct {
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
				} `json:"revision"`
			} `json:"cluster-members-access-policy-revision,omitempty"`
			AccessPolicyInstalled        bool   `json:"access-policy-installed,omitempty"`
			AccessPolicyName             string `json:"access-policy-name,omitempty"`
			AccessPolicyInstallationDate struct {
				Posix   int64  `json:"posix"`
				Iso8601 string `json:"iso-8601"`
			} `json:"access-policy-installation-date,omitempty"`
			ThreatPolicyInstalled bool `json:"threat-policy-installed,omitempty"`
		} `json:"policy,omitempty"`
		Ipv4Address string `json:"ipv4-address"`
		Interfaces  []struct {
			InterfaceName   string `json:"interface-name,omitempty"`
			Ipv4Address     string `json:"ipv4-address,omitempty"`
			Ipv4NetworkMask string `json:"ipv4-network-mask,omitempty"`
			Ipv4MaskLength  int    `json:"ipv4-mask-length,omitempty"`
			DynamicIp       bool   `json:"dynamic-ip,omitempty"`
			Topology        struct {
				LeadsToInternet              bool   `json:"leads-to-internet"`
				IpAddressBehindThisInterface string `json:"ip-address-behind-this-interface"`
				LeadsToDmz                   bool   `json:"leads-to-dmz"`
				SecurityZone                 struct {
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
				} `json:"security-zone,omitempty"`
			} `json:"topology,omitempty"`
			Uid    string `json:"uid,omitempty"`
			Name   string `json:"name,omitempty"`
			Type   string `json:"type,omitempty"`
			Domain struct {
				Uid        string `json:"uid"`
				Name       string `json:"name"`
				DomainType string `json:"domain-type"`
			} `json:"domain,omitempty"`
			Subnet4     string        `json:"subnet4,omitempty"`
			MaskLength4 int           `json:"mask-length4,omitempty"`
			SubnetMask  string        `json:"subnet-mask,omitempty"`
			Comments    string        `json:"comments,omitempty"`
			Color       string        `json:"color,omitempty"`
			Icon        string        `json:"icon,omitempty"`
			Tags        []interface{} `json:"tags,omitempty"`
		} `json:"interfaces"`
		SicStatus string        `json:"sic-status,omitempty"`
		Tags      []interface{} `json:"tags"`
		Icon      string        `json:"icon"`
		Groups    []interface{} `json:"groups"`
		Comments  string        `json:"comments"`
		Color     string        `json:"color"`
		MetaInfo  struct {
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
			Creator          string `json:"creator"`
			LockingAdmin     string `json:"locking-admin,omitempty"`
			LockingSessionId string `json:"locking-session-id,omitempty"`
		} `json:"meta-info"`
		ReadOnly              bool   `json:"read-only"`
		OperatingSystem       string `json:"operating-system,omitempty"`
		Hardware              string `json:"hardware,omitempty"`
		Version               string `json:"version,omitempty"`
		NetworkSecurityBlades struct {
			Firewall bool `json:"firewall"`
		} `json:"network-security-blades,omitempty"`
		ManagementBlades struct {
			NetworkPolicyManagement bool `json:"network-policy-management,omitempty"`
			UserDirectory           bool `json:"user-directory,omitempty"`
			Compliance              bool `json:"compliance,omitempty"`
			LoggingAndStatus        bool `json:"logging-and-status,omitempty"`
			SmartEventServer        bool `json:"smart-event-server,omitempty"`
			SmartEventCorrelation   bool `json:"smart-event-correlation,omitempty"`
			EndpointPolicy          bool `json:"endpoint-policy,omitempty"`
			Secondary               bool `json:"secondary,omitempty"`
			IdentityLogging         bool `json:"identity-logging,omitempty"`
		} `json:"management-blades,omitempty"`
		ClusterMemberNames  []string `json:"cluster-member-names,omitempty"`
		VpnEncryptionDomain string   `json:"vpn-encryption-domain,omitempty"`
		NatSettings         struct {
			AutoRule bool `json:"auto-rule"`
		} `json:"nat-settings,omitempty"`
		Os           string `json:"os,omitempty"`
		SicName      string `json:"sic-name,omitempty"`
		SicState     string `json:"sic-state,omitempty"`
		LogsSettings struct {
			EnableLogIndexing                            bool   `json:"enable-log-indexing"`
			SmartEventIntroCorrelationUnit               bool   `json:"smart-event-intro-correlation-unit"`
			AcceptSyslogMessages                         bool   `json:"accept-syslog-messages"`
			RotateLogByFileSize                          bool   `json:"rotate-log-by-file-size"`
			RotateLogFileSizeThreshold                   int    `json:"rotate-log-file-size-threshold"`
			RotateLogOnSchedule                          bool   `json:"rotate-log-on-schedule"`
			AlertWhenFreeDiskSpaceBelowMetrics           string `json:"alert-when-free-disk-space-below-metrics"`
			AlertWhenFreeDiskSpaceBelow                  bool   `json:"alert-when-free-disk-space-below"`
			AlertWhenFreeDiskSpaceBelowThreshold         int    `json:"alert-when-free-disk-space-below-threshold"`
			AlertWhenFreeDiskSpaceBelowType              string `json:"alert-when-free-disk-space-below-type"`
			DeleteWhenFreeDiskSpaceBelowMetrics          string `json:"delete-when-free-disk-space-below-metrics"`
			DeleteWhenFreeDiskSpaceBelow                 bool   `json:"delete-when-free-disk-space-below"`
			DeleteWhenFreeDiskSpaceBelowThreshold        int    `json:"delete-when-free-disk-space-below-threshold"`
			BeforeDeleteKeepLogsFromTheLastDays          bool   `json:"before-delete-keep-logs-from-the-last-days"`
			BeforeDeleteKeepLogsFromTheLastDaysThreshold int    `json:"before-delete-keep-logs-from-the-last-days-threshold"`
			BeforeDeleteRunScript                        bool   `json:"before-delete-run-script"`
			BeforeDeleteRunScriptCommand                 string `json:"before-delete-run-script-command"`
			StopLoggingWhenFreeDiskSpaceBelowMetrics     string `json:"stop-logging-when-free-disk-space-below-metrics"`
			StopLoggingWhenFreeDiskSpaceBelow            bool   `json:"stop-logging-when-free-disk-space-below"`
			StopLoggingWhenFreeDiskSpaceBelowThreshold   int    `json:"stop-logging-when-free-disk-space-below-threshold"`
			DeleteIndexFilesOlderThanDays                bool   `json:"delete-index-files-older-than-days"`
			DeleteIndexFilesOlderThanDaysThreshold       int    `json:"delete-index-files-older-than-days-threshold"`
			ForwardLogsToLogServer                       bool   `json:"forward-logs-to-log-server"`
			UpdateAccountLogEvery                        int    `json:"update-account-log-every"`
			DetectNewCitrixIcaApplicationNames           bool   `json:"detect-new-citrix-ica-application-names"`
			TurnOnQosLogging                             bool   `json:"turn-on-qos-logging"`
		} `json:"logs-settings,omitempty"`
	} `json:"objects"`
	From  int `json:"from"`
	To    int `json:"to"`
	Total int `json:"total"`
}

type InterfaceStructure struct {
	InterfaceName   string `json:"interface-name,omitempty"`
	Ipv4Address     string `json:"ipv4-address,omitempty"`
	Ipv4NetworkMask string `json:"ipv4-network-mask,omitempty"`
	Ipv4MaskLength  int    `json:"ipv4-mask-length,omitempty"`
	DynamicIp       bool   `json:"dynamic-ip,omitempty"`
	Topology        struct {
		LeadsToInternet              bool   `json:"leads-to-internet"`
		IpAddressBehindThisInterface string `json:"ip-address-behind-this-interface"`
		LeadsToDmz                   bool   `json:"leads-to-dmz"`
		SecurityZone                 struct {
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
		} `json:"security-zone,omitempty"`
	} `json:"topology,omitempty"`
	Uid    string `json:"uid,omitempty"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Domain struct {
		Uid        string `json:"uid"`
		Name       string `json:"name"`
		DomainType string `json:"domain-type"`
	} `json:"domain,omitempty"`
	Subnet4     string        `json:"subnet4,omitempty"`
	MaskLength4 int           `json:"mask-length4,omitempty"`
	SubnetMask  string        `json:"subnet-mask,omitempty"`
	Comments    string        `json:"comments,omitempty"`
	Color       string        `json:"color,omitempty"`
	Icon        string        `json:"icon,omitempty"`
	Tags        []interface{} `json:"tags,omitempty"`
}
