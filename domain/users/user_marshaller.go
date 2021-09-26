package users

import "encoding/json"

type PublicUser struct {
	Id int64 `json:"id"`
	//FirstName   string `json:"first_name"`
	//LastName    string `json:"last_name"`
	//Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"`
}

func (u *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          u.Id,
			DateCreated: u.DateCreated,
			Status:      u.Status,
		}
	}

	//this approach to map is only used when you know the json are matching in both strucrts
	userJson, _ := json.Marshal(u)
	var privateuser PrivateUser
	json.Unmarshal(userJson, &privateuser)
	return privateuser
}
