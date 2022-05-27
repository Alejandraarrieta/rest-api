package models

import (
	"go.mongodb.org/mongo-driver/bson"
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
	PriceId         string             `bson:"price_id,omitempty" json:"price_id,omitempty"`
	Harvest         string             `bson:"harvest,omitempty" json:"harvest,omitempty"`
	Bonus           string             `bson:"bonus,omitempty" json:"bonus,omitempty"`
	Reduction       string             `bson:"reduction,omitempty" json:"reduction,omitempty"`
	Pesification    string             `bson:"pesification,omitempty" json:"pesification,omitempty"`
	Observations    Observations       `json:"observations,inline" bson:"observations,inline"`
	PlaceOfDelivery PlaceOfDelivery    `json:"placeOfDelivery,inline" bson:"placeOfDelivery,inline"`
	CreateAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

type Order struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Product           Product            `json:"product,inline" bson:"product,inline"`
	StatusOrder       string             `json:"productname,omitempty" bson:"productname,omitempty"`
	SellerNAme        string             `json:"productname,omitempty" bson:"productname,omitempty"`
	Tons              string             `json:"productname,omitempty" bson:"productname,omitempty"`
	BusinessSpecifics string             `json:"productname,omitempty" bson:"productname,omitempty"`
	HourO             string             `json:"productname,omitempty" bson:"productname,omitempty"`
	Particularidades  string             `bson:"particularidades,omitempty" json:"particularidades,omitempty"`
}
