package models

type RouteTableTask struct {
	Tasks []struct {
		Uid    string `json:"uid"`
		Name   string `json:"name"`
		Type   string `json:"type"`
		Domain struct {
			Uid        string `json:"uid"`
			Name       string `json:"name"`
			DomainType string `json:"domain-type"`
		} `json:"domain"`
		TaskId             string `json:"task-id"`
		TaskName           string `json:"task-name"`
		Status             string `json:"status"`
		ProgressPercentage int    `json:"progress-percentage"`
		StartTime          struct {
			Posix   int64  `json:"posix"`
			Iso8601 string `json:"iso-8601"`
		} `json:"start-time"`
		LastUpdateTime struct {
			Posix   int64  `json:"posix"`
			Iso8601 string `json:"iso-8601"`
		} `json:"last-update-time"`
		Suppressed  bool `json:"suppressed"`
		TaskDetails []struct {
			Uid    string      `json:"uid"`
			Name   interface{} `json:"name"`
			Domain struct {
				Uid        string `json:"uid"`
				Name       string `json:"name"`
				DomainType string `json:"domain-type"`
			} `json:"domain"`
			Color             string `json:"color"`
			StatusCode        string `json:"statusCode"`
			StatusDescription string `json:"statusDescription"`
			TaskNotification  string `json:"taskNotification"`
			GatewayId         string `json:"gatewayId"`
			GatewayName       string `json:"gatewayName"`
			TransactionId     int    `json:"transactionId"`
			ResponseMessage   string `json:"responseMessage"`
			ResponseError     string `json:"responseError"`
			MetaInfo          struct {
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
			Tags         []interface{} `json:"tags"`
			Icon         string        `json:"icon"`
			Comments     string        `json:"comments"`
			DisplayName  string        `json:"display-name"`
			CustomFields interface{}   `json:"customFields"`
		} `json:"task-details"`
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
	} `json:"tasks"`
}
