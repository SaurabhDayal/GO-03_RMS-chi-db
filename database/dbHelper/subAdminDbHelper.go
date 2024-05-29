package dbHelper

import (
	"03_RMS/database"
	"03_RMS/errorHandling"
	"03_RMS/models"
	"database/sql"
	"fmt"
	"time"
)

func CreateNewRestaurant(r *models.Restaurants, token string) (*models.Restaurants, error) {
	ownerId, err := CheckUserId(token)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	var rest models.Restaurants
	SQL := `INSERT INTO restaurants (restaurant_name, restaurant_address, user_id) 
			VALUES ($1,$2,$3) RETURNING id,restaurant_name,restaurant_address,user_id`
	err = database.RMS.Get(&rest, SQL, r.RestaurantName, r.RestaurantAddress, ownerId)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	return &rest, nil
}

func CreateNewDish(d *models.Dishes, token string) (*models.Dishes, error) {
	ownerId, err := CheckUserId(token)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	isCorrect, err := CheckOwnerRestaurantsIds(ownerId, d.RestaurantId)
	if err != nil || isCorrect == false {
		return nil, errorHandling.UnableToAccessDB()
	}
	timer := d.PreparingTime
	timevalue, err := time.Parse("03:04:05", timer)
	fmt.Println(timevalue)
	var dish models.Dishes
	SQL := `INSERT INTO dishes (dish_name, dish_cost, restaurant_id, user_id, preparing_time) 
			VALUES ($1,$2,$3,$4,$5) RETURNING id,dish_name,dish_cost,restaurant_id, user_id, preparing_time`
	err = database.RMS.Get(&dish, SQL, d.DishName, d.DishCost, d.RestaurantId, ownerId, d.PreparingTime)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	return &dish, nil
}

func GetMyRestaurants(token string) ([]models.Restaurants, error) {
	rests := make([]models.Restaurants, 0)
	id, err := CheckUserId(token)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	SQL := `SELECT id, restaurant_name, restaurant_address, user_id 
			FROM restaurants WHERE user_id = $1 AND archived_at IS NULL`
	err = database.RMS.Select(&rests, SQL, id)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	return rests, nil
}

func GetMyDishes(token string) ([]models.Dishes, error) {
	dishes := make([]models.Dishes, 0)
	id, err := CheckUserId(token)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	SQL := `SELECT id, dish_name, dish_cost, restaurant_id, user_id, preparing_time 
			FROM dishes WHERE user_id = $1 AND archived_at IS NULL`
	err = database.RMS.Select(&dishes, SQL, id)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	return dishes, nil
}

func CreateOwnerNewAddress(a models.Addresses) (*models.Addresses, error) {
	var add models.Addresses
	SQL := `INSERT INTO addresses (address_name, address_lat, address_lng) 
			VALUES ($1,$2,$3) RETURNING id, address_name, address_lat, address_lng`
	err := database.RMS.Get(&add, SQL, a.AddressName, a.AddressLat, a.AddressLng)
	if err != nil {
		return nil, errorHandling.UnableToAccessDB()
	}
	return &add, nil
}

func GetUsersLimitOffset(limit int, offset int) (*models.UsersListPag, error) {
	users := make([]models.UsersClient, 0)
	var count int
	var usrList models.UsersListPag
	SQL := `SELECT u.id, u.user_name, u.user_password, u.user_email, u.credit
			FROM users u INNER JOIN user_roles r ON u.id = r.user_id 
			WHERE r.role=$1 AND u.archived_at IS NULL LIMIT $2 OFFSET $3`
	err := database.RMS.Select(&users, SQL, "user", limit, offset)
	if err != nil && err != sql.ErrNoRows {
		return nil, errorHandling.UnableToAccessDB()
	} else if err == sql.ErrNoRows {
		return nil, errorHandling.NoContentDB()
	} else {
		SQL1 := `SELECT COUNT(*) FROM users WHERE archived_at IS NULL`
		err2 := database.RMS.Get(&count, SQL1)
		if err2 != nil {
			return nil, errorHandling.UnableToAccessDB()
		}
		usrList.Users = users
		usrList.TotalUserCount = count
	}
	return &usrList, nil
}
