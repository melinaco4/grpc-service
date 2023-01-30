package ticket

import "context"

type Ticket struct {
	ID    string
	Name  string
	Type  string
	Event int
}

type Store interface {
	GetTicketByID(id string) (Ticket, error)
	InsertTicket(t Ticket) (Ticket, error)
	DeleteRocket(id string) error
}

// Service - Ticket service that updates ticket inventory

type Service struct {
	Store Store
}

func New(store Store) Service {
	return Service{
		Store: store,
	}
}

func (s Service) GetTicketById(ctx context.Context, id string) (Ticket, error) {
	t, err := s.Store.GetTicketByID(id)
	if err != nil {
		return Ticket{}, err
	}
	return t, nil
}

func (s Service) InsertTicket(ctx context.Context, t Ticket) (Ticket, error) {
	t, err := s.Store.InsertTicket(t)
	if err != nil {
		return Ticket{}, err
	}
	return t, nil
}

func (s *Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}
	return nil
}
