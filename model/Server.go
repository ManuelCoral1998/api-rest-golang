package model

// Server is ...
type Server struct {
	Address  string `json:"ipAddress" newtag:"address"`
	SslGrade string `json:"grade" newtag:"ssl_grade"`
	Country  string `json:"country" newtag:"country"`
	Owner    string `json:"owner" newtag:"owner"`
}
