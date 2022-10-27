package dog

import "time"

type LogResponse struct {
	Data []AttributeHolder `json:"data"`
	Meta MetaHolder        `json:"meta"`
}

type MetaHolder struct {
	Page PageHolder `json:"page"`
}
type PageHolder struct {
	After string `json:"after"`
}

type AttributeHolder struct {
	Attributes Attribute `json:"attributes"`
	Id         string    `json:"id"`
}

type Attribute struct {
	Service       string       `json:""service"`
	Tags          []string     `json:"tags"`
	Status        string       `json:"status"`
	Timestamp     time.Time    `json:"timestamp"`
	Host          string       `json:"host"`
	SubAttributes SubAttribute `json:"attributes"`
	Message       string       `json"message"`
}

type SubAttribute struct {
	Msg       string `json:"msg"`
	Exception string `json:"exception"`
}
