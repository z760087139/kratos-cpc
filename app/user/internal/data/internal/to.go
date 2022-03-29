package internal

import "kratos-project/app/user/internal/biz"

func ToUser(user biz.User) User {
	// do someting
	return User{
		ID:       user.ID,
		Name:     user.Name,
		Password: "", // do someting
	}
}
