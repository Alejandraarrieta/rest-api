package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Observations struct {
	IsFeatured        string
	IsFixed           string
	IsMessage         string
	Message           string
	IsBussinesGravanz string
	IsPort            string
}

type PlaceOfDelivery struct {
	Zone string
	Port string
}

type Product struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ProductName     string             `json:"productname,omitempty" bson:"productname,omitempty"`
	Price           string             `json:"price,omitempty" bson:"price,omitempty"`
	TypeCoin        string             `json:"typecoin,omitempty" bson:"typecoin,omitempty"`
	TypeBusiness    string             `json:"typebusiness,omitempty" bson:"typebusiness,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
	DeliveryPeriod  string             `json:"deliveryperiod,omitempty" bson:"deliveryperiod,omitempty"`
	Quality         string             `json:"quality,omitempty" bson:"quality,omitempty"`
	WayPay          string             `json:"waypay,omitempty" bson:"waypay,omitempty"`
	Expiration      string             `json:"expiration,omitempty" bson:"expiration,omitempty"`
	HourP           string             `json:"hourp,omitempty" bson:"hourp,omitempty"`
	PriceId         string             `json:"price_id,omitempty" bson:"price_id,omitempty"`
	Harvest         string             `json:"harvest,omitempty" bson:"harvest,omitempty"`
	Bonus           string             `json:"bonus,omitempty" bson:"bonus,omitempty"`
	Reduction       string             `json:"reduction,omitempty" bson:"reduction,omitempty"`
	Pesification    string             `json:"pesification,omitempty" bson:"pesification,omitempty"`
	Observations    Observations       `json:"observations,inline" bson:"observations,inline"`
	PlaceOfDelivery PlaceOfDelivery    `json:"placeOfDelivery,inline" bson:"placeOfDelivery,inline"`
	CreateAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

type Order struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Product           Product            `json:"product,inline" bson:"product,inline"`
	StatusOrder       string             `json:"productname,omitempty" bson:"productname,omitempty"`
	SellerNAme        string             `json:"productname,omitempty" bson:"productname,omitempty"`
	Tons              string             `json:"productname,omitempty" bson:"productname,omitempty"`
	BusinessSpecifics string             `json:"productname,omitempty" bson:"productname,omitempty"`
	HourO             string             `json:"productname,omitempty" bson:"productname,omitempty"`
	Particularidades  string             `json:"particularidades,omitempty" bson:"particularidades,omitempty"`
}
