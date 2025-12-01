package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Marker struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Location GeoJSONPoint       `json:"location" bson:"location"`
}

type GeoJSONPoint struct {
	Type        string     `json:"type" bson:"type"`
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}

type Area struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Coordinates GeoJSONPolygon     `json:"coordinates" bson:"coordinates"`
}

type GeoJSONPolygon struct {
	Type        string         `json:"type" bson:"type"`
	Coordinates [][][]float64  `json:"coordinates" bson:"coordinates"`
}