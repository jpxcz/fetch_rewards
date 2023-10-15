package reciept

import (
	"github.com/jpxcz/fetch_rewards/internal/repository/calculations"
)

const RETAILER_NAME_POINT = 1
const ZERO_CENTS_POINT = 50
const TWO_ITEMS_POINT = 5
const MULTIPLE_OF_25_POINT = 25
const DATE_ODD_POINT = 6
const BETWEEN_TIME_POINT = 10

type Reciept struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (r *Reciept) CreatePoints() (int, error) {
	retailerNameTotal := calculations.AlphanumericCharacters(r.Retailer, RETAILER_NAME_POINT)
	twoItemsTotal := calculations.TwoItemsPoints(len(r.Items), TWO_ITEMS_POINT)
	zeroCentsTotal, err := calculations.IsZeroCents(r.Total, ZERO_CENTS_POINT)
	if err != nil {
		return 0, err
	}

	multiple25Total, err := calculations.IsMultipleOf25(r.Total, MULTIPLE_OF_25_POINT)
	if err != nil {
		return 0, err
	}
	dateOddTotal, err := calculations.DateIsOddPoints(r.PurchaseDate, DATE_ODD_POINT)
	if err != nil {
		return 0, err
	}
	betweenHoursTotal, err := calculations.BetweenTimePoints(r.PurchaseTime, 14, 16, BETWEEN_TIME_POINT)
	if err != nil {
		return 0, err
	}
	itemsDescriptionTotal, err := r.calculateItemsTotal()
	if err != nil {
		return 0, err
	}

	total := retailerNameTotal + zeroCentsTotal + twoItemsTotal + multiple25Total + betweenHoursTotal + dateOddTotal + itemsDescriptionTotal
	return total, nil
}

func (r *Reciept) calculateItemsTotal() (int, error) {
	var totalPoints = 0
	for _, i := range r.Items {
		p, err := calculations.TrimmedDescriptionPoints(i.ShortDescription, i.Price)
		if err != nil {
			return totalPoints, err
		}
		totalPoints = totalPoints + p
	}

	return totalPoints, nil
}
