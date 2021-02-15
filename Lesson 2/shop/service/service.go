package service

import (
	"errors"
	"fmt"
	"log"
	"shop/models"
	"shop/notification"
	"shop/repository"
)

type Service interface {
	CreateOrder(order *models.Order) (*models.Order, error)
}

type service struct {
	notif notification.Notification
	rep   repository.Repository
	notifEmail notification.SendEmailNotif
}

var (
	ErrItemNotExists = errors.New("item not exists")
)

func (s *service) CreateOrder(order *models.Order) (*models.Order, error) {
	for _, itemID := range order.ItemIDs {
		_, err := s.rep.GetItem(itemID)
		if err != nil && err != repository.ErrNotFound {
			return nil, err
		}
		if err == repository.ErrNotFound {
			return nil, ErrItemNotExists
		}
	}

	order, err := s.rep.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	if err := s.notif.SendOrderCreated(order); err != nil {
		log.Println(err)
	}

	if err := s.notifEmail.SendEmailOrderNotification(order); err != nil {
		log.Println(err)
	}
	fmt.Println("Send email")
	return order, nil
}

func NewService(rep repository.Repository, notif notification.Notification, notifEmail notification.SendEmailNotif) Service {
	return &service{
		notif: notif,
		rep:   rep,
		notifEmail: notifEmail,
	}
}
