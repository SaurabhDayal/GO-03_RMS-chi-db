package server

import (
	"03_RMS/handlers"
	"03_RMS/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// master
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewares.JsonMiddleware)

	// Login-Logout & add Address
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
	r.Route("/", func(r chi.Router) {
		r.Use(middlewares.VerifyUserMidd)
		r.Post("/logout", handlers.Logout)
	})

	// Admin routes
	r.Route("/admin", func(r chi.Router) {
		r.Use(middlewares.VerifyAdminMidd)
		r.Post("/subAdmin", handlers.CreateSubAdmin)
		r.Get("/subAdmin", handlers.GetSubAdminList)
	})

	// SubAdmin routes
	r.Route("/subAdmin", func(r chi.Router) {
		r.Use(middlewares.VerifySubAdminMidd)
		r.Post("/address", handlers.AddOwnerAddress)
		r.Post("/restaurant", handlers.CreateRestaurant)
		r.Post("/dish", handlers.CreateDish)
		r.Get("/restaurants", handlers.GetRestaurantByOwnerId)
		r.Get("/dishes", handlers.GetDishByOwnerId)
		r.Get("/users", handlers.GetUsersList)
	})

	// User Routes
	r.Route("/user", func(r chi.Router) {
		r.Use(middlewares.VerifyUserMidd)
		r.Post("/address", handlers.AddUserAddress)
		// Getting Info Routes
		r.Route("/restaurant", func(r chi.Router) {
			r.Get("/", handlers.GetRestaurantList)
			r.Get("/{resId}/dishes", handlers.GetDishList)
			// Get distance from RESTAURANT to a particular USER address
			r.Get("/{resId}/distance/{userAddId}", handlers.GetDistance)
		})
		// Order Routes
		r.Route("/order", func(r chi.Router) {
			r.Post("/{dishId}", handlers.AddOrder)
			r.Delete("/{orderId}", handlers.CancelOrder)
			r.Put("/{orderId}", handlers.OkOrder)
		})
	})

	return r
}
