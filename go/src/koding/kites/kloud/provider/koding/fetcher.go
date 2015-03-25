package koding

import (
	"encoding/json"
	"errors"
	"fmt"
	"koding/db/models"
	"koding/db/mongodb"
	"net/http"
	"net/url"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type PaymentFetcher interface {
	Fetch(username string) (*PaymentResponse, error)
}

type SubscriptionsResponse struct {
	AccountId          string    `json:"accountId"`
	PlanTitle          string    `json:"planTitle"`
	PlanInterval       string    `json:"planInterval"`
	State              string    `json:"state"`
	CurrentPeriodStart time.Time `json:"currentPeriodStart"`
	CurrentPeriodEnd   time.Time `json:"currentPeriodEnd"`
	Description        string    `json:"description"`
	Error              string    `json:"error"`
}

type PaymentResponse struct {
	Plan  Plan
	State string
}

type Payment struct {
	DB              *mongodb.MongoDB
	PaymentEndpoint string
}

func (p *Payment) Fetch(username string) (*PaymentResponse, error) {
	if p.PaymentEndpoint == "" {
		return nil, errors.New("Payment endpoint is not set")
	}

	userEndpoint, err := url.Parse(p.PaymentEndpoint)
	if err != nil {
		return nil, err
	}

	var account *models.Account
	if err := p.DB.Run("jAccounts", func(c *mgo.Collection) error {
		return c.Find(bson.M{"profile.nickname": username}).One(&account)
	}); err != nil {
		return nil, err
	}

	q := userEndpoint.Query()
	q.Set("account_id", account.Id.Hex())
	userEndpoint.RawQuery = q.Encode()

	resp, err := http.Get(userEndpoint.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var subscription *SubscriptionsResponse
	e := json.NewDecoder(resp.Body)
	if err := e.Decode(&subscription); err != nil {
		return nil, err
	}

	// return an error back for a non 200 status
	if resp.StatusCode != 200 {
		if subscription.Error != "" {
			return nil, fmt.Errorf("could not fetch subscription. err: '%s'", subscription.Description)
		}

		return nil, fmt.Errorf("could not fetch subscription. status code: %d", resp.StatusCode)
	}

	plan, ok := plans[subscription.PlanTitle]
	if !ok {
		return nil, fmt.Errorf("could not find plan. There is no plan called '%s'", subscription.PlanTitle)
	}

	return &PaymentResponse{
		Plan:  plan,
		State: subscription.State,
	}, nil
}
