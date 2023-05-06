package capsule

// HTTPRequest is the data of the http request
type HTTPRequest struct {
	Body    string
	JSONBody   string
	TextBody   string
	URI     string
	Method  string
	Headers string
}

// HTTPResponse is the data of the http response
type HTTPResponse struct {
	//Body    string
	JSONBody   string
	TextBody   string
	Headers    string
	StatusCode int
}
