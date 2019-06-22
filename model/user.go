package model

import (
	"github.com/ChimeraCoder/anaconda"
)

// 自分の情報を返す用の構造体
type MyUserDataToCheck struct {
	Login bool
	User  anaconda.User
}

// func TellMyId () (int, error) {
// 	id := api.
// }

func TellMe() (MyUserDataToCheck, error) {
	myDataToCheck := MyUserDataToCheck{}
	myData, err := api.GetSelf(nil)

	if err != nil {
		myDataToCheck.Login = false
		return myDataToCheck, err
	}

	myDataToCheck.Login = true
	myDataToCheck.User = myData
	return myDataToCheck, nil
}

func TellMyUserId() (int64, error) {
	myData, err := api.GetSelf(nil)
	if err != nil {
		return 0, err
	}

	return myData.Id, nil
}

func TellMyScreenName() (string, error) {
	myData, err := api.GetSelf(nil)
	if err != nil {
		return "", err
	}

	return myData.ScreenName, nil
}
