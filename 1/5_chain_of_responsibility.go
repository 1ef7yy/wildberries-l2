package main

import "fmt"

type Handler interface {
	Handle(request *TransferRequest) error
}

type TransferRequest struct {
	UserID  int
	Amount  float64
	Account string
}

type AuthenticationHandler struct {
	next Handler
}

func (h *AuthenticationHandler) Handle(request *TransferRequest) error {
	if request.UserID == 0 {
		return fmt.Errorf("пользователь не аутентифицирован")
	}
	if h.next != nil {
		return h.next.Handle(request)
	}
	return nil
}

type AuthorizationHandler struct {
	next Handler
}

func (h *AuthorizationHandler) Handle(request *TransferRequest) error {
	if request.Account == "" {
		return fmt.Errorf("недостаточно прав доступа")
	}
	if h.next != nil {
		return h.next.Handle(request)
	}
	return nil
}

type ValidationHandler struct {
	next Handler
}

func (h *ValidationHandler) Handle(request *TransferRequest) error {
	if request.Amount <= 0 {
		return fmt.Errorf("неверная сумма перевода")
	}
	if h.next != nil {
		return h.next.Handle(request)
	}
	return nil
}

type TransferHandler struct{}

func (h *TransferHandler) Handle(request *TransferRequest) error {
	fmt.Printf("Перевод %f на счет %s успешно выполнен\n", request.Amount, request.Account)
	return nil
}

// func main() {
// 	authHandler := &AuthenticationHandler{}
// 	authzHandler := &AuthorizationHandler{}
// 	validationHandler := &ValidationHandler{}
// 	transferHandler := &TransferHandler{}

// 	authHandler.next = authzHandler
// 	authzHandler.next = validationHandler
// 	validationHandler.next = transferHandler

// 	request := &TransferRequest{
// 		UserID:  1,
// 		Amount:  100,
// 		Account: "1234567890",
// 	}

// 	if err := authHandler.Handle(request); err != nil {
// 		fmt.Println(err)
// 	}
// }
