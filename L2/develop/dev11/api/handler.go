package api

import (
	"encoding/json"
	"httpServer/calendar"
	"log"
	"net/http"
	"time"
)

var CalendarExample = calendar.NewCalendar()

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
func initHeadersAndDecodeRequestBody(writer http.ResponseWriter, req *http.Request) (calendar.Event, error) {
	initHeaders(writer)

	var eventsToPost calendar.Event
	err := json.NewDecoder(req.Body).Decode(&eventsToPost)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return calendar.Event{}, err
	}

	return eventsToPost, nil
}
func CreateEventsHandler(writer http.ResponseWriter, req *http.Request) {
	eventsToPost, err := initHeadersAndDecodeRequestBody(writer, req)
	if err != nil {
		log.Fatal(err)
	}
	err = CalendarExample.CreateEvent(eventsToPost)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusServiceUnavailable,
			Message:    err.Error(),
			IsError:    true,
		}
		writer.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func UpdateEventsHandler(writer http.ResponseWriter, req *http.Request) {
	eventsToPost, err := initHeadersAndDecodeRequestBody(writer, req)
	if err != nil {
		log.Fatal(err)
	}
	err = CalendarExample.UpdateEvent(eventsToPost)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusServiceUnavailable,
			Message:    err.Error(),
			IsError:    true,
		}
		writer.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
func DeleteEventsHandler(writer http.ResponseWriter, req *http.Request) {
	eventsToPost, err := initHeadersAndDecodeRequestBody(writer, req)
	if err != nil {
		log.Fatal(err)
	}
	err = CalendarExample.DeleteEvent(eventsToPost.ID)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusServiceUnavailable,
			Message:    err.Error(),
			IsError:    true,
		}
		writer.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func GetEventsHandlerPerPeriod(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	// Получаем параметры из строки запроса
	startDateString := req.URL.Query().Get("start_date")
	endDateString := req.URL.Query().Get("end_date")

	startDate, err := time.Parse(time.RFC3339, startDateString)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid start date format",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateString)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid end date format",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	// Используем глобальный экземпляр календаря
	events := CalendarExample.GetEventsForPeriod(startDate, endDate)
	json.NewEncoder(writer).Encode(events)
}
func GetEventsHandlerPerDay(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	// Получаем параметры из строки запроса
	startDateString := req.URL.Query().Get("start_date")

	startDate, err := time.Parse(time.RFC3339, startDateString)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid start date format",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	events := CalendarExample.GetEventsForDay(startDate)
	json.NewEncoder(writer).Encode(events)
	return
}
func GetEventsHandlerPerMonth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	// Получаем параметры из строки запроса
	startDateString := req.URL.Query().Get("start_date")

	startDate, err := time.Parse(time.RFC3339, startDateString)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid start date format",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	events := CalendarExample.GetEventsForMonth(startDate)
	json.NewEncoder(writer).Encode(events)
	return
}
