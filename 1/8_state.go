package main

import "fmt"

type State interface {
	Next(order *Order)
	Previous(order *Order)
	Status() string
}

type Order struct {
	state State
}

func NewOrder() *Order {
	order := &Order{}
	order.setState(&PendingState{})
	return order
}

func (o *Order) setState(state State) {
	o.state = state
}

func (o *Order) Next() {
	o.state.Next(o)
}

func (o *Order) Previous() {
	o.state.Previous(o)
}

func (o *Order) Status() string {
	return o.state.Status()
}

type PendingState struct{}

func (p *PendingState) Next(order *Order) {
	fmt.Println("Order is now Shipped.")
	order.setState(&ShippedState{})
}

func (p *PendingState) Previous(order *Order) {
	fmt.Println("Order is already in the Pending state. No previous state.")
}

func (p *PendingState) Status() string {
	return "Pending"
}

type ShippedState struct{}

func (s *ShippedState) Next(order *Order) {
	fmt.Println("Order is now Delivered.")
	order.setState(&DeliveredState{})
}

func (s *ShippedState) Previous(order *Order) {
	fmt.Println("Order is now in Pending state.")
	order.setState(&PendingState{})
}

func (s *ShippedState) Status() string {
	return "Shipped"
}

type DeliveredState struct{}

func (d *DeliveredState) Next(order *Order) {
	fmt.Println("Order is already delivered. No next state.")
}

func (d *DeliveredState) Previous(order *Order) {
	fmt.Println("Order is now in Shipped state.")
	order.setState(&ShippedState{})
}

func (d *DeliveredState) Status() string {
	return "Delivered"
}

func main() {
	order := NewOrder()

	fmt.Println("Initial status:", order.Status()) // pending
	order.Next()                                   // ship
	fmt.Println("Current status:", order.Status()) // shipped
	order.Next()                                   // deliver
	fmt.Println("Current status:", order.Status()) // delivered

	order.Previous()                               // go back to shipped
	fmt.Println("Current status:", order.Status()) // shipped

	order.Previous()                               // go back to pending
	fmt.Println("Current status:", order.Status()) // pending
}
