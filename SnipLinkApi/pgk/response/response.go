package response

type Request struct {
	Link string `json:"link" validate:"required,url"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Link   string `json:"link,omitempty"`
}

func NewOK() Response {
	return Response{
		Status: "OK",
	}
}

func NewError(msg string) Response {
	return Response{
		Status: "Error",
		Error:  msg,
	}
}

func NewOkLink(link string) Response {
	return Response{
		Status: "OK",
		Link:   link,
	}
}
