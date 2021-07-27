package Models

import "bootcamp-prac/day-4/Config"

func OrderProduct(ord *Order) (err error){
	if err =Config.DB.Model(&ord).Create(ord).Error; err!=nil{
		return err
	}
	return nil
}

func GetOrderByID(ord *Order, id string) (err error){
	if err =Config.DB.Model(&ord).Where("id = ?",id).First(ord).Error; err!=nil{
		return err
	}
	return nil
}
