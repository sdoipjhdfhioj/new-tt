package service

import (
	"awesomeProject1/internal/entity"
	"awesomeProject1/internal/handler/dto"
	"awesomeProject1/internal/repository"
	"awesomeProject1/internal/repository/redis"
	"errors"
	"fmt"
)

type NameService interface {
	Set(nameRequest dto.SetNameRequest) error
	Get(id string) (string, error)
	Remove(id string) error
}

func InitNameService(addr string) *NameS {
	return &NameS{
		dataManager: redis.InitRedis(addr),
	}
}

type NameS struct {
	dataManager repository.DataManager
}

func (ns *NameS) Set(nameRequest dto.SetNameRequest) error {
	name := entity.Name{
		ID:   nameRequest.ID,
		Name: nameRequest.Name,
	}
	err := ns.dataManager.Set(name.ID, name.Name)
	if err != nil {
		return fmt.Errorf("while setting user name error: %w", err)
	}
	return nil
}

func (ns *NameS) Get(id string) (string, error) {
	userName, err := ns.dataManager.Get(id)
	if err != nil {
		if err.Error() == "redis: nil" {
			return "", errors.New("name wasn't set")
		}
		return "", fmt.Errorf("while getting user name error: %w", err)
	}
	return userName, nil
}

func (ns *NameS) Remove(id string) error {
	err := ns.dataManager.Remove(id)
	if err != nil {
		return fmt.Errorf("while removing user name error: %w", err)
	}
	return nil
}
