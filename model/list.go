package model

import (
	"github.com/ChimeraCoder/anaconda"
)

func GetLists() ([]anaconda.List, error) {
	id, err := TellMyUserId()
	if err != nil {
		return nil, err
	}
	lists, err := api.GetListsOwnedBy(id, nil)
	if err != nil {
		return nil, err
	}
	return lists, nil
}
