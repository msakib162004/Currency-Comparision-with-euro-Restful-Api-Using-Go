package structures

type Response struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type Rates_analyze map[string]map[string]float64

type AnalyzerResponse struct {
	Base          string                        `json:"base"`
	Rates_analyze map[string]map[string]float64 `json:"rates_analyze"`
}
