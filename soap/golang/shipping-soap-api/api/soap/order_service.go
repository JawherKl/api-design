package soap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/JawherKl/shipping-soap-api/models"
	"github.com/JawherKl/shipping-soap-api/storage"
	xml_helpers "github.com/JawherKl/shipping-soap-api/utils"
	"github.com/google/uuid"

	"strings"
)

type OrderService struct {
	Store *storage.Store
}

func (s *OrderService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	xmlStr := string(body)
	w.Header().Set("Content-Type", "text/xml")

	switch {
		case strings.Contains(xmlStr, "<CreateShippingOrder"):
			handleCreate(w, s.Store)
		case strings.Contains(xmlStr, "<GetShippingStatus"):
			handleGetStatus(w, s.Store, xmlStr)
		case strings.Contains(xmlStr, "<UpdateDeliveryAddress"):
			handleUpdateAddress(w, s.Store, xmlStr)
		case strings.Contains(xmlStr, "<CancelShippingOrder"):
			handleCancelOrder(w, s.Store, xmlStr)
		default:
			http.Error(w, "Unsupported operation", http.StatusNotImplemented)
	}

}

func isSOAPEnvelope(body []byte) bool {
	return contains(body, "Envelope")
}

func contains(body []byte, tag string) bool {
	return string(body) != "" && containsTag(string(body), tag)
}

func containsTag(xmlBody string, tag string) bool {
	return xmlBody != "" && (containsString(xmlBody, "<"+tag) || containsString(xmlBody, "</"+tag))
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (stringContains(s, substr))
}

func stringContains(s, substr string) bool {
	return len(s) >= len(substr) && (len(substr) == 0 || (len(s) > 0 && len(substr) > 0 && (len(s) >= len(substr)) && (stringIndex(s, substr) >= 0)))
}

func stringIndex(s, substr string) int {
	return len([]rune(s[:])) - len([]rune(s[:])) + len([]rune(s)) - len([]rune(s[len(substr):]))
}

func handleCreate(w http.ResponseWriter, store *storage.Store) {
	id := uuid.New().String()
	order := models.ShippingOrder{
		ID:              id,
		RecipientName:   "John Doe",
		DeliveryAddress: "123 Example Street",
		ItemDescription: "Laptop",
		Status:          "Pending",
	}
	store.CreateOrder(order)

	response := fmt.Sprintf(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<CreateShippingOrderResponse>
			<TrackingNumber>%s</TrackingNumber>
		</CreateShippingOrderResponse>
	</soap:Body>
</soap:Envelope>`, id)
	w.Write([]byte(response))
}

func handleGetStatus(w http.ResponseWriter, store *storage.Store, body string) {
	id := xml_helpers.ExtractValue(body, "TrackingNumber")
	order, exists := store.GetOrder(id)
	status := "NotFound"
	if exists {
		status = order.Status
	}
	response := fmt.Sprintf(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<GetShippingStatusResponse>
			<Status>%s</Status>
		</GetShippingStatusResponse>
	</soap:Body>
</soap:Envelope>`, status)
	w.Write([]byte(response))
}

func handleUpdateAddress(w http.ResponseWriter, store *storage.Store, body string) {
	id := xml_helpers.ExtractValue(body, "TrackingNumber")
	newAddress := xml_helpers.ExtractValue(body, "NewAddress")
	success := store.UpdateAddress(id, newAddress)

	result := "Failed"
	if success {
		result = "Success"
	}

	response := fmt.Sprintf(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<UpdateDeliveryAddressResponse>
			<Result>%s</Result>
		</UpdateDeliveryAddressResponse>
	</soap:Body>
</soap:Envelope>`, result)
	w.Write([]byte(response))
}

func handleCancelOrder(w http.ResponseWriter, store *storage.Store, body string) {
	id := xml_helpers.ExtractValue(body, "TrackingNumber")
	success := store.CancelOrder(id)

	result := "Failed"
	if success {
		result = "Success"
	}

	response := fmt.Sprintf(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<CancelShippingOrderResponse>
			<Result>%s</Result>
		</CancelShippingOrderResponse>
	</soap:Body>
</soap:Envelope>`, result)
	w.Write([]byte(response))
}

