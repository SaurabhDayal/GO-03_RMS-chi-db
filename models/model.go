package models

import "fmt"

const UserCredit = 1000
const SubAdminCredit = 2000

type Users struct {
	UserName     string `json:"userName" db:"user_name"`
	UserPassword string `json:"userPassword" db:"user_password"`
	UserEmail    string `json:"userEmail" db:"user_email"`
}

type UsersClient struct {
	Id           int    `json:"id" db:"id"`
	UserName     string `json:"userName" db:"user_name"`
	UserPassword string `json:"userPassword" db:"user_password"`
	UserEmail    string `json:"userEmail" db:"user_email"`
	Credit       int    `json:"credit" db:"credit"`
}

type UserRoles struct {
	UserId int    `json:"userId" db:"user_id"`
	Role   string `json:"role" db:"role"`
}

type Auths struct {
	UserId    int    `json:"userID" db:"user_id"`
	UserToken string `json:"userToken" db:"user_token"`
}

type Addresses struct {
	Id          int     `json:"id" db:"id"`
	AddressName string  `json:"addressName" db:"address_name"`
	AddressLat  float64 `json:"addressLat" db:"address_lat"`
	AddressLng  float64 `json:"addressLng" db:"address_lng"`
}

type Restaurants struct {
	Id                int    `json:"id" db:"id"`
	RestaurantName    string `json:"restaurantName" db:"restaurant_name"`
	RestaurantAddress int    `json:"restaurantAddress" db:"restaurant_address"`
	UserId            int    `json:"userId" db:"user_id"`
}

type Dishes struct {
	Id            int    `json:"id" db:"id"`
	DishName      string `json:"dishName" db:"dish_name"`
	DishCost      int    `json:"dishCost" db:"dish_cost" `
	RestaurantId  int    `json:"restaurantId" db:"restaurant_id"`
	UserId        int    `json:"userId" db:"user_id"`
	PreparingTime string `json:"preparingTime" db:"preparing_time"`
}

type AddressDistance struct {
	UserAddId int     `json:"userAddId"`
	RestId    int     `json:"restId"`
	Distance  float64 `json:"distance"`
}

type Orders struct {
	Id           int    `json:"orderId" db:"id"`
	DishId       int    `json:"dishId" db:"dish_id"`
	DeliveryTime string `json:"deliveryTimeTime" db:"delivery_time"`
	UserId       int    `json:"userId" db:"user_id"`
	IsDelivered  bool   `json:"isDelivered" db:"is_delivered"`
}

type CustomClientError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (e *CustomClientError) Error() string {
	return fmt.Sprintf("Error\nMessage- %v\nStatus code - %v", e.Message, e.StatusCode)
}

func (e *CustomClientError) SC() int {
	return e.StatusCode
}
