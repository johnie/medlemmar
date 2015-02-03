package medlemmar

type Error struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

var Codes = map[int]string{
	400: "Bad Request",
	404: "Not Found",
}
