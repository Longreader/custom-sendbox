package models

type LoginStructure struct {
	Uid            string `json:"uid"`
	Sid            string `json:"sid"`
	Url            string `json:"url"`
	SessionTimeout int    `json:"session-timeout"`
	LastLoginWasAt struct {
		Posix   int64  `json:"posix"`
		Iso8601 string `json:"iso-8601"`
	} `json:"last-login-was-at"`
	ApiServerVersion string `json:"api-server-version"`
}

type TaskStructure struct {
	Tasks []struct {
		Target string `json:"target"`
		TaskId string `json:"task-id"`
	} `json:"tasks"`
}
