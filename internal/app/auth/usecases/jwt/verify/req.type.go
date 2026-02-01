package verify

type Request struct {
	Token string
}

type Response struct {
	StaffID    string
	PositionID string
}