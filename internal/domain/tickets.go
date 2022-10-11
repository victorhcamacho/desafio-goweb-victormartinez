package domain

type Ticket struct {
	Id      string
	Name    string
	Email   string
	Country string
	Time    string
	Price   float64
}

func NewTicket(id string, name string, email string, country string, time string, price float64) *Ticket {
	return &Ticket{id, name, email, country, time, price}
}
