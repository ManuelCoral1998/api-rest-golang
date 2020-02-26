package model

// Domain is ...
type Domain struct {
	Servers          []Server `json:"endpoints" newtag:"servers"`
	ServerChanged    bool     `json:"serverChanged" newtag:"server_changed"`
	SslGrade         string   `json:"sslGrade" newtag:"ssl_grade"`
	PreviousSSLGrade string   `json:"previoussslGrade" newtag:"previous_ssl_grade"`
	Logo             string   `json:"logo" newtag:"logo"`
	Title            string   `json:"title" newtag:"title"`
	IsDown           bool     `json:"isDown" newtag:"is_down"`
}
