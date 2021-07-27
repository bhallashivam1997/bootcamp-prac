//Models/Product.go
package Models
import (
	"bootcamp-prac/day-4/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//GetAllUsers Fetch all user data
func GetAllProducts(prod *[]Product) (err error) {
	if err = Config.DB.Find(prod).Error; err != nil {
		return err
	}
	return nil
}
//CreateUser ... Insert New data
func CreateProduct(prod *Product) (err error) {
	if err = Config.DB.Model(&prod).Create(prod).Error; err != nil {
		return err
	}
	return nil
}
//GetUserByID ... Fetch only one user by Id
func GetProductByID(prod *Product, id string) (err error) {
	if err = Config.DB.Model(&prod).Where("id = ?", id).First(prod).Error; err != nil {
		return err
	}
	return nil
}
//UpdateProduct ... Update user
func UpdateProduct(prod *Product, id string) (err error) {
	fmt.Println(prod)
	Config.DB.Save(prod)
	return nil
}
//DeleteUser ... Delete user
func DeleteProduct(prod *Product, id string) (err error) {
	Config.DB.Model(&prod).Where("id = ?", id).Delete(prod)
	return nil
}