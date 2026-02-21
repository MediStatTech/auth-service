package create

type Request struct {
	Staff Staff
}

type Staff struct {
	PositionID string
	FirstName  string
	LastName   string
	Email      string
	Password   string
}
