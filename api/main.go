package main

import (
	"mini-mall/api/metric"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	router := chi.NewRouter()

	// Global middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{}))

	// Routes
	// router.Get("/", v1.IndexApi)

	// environ := vars.Environment

	env := "dev"
	if env == "dev" {
		// router.Mount("/debug", middleware.Profiler())
		router.Get("/debug/metrics", metric.MetricsApi)
	}

	router.Route("/api/v1", func(r1 chi.Router) {
		r1.Route("/auth", func(r2 chi.Router) {
			r2.Post("/token", func(w http.ResponseWriter, r *http.Request) {
				// Implement token generation logic here
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{ "token": "your_generated_token" }`))
			})
		})
	})

	/**
	User Authentication and Registration

	User registration
	Login with phone and verification code
	Login with phone and password
	Login with account and password
	Password reset


	User Management

	Get user information
	List users
	Update user address


	Merchant Management

	Submit merchant materials
	Apply for a shop
	Shop pledging (deposit)


	Product Management

	Add product to cart
	Remove product from cart
	List products in cart
	Put away products (for sellers)
	Supplement product properties
	List products (SKUs)


	Order Management

	Generate order code
	Create trade order
	Process order payment
	Get order report
	Get shop ranking by orders
	Get product ranking by orders


	Logistics

	Apply for logistics
	Query logistics record
	Update logistics record


	User Account

	Account recharge


	Comments and Ratings

	Modify comment tags
	Get shop comments list
	Create order comments
	Get comments tags list


	Search Functionality

	Search product inventory
	Search shops
	Search user information
	Search merchant information
	Search trade orders
	*/

	http.ListenAndServe(":8000", router)
}
