package main

import "final-project/data"

// UserInterface is the interface for the user type. In order
// to satisfy this interface, all specified methods must be implemented.
// We do this so we can test things easily. Both data.User and data.UserTest
// implement this interface.
type UserInterface interface {
	GetAll() ([]*data.User, error)
	GetByEmail(email string) (*data.User, error)
	GetOne(id int) (*data.User, error)
	Update(user data.User) error
	// Delete() error
	DeleteByID(id int) error
	Insert(user data.User) (int, error)
	ResetPassword(password string) error
	PasswordMatches(plainText string) (bool, error)
}

// PlanInterface is the type for the plan type. Both data.Plan and data.PlanTest
// implement this interface.
type PlanInterface interface {
	GetAll() ([]*data.Plan, error)
	GetOne(id int) (*data.Plan, error)
	SubscribeUserToPlan(user data.User, plan data.Plan) error
	AmountForDisplay() string
}
