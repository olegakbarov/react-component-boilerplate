package db

import (
	"time"

	"github.com/guregu/null"
)

type PropertyMap map[string]interface{}

type Conf struct {
	Id                string      `json:"id"`
	Title             string      `json:"name"`
	Added_by          string      `json:"added_by"`
	Start_date        time.Time   `json:"start_date"`
	End_date          time.Time   `json:"end_date"`
	Description       string      `json:"description"`
	Picture           null.String `json:"picture"`
	Country           string      `json:"country"`
	City              string      `json:"city"`
	Address           string      `json:"address"`
	Category          string      `json:"category"`
	Tickets_available null.Bool   `json:"tickets_available"`
	Discount_program  null.Bool   `json:"discount_program"`
	Min_price         null.Int    `json:"min_price"`
	Max_price         null.Int    `json:"max_price"`
	Facebook          null.String `json:"facebook"`
	Youtube           null.String `json:"youtube"`
	Twitter           null.String `json:"twitter"`
	Details           PropertyMap `json:"details"`
}

// This stored in database and have private fields
type RawConf struct {
	Conf
	Verified   bool      `json:"verified"`
	Deleted    bool      `json:"deleted"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (r RawConf) PublicFields() Conf {
	return r.Conf
}

type User struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Locale    null.String `json:"locale"`
	City      null.String `json:"city"`
	Userpic   null.String `json:"userpic"`
	Email     string      `json:"email"`
	Verified  bool        `json:"verified"`
	Settings  PropertyMap `json:"settings"`
}
