package calendar

import (
	"errors"
	"time"
)

type Event struct {
	ID   string    `json:"ID"`
	Date time.Time `json:"date"`
}

type Calendar struct {
	events map[string]Event
}

func NewCalendar() Calendar {
	return Calendar{
		events: make(map[string]Event),
	}
}
func (c *Calendar) CreateEvent(e Event) error {
	if _, ok := c.events[e.ID]; ok {
		return errors.New("the event with this ID already exists")
	}
	c.events[e.ID] = e
	return nil
}
func (c *Calendar) UpdateEvent(e Event) error {
	if _, ok := c.events[e.ID]; !ok {
		return errors.New("event with this ID doesn't exist, please firstly create it")
	}
	c.events[e.ID] = e
	return nil
}
func (c *Calendar) DeleteEvent(ID string) error {
	if _, ok := c.events[ID]; !ok {
		return errors.New("event with this ID doesn't exist")
	}
	delete(c.events, ID)
	return nil
}
func (c *Calendar) GetEventsForPeriod(startDate time.Time, endDate time.Time) []Event {
	eventsPerPeriod := make([]Event, 0)
	if startDate.Equal(endDate) {
		for _, val := range c.events {
			if val.Date.Equal(endDate) {
				eventsPerPeriod = append(eventsPerPeriod, val)
			}
		}
	} else {
		for _, val := range c.events {
			if val.Date.Before(endDate) && val.Date.After(startDate) {
				eventsPerPeriod = append(eventsPerPeriod, val)
			}
		}
	}
	return eventsPerPeriod
}
func (c *Calendar) GetEventsForDay(day time.Time) []Event {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 0, 1)
	return c.GetEventsForPeriod(start, end)
}
func (c *Calendar) GetEventsForMonth(day time.Time) []Event {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, 0)
	return c.GetEventsForPeriod(start, end)
}
