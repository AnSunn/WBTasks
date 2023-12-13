package service

import (
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/cache"
	"github.com/AnSunn/WBTasks/L0/StanSubscriber/internal/model"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
)

// Structure to display the order
type TemplateData struct {
	HasValue bool
	Data     model.Order
}

var (
	templates         = template.Must(template.ParseFiles("service/OrderPage.html"))
	infoLogWebServer  = log.New(os.Stdout, "INFO[WebServer] ", log.Ldate|log.Ltime)
	errorLogWebServer = log.New(os.Stderr, "ERROR[WebServer] ", log.Ldate|log.Ltime|log.Lshortfile)
)

// Handler for "/get"
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Check the method
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Take the orderUID from the form (static/OrderPage.html)
	orderUID := r.FormValue("orderUID")

	// Take the order from the cache
	order, ok := cache.GetItemFromCache(orderUID)

	// Preparing data to send to the template
	var templateData = TemplateData{
		HasValue: ok,
		Data:     order,
	}

	// Show HTML-page with populated data
	err := templates.ExecuteTemplate(w, "OrderPage.html", templateData)
	if err != nil {
		errorLogWebServer.Printf(err.Error())
		return
	}
}

// Web server provides order data by order_id
func StartWebServer(wg1 *sync.WaitGroup) {
	defer wg1.Done()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/get", indexHandler)

	// Launch web-server on 8181 port
	go func() {
		if err := http.ListenAndServe(":8181", nil); err != nil {
			errorLogWebServer.Printf("HTTP server error: %v\n", err)
		}
	}()
	infoLogWebServer.Println("Web-server is running")
}
