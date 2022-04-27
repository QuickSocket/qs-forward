package model

type Callback struct {
	AuthToken1     string `json:"authToken1"`
	AuthToken2     string `json:"authToken2"`
	Signature      string `json:"signature"`
	SerializedData string `json:"serializedData"`
}
