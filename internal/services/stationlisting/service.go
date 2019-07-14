package stationlisting

import (
	"docomo-bike/internal/libs/docomo/getstation"
	"docomo-bike/internal/services/auth"
)

type Station struct {
	ID    string
	Name  string
	Bikes []*Bike
}
type Bike struct {
	ID string
}

type Service interface {
	GetStation(auth *auth.Auth, stationID string) (*Station, error)
}

func NewService(getStationClient getstation.Client) Service {
	return &DocomoService{
		GetStationClient: getStationClient,
	}
}

type DocomoService struct {
	GetStationClient getstation.Client
}

func (serv *DocomoService) GetStation(auth *auth.Auth, stationID string) (*Station, error) {
	s, err := serv.GetStationClient.GetStation(auth.UserID, auth.SessionKey, stationID)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}
	bb := []*Bike{}
	for _, b := range s.Bikes {
		bb = append(bb, &Bike{
			ID: b.ID,
		})
	}
	return &Station{
		ID:    stationID,
		Name:  s.Name,
		Bikes: bb,
	}, nil
}
