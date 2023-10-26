package manager

import "payment-application/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	ProfileRepo() repository.ProfileRepository
	MerchantRepo() repository.MerchantRepository
	TransferRepo() repository.TransferRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Connection())
}

func (r *repoManager) ProfileRepo() repository.ProfileRepository {
	return repository.NewProfileRepository(r.infra.Connection())
}

func (r *repoManager) MerchantRepo() repository.MerchantRepository {
	return repository.NewMerchantRepository(r.infra.Connection())
}

func (r *repoManager) TransferRepo() repository.TransferRepository {
	return repository.NewTransferRepository(r.infra.Connection())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
