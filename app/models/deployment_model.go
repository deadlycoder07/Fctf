package models

type Deployments struct {
	Id           int    `pg:"id,pk"`
	ContainerId  int    `json:"containerId"`
	HealthStatus string `json:"healthStatus"`
	ChallengesId int    `pg:"challenges_id"`
}
