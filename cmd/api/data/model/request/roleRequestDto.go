package request

type RoleRequestDto struct {
	Name        string   `json:"name" validation:"required"` // role name
	Permissions []string `json:"permissions"`                // permission ids
}
