package practice

import (
	"errors"
	"fmt"
)

var ErrDatabaseDown = errors.New("database is down")
var ErrDataEmpty = errors.New("data is empty")

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
}

func (emailNotifier *EmailNotifier) Send(message string) error {
	fmt.Println("Sending email")
	return nil
}

type SMSNotifier struct {
}

func (smsNotifier *SMSNotifier) Send(message string) error {
	fmt.Println("Sending SMS")
	return nil
}

func NotifyUser(n Notifier, message string) error {
	if err := n.Send(message); err != nil {
		return err
	}
	return nil
}

type Storage interface {
	Save(data string) error
}

type PostgresStorage struct{}

func (postgresStorage *PostgresStorage) Save(data string) error {
	if data == "" {
		return ErrDataEmpty
	}
	fmt.Println("Save to Postgres: ", data)
	return ErrDatabaseDown
}

type MongoStorage struct{}

func (mongoStorage *MongoStorage) Save(data string) error {
	if data == "" {
		return ErrDataEmpty
	}
	fmt.Println("Save to Mongo: ", data)
	return ErrDatabaseDown
}

type CachedStorage struct {
	Storage
}

func (cachedStorage *CachedStorage) Save(data string) error {
	fmt.Println("Checking cached...", data)
	err := cachedStorage.Storage.Save(data)
	if err != nil {
		return fmt.Errorf("save to cached: %w", err)
	}
	return nil
}
