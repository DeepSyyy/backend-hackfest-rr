package response

import "time"

type ClanResponse struct {
	ClanTag     string    `json:"clan_tag"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
