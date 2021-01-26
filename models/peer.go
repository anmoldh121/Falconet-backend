package models

type Peer struct {
	PeerId   string `bson:"_id,omitempty"`
	Ph       string `bson:"ph,omitempty"`
	Endpoint string `bson:"endpoint,omitempty"`
}
