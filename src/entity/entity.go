package entity

// TradeMade struct
type TradeMade struct {
	Endpoint       string                   `json:"endpoint"`
	Quotes         []map[string]interface{} `json:"quotes"`
	Requested_time string                   `json:"requested_time"`
	Timestamp      int32                    `json:"timestamp"`
}
