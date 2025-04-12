package main

import "fmt"

// Observer interface define o método para receber atualizações sobre o pedido
type Observer interface {
	Update(order *Order)
}

// EmailNotificationObserver representa um observador que envia e-mail
type EmailNotificationObserver struct{}

// NewEmailNotificationObserver cria uma nova instância de EmailNotificationObserver
func NewEmailNotificationObserver() *EmailNotificationObserver {
	return &EmailNotificationObserver{}
}

// Update implementa o método para enviar e-mail
func (o *EmailNotificationObserver) Update(order *Order) {
	fmt.Printf("Enviando e-mail para %s sobre o status do pedido %d: %s\n", order.CustomerEmail, order.ID, order.Status)
}

// WhatsappNotificationObserver representa um observador que envia mensagem via WhatsApp
type WhatsappNotificationObserver struct{}

// NewWhatsappNotificationObserver cria uma nova instância de WhatsappNotificationObserver
func NewWhatsappNotificationObserver() *WhatsappNotificationObserver {
	return &WhatsappNotificationObserver{}
}

// Update implementa o método para enviar mensagem via WhatsApp
func (o *WhatsappNotificationObserver) Update(order *Order) {
	fmt.Printf("Enviando WhatsApp para %s sobre o status do pedido %d: %s\n", order.CustomerPhone, order.ID, order.Status)
}

// SMSNotificationObserver representa um observador que envia SMS
type SMSNotificationObserver struct{}

// NewSMSNotificationObserver cria uma nova instância de SMSNotificationObserver
func NewSMSNotificationObserver() *SMSNotificationObserver {
	return &SMSNotificationObserver{}
}

// Update implementa o método para enviar SMS
func (o *SMSNotificationObserver) Update(order *Order) {
	fmt.Printf("Enviando SMS para %s sobre o status do pedido %d: %s\n", order.CustomerPhone, order.ID, order.Status)
}

// Order representa um pedido
type Order struct {
	ID            int
	Status        string
	CustomerEmail string
	CustomerPhone string
}

// NewOrder cria uma nova instância de Order
func NewOrder(id int, email string, phone string) *Order {
	return &Order{ID: id, CustomerEmail: email, CustomerPhone: phone, Status: "Pendente"}
}

// UpdateStatus atualiza o status do pedido
func (o *Order) UpdateStatus(newStatus string) {
	o.Status = newStatus
}

// OrderSubject interface define métodos para registrar, remover e notificar observadores de pedidos
type OrderSubject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// ConcreteOrderSubject representa o sistema de pedidos
type ConcreteOrderSubject struct {
	observers []Observer
	order     *Order
}

// NewConcreteOrderSubject cria uma nova instância de ConcreteOrderSubject
func NewConcreteOrderSubject(order *Order) *ConcreteOrderSubject {
	return &ConcreteOrderSubject{
		observers: make([]Observer, 0),
		order:     order,
	}
}

// RegisterObserver adiciona um observador à lista
func (s *ConcreteOrderSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver remove um observador da lista
func (s *ConcreteOrderSubject) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifica todos os observadores sobre a mudança no status do pedido
func (s *ConcreteOrderSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.order)
	}
}

// UpdateOrderStatus atualiza o status do pedido e notifica os observadores
func (s *ConcreteOrderSubject) UpdateOrderStatus(newStatus string) {
	fmt.Printf("Pedido %d teve seu status atualizado para: %s\n", s.order.ID, newStatus)
	s.order.UpdateStatus(newStatus)
	s.NotifyObservers()
}

func main() {
	pedido := NewOrder(123, "cliente@email.com", "551199999999")
	sistemaPedidos := NewConcreteOrderSubject(pedido)

	emailObserver := NewEmailNotificationObserver()
	whatsappObserver := NewWhatsappNotificationObserver()
	smsObserver := NewSMSNotificationObserver()

	sistemaPedidos.RegisterObserver(emailObserver)
	sistemaPedidos.RegisterObserver(whatsappObserver)

	sistemaPedidos.UpdateOrderStatus("Processando")

	fmt.Println("Removendo notificação por WhatsApp...")
	sistemaPedidos.RemoveObserver(whatsappObserver)

	sistemaPedidos.UpdateOrderStatus("Enviado")

	sistemaPedidos.RegisterObserver(smsObserver)

	sistemaPedidos.UpdateOrderStatus("Entregue")
}
