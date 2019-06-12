package cosmos

import (
	"context"
	"fmt"
)

type Offer struct {
	client  Client
	offerID string
}

type OfferDefinition struct {
	OfferVersion string `json:"offerVersion"`
	OfferType    string `json:"offerType"`
	Content      struct {
		OfferThroughput int `json:"offerThroughput"`
	} `json:"content"`
	OfferResource   string `json:"resource"`
	OfferResourceID string `json:"offerResourceId"`
	Resource
}

type Offers struct {
	client Client
}

func newOffer(client Client, offerID string) *Offer {
	client.path = "offers/" + offerID
	client.rType = "offers"
	client.rLink = offerID
	offer := &Offer{
		client:  client,
		offerID: offerID,
	}

	return offer
}

func newOffers(client Client) *Offers {
	client.path = "offers"
	client.rType = "offers"
	fmt.Println("URL", client.getURL())
	offers := &Offers{
		client: client,
	}

	return offers
}

func (u *Offers) Query(ctx context.Context, query *SqlQuerySpec, opts ...CallOption) ([]OfferDefinition, error) {
	data := struct {
		Offers []OfferDefinition `json:"Offers,omitempty"`
		Count  int               `json:"_count,omitempty"`
	}{}
	_, err := u.client.query(ctx, query, &data, opts...)
	return data.Offers, err
}

func (u *Offer) Replace(ctx context.Context, offer *OfferDefinition, opts ...CallOption) (*OfferDefinition, error) {
	var updatedOffer *OfferDefinition
	_, err := u.client.replace(ctx, offer, updatedOffer, opts...)
	if err != nil {
		return nil, err
	}

	return updatedOffer, err
}

func (o *Offers) ReadAll(ctx context.Context, opts ...CallOption) ([]OfferDefinition, error) {
	data := struct {
		Offers []OfferDefinition `json:"offers,omitempty"`
		Count  int               `json:"_count,omitempty"`
	}{}

	_, err := o.client.read(ctx, &data, opts...)
	if err != nil {
		return nil, err
	}
	return data.Offers, err
}

func (o *Offer) Read(ctx context.Context, opts ...CallOption) (*OfferDefinition, error) {
	var offer OfferDefinition
	_, err := o.client.read(ctx, &offer, opts...)
	return &offer, err
}
