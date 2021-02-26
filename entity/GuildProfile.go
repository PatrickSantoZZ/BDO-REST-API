package entity

import "time"

type GuildProfile struct {
	Name 		string 		`json:"name"`
	Region 		string 		`json:"region"`
	CreatedOn 	*time.Time 	`json:"createdOn,omitempty"`
	GuildMaster *Profile 	`json:"guildMaster,omitempty"`
	Members 	[]Profile 	`json:"members,omitempty"`
}