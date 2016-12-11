package shared

type Station struct {
	Id     uint64 `json:"id"`
	Name   string `json:"name"`
	System string `json:"system"`
	Region string `json:"region"`
}
