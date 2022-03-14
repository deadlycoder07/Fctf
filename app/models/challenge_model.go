package models

type Challenges struct {
	Id          int            `pg:"id,pk"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MaxAttempts int            `json:"maxAttempts"`
	Value       int            `json:"value"`
	Category    string         `json:"category"`
	Type        string         `json:"type"`
	State       bool           `json:"state"`
	Files       []*Files       `pg:"rel:has-many"`
	Deployments []*Deployments `pg:"rel:has-many"`
	Replicas    int            `json:"replicas"`
}
