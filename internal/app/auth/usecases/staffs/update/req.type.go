package sign_in

type Request struct {
	StaffID string
	PositionID string
}

type Response struct {
	Token string
}