package facility

import (
	_entities "capstone/entities"
	_facilityRepository "capstone/repository/facility"
)

type FacilityUseCase struct {
	facilityRepository _facilityRepository.FacilityRepositoryInterface
}

func NewFacilityUseCase(facilityRepo _facilityRepository.FacilityRepositoryInterface) FacilityUseCaseInterface {
	return &FacilityUseCase{
		facilityRepository: facilityRepo,
	}
}

func (cuc *FacilityUseCase) GetAllFacility() ([]_entities.Facility, error) {
	facility, err := cuc.facilityRepository.GetAllFacility()
	return facility, err
}

func (cuc *FacilityUseCase) CreateFacility(request _entities.Facility) (_entities.Facility, error) {
	facility, err := cuc.facilityRepository.CreateFacility(request)
	return facility, err
}

func (cuc *FacilityUseCase) UpdateFacility(id uint, request _entities.Facility) (_entities.Facility, int, error) {
	facility, rows, err := cuc.facilityRepository.UpdateFacility(id, request)
	return facility, rows, err
}

func (cuc *FacilityUseCase) DeleteFacility(id int) error {
	err := cuc.facilityRepository.DeleteFacility(id)
	return err
}
