package app

import (
	"context"
	"io"

	"github.com/dingowd/hw12_13_14_15_calendar/internal/storage"
	"github.com/sirupsen/logrus"
)

type App struct {
	Event   storage.Event
	Logg    Logger
	Storage Storage
}

type Logger interface {
	SetLevel() logrus.Level
	SetOutput(output io.Writer)
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Warn(msg string)
}

type Storage interface {
	Connect(ctx context.Context, dsn string) error
	Close() error
	IsEventExist(e storage.Event) (bool, error)
	Create(e storage.Event) error
	Update(id int, e storage.Event) error
	Delete(id int) error
	GetIntervalEvent(day string, n int) ([]storage.Event, error)
	GetDayEvent(day string) ([]storage.Event, error)
	GetWeekEvent(day string) ([]storage.Event, error)
	GetMonthEvent(day string) ([]storage.Event, error)
}

func New(logger Logger, storage Storage) *App {
	a := new(App)
	a.Logg = logger
	a.Storage = storage
	return a
}

func (a *App) CreateEvent(ctx context.Context, e storage.Event) error {
	var err error
	if err = a.Storage.Create(e); err != nil {
		a.Logg.Error(err.Error())
	}
	return err
}

func (a *App) UpdateEvent(ctx context.Context, id int, e storage.Event) error {
	var err error
	if err = a.Storage.Update(id, e); err != nil {
		a.Logg.Error(err.Error())
	}
	return err
}

func (a *App) DeleteEvent(ctx context.Context, id int) error {
	var err error
	if err = a.Storage.Delete(id); err != nil {
		a.Logg.Error(err.Error())
	}
	return err
}

func (a *App) GetDayEvent(day string) ([]storage.Event, error) {
	var err error
	var events []storage.Event
	if events, err = a.Storage.GetDayEvent(day); err != nil {
		a.Logg.Error(err.Error())
	}
	return events, err
}

func (a *App) GetWeekEvent(day string) ([]storage.Event, error) {
	var err error
	var events []storage.Event
	if events, err = a.Storage.GetWeekEvent(day); err != nil {
		a.Logg.Error(err.Error())
	}
	return events, err
}

func (a *App) GetMonthEvent(day string) ([]storage.Event, error) {
	var err error
	var events []storage.Event
	if events, err = a.Storage.GetMonthEvent(day); err != nil {
		a.Logg.Error(err.Error())
	}
	return events, err
}
