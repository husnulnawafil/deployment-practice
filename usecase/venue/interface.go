package venue

import (
	_entities "capstone/entities"
)

type VenueUseCaseInterface interface {
	GetAllList(name string, category int) ([]_entities.GetVenuesResponse, error)
	CreateStep2(request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error)
	CreateStep1(request _entities.Venue, image string) (_entities.Venue, int, error)
	UpdateStep1(request _entities.Venue, id uint) (_entities.Venue, int, error)
	UpdateStep2(VenueID uint, request []_entities.Step2, facility []_entities.VenueFacility) ([]_entities.Step2, int, error)
	DeleteVenue(id uint) (int, error)
	GetVenueById(id int) (_entities.Venue, int, error)
	UpdateVenueImage(image string, id uint) (int, error)
}
