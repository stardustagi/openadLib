package http_service

type Config struct {
	IP    string   `json:"ip"`
	Port  int      `json:"port"`
	Group []string `json:"group"`
}
