package repository

import "event/config"

type Repository struct{
	config *config.Config
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := Repository{
		config: config,
	}

	return &r, nil
}
