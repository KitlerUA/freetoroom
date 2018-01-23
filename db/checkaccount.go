package db

import "github.com/KitlerUA/freetoroom/hotel"

//CheckAccount - return true if user is registered, false otherwise
func (d *GormDB) CheckAccount(username, password string) (bool, error) {
	account := &hotel.Account{}
	if err := d.db.Where("username = ? and password = ?", username, password).First(account).Error;err!=nil{
		return false, err
	}
	return true, nil
}