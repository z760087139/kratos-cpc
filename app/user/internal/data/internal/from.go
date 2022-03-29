package internal

import "kratos-project/app/user/internal/biz"

func FromUser(user User) biz.User {
	// do something
	return biz.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
