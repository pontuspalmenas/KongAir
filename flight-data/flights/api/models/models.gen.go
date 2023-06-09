// Package models provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package models

import (
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Flight defines model for Flight.
type Flight struct {
	Number string `json:"number"`
	Route  struct {
		Destination *string `json:"destination,omitempty"`
		Origin      *string `json:"origin,omitempty"`
	} `json:"route"`
	ScheduledArrival   time.Time `json:"scheduled_arrival"`
	ScheduledDeparture time.Time `json:"scheduled_departure"`
}

// FlightDetails defines model for FlightDetails.
type FlightDetails struct {
	AircraftType          string   `json:"aircraft_type,omitempty"`
	FlightNumber          string   `json:"flight_number,omitempty"`
	InFlightEntertainment bool     `json:"in_flight_entertainment,omitempty"`
	MealOptions           []string `json:"meal_options,omitempty"`
}

// GetFlightsParams defines parameters for GetFlights.
type GetFlightsParams struct {
	// Date Filter by date (defaults to current day)
	Date *openapi_types.Date `form:"date,omitempty" json:"date,omitempty"`
}
