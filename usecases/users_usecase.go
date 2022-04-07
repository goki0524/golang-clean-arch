package usecases

import (
	"github.com/goki0524/clean-arch/infrastructure/repository"
	"github.com/goki0524/clean-arch/usecases/dto/usersdto"
)

type IUsersUseCase interface {
	Get(usersdto.GetInput) usersdto.GetOutput
}

type UsersUseCase struct {
	companyBedRepo repository.ICompanyBedRepository
}

func NewUsersUseCase(
	companyBedRepo repository.ICompanyBedRepository,
) *UsersUseCase {
	return &UsersUseCase{
		companyBedRepo: companyBedRepo,
	}
}

func (u *UsersUseCase) Get(in usersdto.GetInput) usersdto.GetOutput {
	id := in.ID
	companyBed := u.companyBedRepo.FindOne(id)

	if companyBed == nil {
		return usersdto.GetOutput{}
	}

	out := usersdto.GetOutput{}
	out.CompanyBed = *companyBed
	return out
}
