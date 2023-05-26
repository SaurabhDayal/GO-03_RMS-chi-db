package dbHelper

import (
	"06_RMS-chi-db/database"
	"06_RMS-chi-db/models"
	"fmt"
	"time"
)

func CreateNewRestaurant(r *models.Restaurants, token string) (*models.Restaurants, error) {
	ownerId, err := CheckUserId(token)
	if err != nil {
		return nil, err
	}
	SQL := `INSERT INTO restaurants (restaurant_name, restaurant_address, user_id) VALUES ($1,$2,$3)`
	_, err = database.RMS.Exec(SQL, r.RestaurantName, r.RestaurantAddress, ownerId)
	if err != nil {
		return nil, err
	}
	r.UserId = ownerId
	return r, nil
}

func CreateNewDish(d *models.Dishes, token string) (*models.Dishes, error) {
	ownerId, err := CheckUserId(token)
	if err != nil {
		return nil, err
	}
	isCorrect, err := CheckOwnerRestaurantsIds(ownerId, d.RestaurantId)
	if err != nil || isCorrect == false {
		return nil, err
	}

	timer := d.PreparingTime
	timevalue, err := time.Parse("03:04:05", timer)
	fmt.Println(timevalue)

	SQL := `INSERT INTO dishes (dish_name, dish_cost, restaurant_id, user_id, preparing_time) VALUES ($1,$2,$3,$4,$5)`
	_, err = database.RMS.Exec(SQL, d.DishName, d.DishCost, d.RestaurantId, ownerId, d.PreparingTime)
	if err != nil {
		return nil, err
	}
	d.UserId = ownerId
	return d, nil
}

func GetMyRestaurants(token string) ([]models.Restaurants, error) {
	rests := make([]models.Restaurants, 0)
	id, err := CheckUserId(token)
	if err != nil {
		return nil, err
	}
	SQL := `SELECT restaurant_name, restaurant_address, user_id FROM restaurants WHERE user_id = $1 AND archived_at IS NULL`
	err = database.RMS.Select(&rests, SQL, id)
	if err != nil {
		return nil, err
	}
	return rests, nil
}

func GetMyDishes(token string) ([]models.Dishes, error) {
	dishes := make([]models.Dishes, 0)
	id, err := CheckUserId(token)
	if err != nil {
		return nil, err
	}
	SQL := `SELECT dish_name, dish_cost, restaurant_id, user_id, preparing_time FROM dishes WHERE user_id = $1 AND archived_at IS NULL`
	err = database.RMS.Select(&dishes, SQL, id)
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func CreateOwnerNewAddress(a *models.Addresses) (*models.Addresses, error) {
	SQL := `INSERT INTO addresses (address_name, address_lat, address_lng) VALUES ($1,$2,$3)`
	_, err := database.RMS.Exec(SQL, a.AddressName, a.AddressLat, a.AddressLng)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func GetUsersLimitOffset(limit int, offset int) ([]models.Users, error) {
	users := make([]models.Users, 0)
	SQL := `SELECT u.user_name, u.user_password, u.user_email, u.credit FROM users u INNER JOIN user_roles r ON u.id = r.user_id WHERE r.role=$1 AND u.archived_at IS NULL LIMIT $2 OFFSET $3`
	err := database.RMS.Select(&users, SQL, "user", limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}
