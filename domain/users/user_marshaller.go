package users

type PublicUser struct {
	Id        int64  `json:"id"`
	Firstname string `json:"first_name"`
	//LastName  string `json:"last_name"`
	//Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	//Password    string `json:"password"`
}

func (users Users) Marshal(isPrivate bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshal(isPrivate)
	}
	return result

}

func (user *User) Marshal(isPrivate bool) interface{} {
	if !isPrivate {
		return PublicUser{
			Id:          user.Id,
			Firstname:   user.Firstname,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	} else {
		return PrivateUser{
			Id:          user.Id,
			Firstname:   user.Firstname,
			LastName:    user.LastName,
			Email:       user.Email,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
}
