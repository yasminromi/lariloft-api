package service

import (
	"context"
	"errors"

	"github.com/yasminromi/lariloft-api/model"
	"gopkg.in/olivere/elastic.v6"
)

const (
	indexName = "interests"
	docType   = "log"
	appName   = "lariloft-api"
)

type ElasticService struct {
	ElasticCLI *elastic.Client
}

func (e *ElasticService) SaveToElastic(ctx context.Context, payload model.Interest) error {

	exists, err := e.ElasticCLI.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		res, error := e.ElasticCLI.CreateIndex(indexName).Do(ctx)
		if error != nil {
			return error
		}
		if !res.Acknowledged {
			return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
		}
	}

	var user = model.User{
		Name:            payload.User.Name,
		Age:             payload.User.Age,
		Work:            payload.User.Work,
		ActualNeighbor:  payload.User.ActualNeighbor,
		DesiredNeighbor: payload.User.DesiredNeighbor,
		Motive:          payload.User.Motive,
		Rooms:           payload.User.Rooms,
		CoLivers:        payload.User.CoLivers,
		Kids:            payload.User.Kids,
		Pet:             payload.User.Pet,
		InterestHome:    payload.User.InterestHome,
		InterestOutside: payload.User.InterestOutside,
		InterestRooms:   payload.User.InterestRooms,
		Budget:          payload.User.Budget,
		Msisdn:          payload.User.Msisdn,
	}

	var apartment = model.Apartment{
		Lat:              payload.Apartment.Lat,
		Long:             payload.Apartment.Long,
		NumberOfRooms:    payload.Apartment.NumberOfRooms,
		MaxEstimate:      payload.Apartment.MaxEstimate,
		MinEstimate:      payload.Apartment.MinEstimate,
		KitchenSize:      payload.Apartment.KitchenSize,
		LivingRoomSize:   payload.Apartment.LivingRoomSize,
		HasOffice:        payload.Apartment.HasOffice,
		HasPool:          payload.Apartment.HasPool,
		AcceptPet:        payload.Apartment.AcceptPet,
		PriceCondominium: payload.Apartment.PriceCondominium,
	}

	var interest = model.Interest{
		CreatedAt: payload.CreatedAt,
		Apartment: apartment,
		User:      user,
	}

	_, error := e.ElasticCLI.Index().
		Index(indexName).
		Type(docType).
		BodyJson(interest).
		Do(ctx)

	if error != nil {
		return error
	}

	return nil
}
