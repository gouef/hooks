<img align=right width="168" src="docs/gouef_logo.png">

# hooks
A simple event-driven package that allows you to register hooks (listeners) and subscribers for specific events. Inspired by Symfony's event system, this package provides an easy way to implement hooks and event subscribers.

[![Static Badge](https://img.shields.io/badge/Github-gouef%2Fhooks-blue?style=for-the-badge&logo=github&link=github.com%2Fgouef%2Fhooks)](https://github.com/gouef/hooks)

[![GoDoc](https://pkg.go.dev/badge/github.com/gouef/hooks.svg)](https://pkg.go.dev/github.com/gouef/hooks)
[![GitHub stars](https://img.shields.io/github/stars/gouef/hooks?style=social)](https://github.com/gouef/hooks/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouef/hooks)](https://goreportcard.com/report/github.com/gouef/hooks)
[![codecov](https://codecov.io/github/gouef/hooks/branch/main/graph/badge.svg?token=YUG8EMH6Q8)](https://codecov.io/github/gouef/hooks)

## Versions
![Stable Version](https://img.shields.io/github/v/release/gouef/hooks?label=Stable&labelColor=green)
![GitHub Release](https://img.shields.io/github/v/release/gouef/hooks?label=RC&include_prereleases&filter=*rc*&logoSize=diago)
![GitHub Release](https://img.shields.io/github/v/release/gouef/hooks?label=Beta&include_prereleases&filter=*beta*&logoSize=diago)


## Installation

To install the `hooks` package, use the following Go command:

```bash
go get -u github.com/gouef/hooks
```

## Features
- Register hooks to events using event names (e.g., User::create).
- Subscribers can listen for specific events and execute their corresponding handlers.
- Supports multiple hooks for a single event.
- Allows easy testing using mocks.

## Example Usage

### Define Your Event Interface
You need to define an event that implements the Event interface. This interface should have a `GetName() string` method to return the name of the event.

```go
package app

// UserCreatedEvent represents an event triggered when a user is created
type UserCreatedEvent struct {
    UserID string
}

// GetName returns the name of the event
func (e *UserCreatedEvent) GetName() string {
    return "User::create"
}
```

### Create a Subscriber
A subscriber will listen for specific events and handle them when triggered.

```go
package app

import "fmt"

// Subscriber listens for events and handles them
type Subscriber struct {
    Name string
}

// OnUserCreated is called when the User::create event is triggered
func (s *Subscriber) OnUserCreated(event interface{}) {
    if userEvent, ok := event.(*UserCreatedEvent); ok {
        fmt.Printf("%s: User %s was created\n", s.Name, userEvent.UserID)
    }
}
```

### Define a User with Event Triggering

```go
package app

import (
	"fmt"
	"github.com/gouef/hooks"
)

// User model
type User struct {
    Name string
}

// Create is a method that triggers the User::create event
func (u *User) Create() {
    fmt.Printf("User %s created\n", u.Name)

    // Trigger the event after user creation
    hooks.Trigger(&UserCreatedEvent{UserID: u.Name})
}

```

### Using the Event System
Now you can use the event system to add hooks and trigger events.

```go
package main

import (
    "fmt"
    "app"
    "github.com/gouef/hooks"
)

func main() {
    // Create a subscriber
    subscriber := &app.Subscriber{Name: "Subscriber 1"}

    // Register a hook for the User::create event
    hooks.AddHook("User::create", func(event interface{}) {
        subscriber.OnUserCreated(event)
    })

    // Create a user, which will trigger the event
    user := &app.User{Name: "John"}
    user.Create() // This will automatically trigger the callback
}

```


### Testing
You can write tests for the events and subscribers using the testify package. Here’s an example test for the user creation event.

```go
package main

import (
    "testing"
    "app"
    "github.com/stretchr/testify/mock"
    "github.com/gouef/hooks"
)

// MockSubscriber is a mock for the subscriber
type MockSubscriber struct {
    mock.Mock
}

// OnUserCreated is a mocked method for testing
func (m *MockSubscriber) OnUserCreated(event interface{}) {
    m.Called(event)
}

func TestUserCreationTriggersEvent(t *testing.T) {
    // Create a mock subscriber
    mockSubscriber := new(MockSubscriber)

    // Register a hook for the User::create event
    hooks.AddHook("User::create", func(event interface{}) {
        mockSubscriber.OnUserCreated(event)
    })

    // Create a user
    user := &app.User{Name: "John"}

    // Expect that the OnUserCreated method will be called once with any event
    mockSubscriber.On("OnUserCreated", mock.Anything).Once()

    // Trigger the Create method, which should trigger the event
    user.Create()

    // Assert that the expected method was called
    mockSubscriber.AssertExpectations(t)
}

```

## Contributing

Read [Contributing](CONTRIBUTING.md)

## Contributors

<div>
<span>
  <a href="https://github.com/actions-user"><img src="https://raw.githubusercontent.com/gouef/hooks/refs/heads/contributors-svg/.github/contributors/actions-user.svg" alt="actions-user" /></a>
</span>
<span>
  <a href="https://github.com/JanGalek"><img src="https://raw.githubusercontent.com/gouef/hooks/refs/heads/contributors-svg/.github/contributors/JanGalek.svg" alt="JanGalek" /></a>
</span>
</div>

## Join our Discord Community! 🎉

[![Discord](https://img.shields.io/discord/1334331501462163509?style=for-the-badge&logo=discord&logoColor=white&logoSize=auto&label=Community%20discord&labelColor=blue&link=https%3A%2F%2Fdiscord.gg%2FwjGqeWFnqK
)](https://discord.gg/wjGqeWFnqK)

Click above to join our community on Discord!
