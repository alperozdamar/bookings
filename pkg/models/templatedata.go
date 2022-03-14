package models

// TempalteData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //If you don't know the data type for now.
	CSRFToken string                 //Cross-Site Request Forgery
	Flash     string
	Warning   string
	Error     string
}
