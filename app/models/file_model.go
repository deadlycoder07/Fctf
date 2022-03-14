package models

type Files struct {
	Id           int    `pg:"id,pk"`
	Type         string `json:"type" pg:"type"`
	Location     string `json:"location" pg:"location"`
	ChallengesId int    `pg:"challenges_id"`
}
