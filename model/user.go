package model

import (
	"github.com/ChimeraCoder/anaconda"
)

// 自分の情報を返す用の構造体
type MyUserDataToCheck struct {
	Login bool
	User  anaconda.User
}

func TellMe(testApi *anaconda.TwitterApi) (MyUserDataToCheck, error) {
	myDataToCheck := MyUserDataToCheck{}
	myData, err := testApi.GetSelf(nil)

	if err != nil {
		myDataToCheck.Login = false
		return myDataToCheck, err
	}

	myDataToCheck.Login = true
	myDataToCheck.User = myData
	return myDataToCheck, nil
}

func TellOtherUserData(api *anaconda.TwitterApi, userScreenName string) (anaconda.User, error) {
	userData, err := api.GetUsersShow(userScreenName, nil)
	if err != nil {
		return userData, err
	}

	return userData, nil
}
