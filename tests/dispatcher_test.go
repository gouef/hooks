package tests

import (
	"fmt"
	"github.com/gouef/hooks"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockSubscriber je mock pro subscriber, který nám umožní testovat volání funkce
type MockSubscriber struct {
	mock.Mock
}

// OnUserCreated je metoda, kterou mockujeme pro testování
func (m *MockSubscriber) OnUserCreated(event interface{}) {
	m.Called(event) // Sledujeme, že byla zavolána
}

// UserCreatedEvent událost pro vytvoření uživatele
type UserCreatedEvent struct {
	UserID string
}

// GetName vrací název události
func (e *UserCreatedEvent) GetName() string {
	return "User::create"
}

type User struct {
	ID   string
	Name string
}

// Create metoda pro vytvoření uživatele
func (u *User) Create() {
	fmt.Printf("Uživatel %s vytvořen\n", u.Name)

	hooks.Trigger(&UserCreatedEvent{UserID: u.Name})
}

func TestDispatcher(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {

		mockSubscriber := new(MockSubscriber)

		hooks.AddHook("User::create", func(event interface{}) {
			mockSubscriber.OnUserCreated(event)
		})

		user := &User{Name: "Jan"}

		mockSubscriber.On("OnUserCreated", mock.Anything).Once()

		user.Create()

		mockSubscriber.AssertExpectations(t)

	})
}
