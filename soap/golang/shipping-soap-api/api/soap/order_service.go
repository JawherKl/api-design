package soap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/JawherKl/shipping-soap-api/models"
	"github.com/JawherKl/shipping-soap-api/storage"
	//xml "github.com/JawherKl/shipping-soap-api/utils"
	"github.com/google/uuid"

	"encoding/xml"
)

type OrderService struct {
	Store *storage.Store
}

func (s *OrderService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var envelope models.Envelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		http.Error(w, "Invalid XML", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/xml")

	switch {
		case envelope.Body.GetStatus != nil:
			handleGetStatus(w, s.Store, envelope.Body.GetStatus)
		case envelope.Body.UpdateAddress != nil:
			handleUpdateAddress(w, s.Store, envelope.Body.UpdateAddress)
		case envelope.Body.CancelShipping != nil:
			handleCancelOrder(w, s.Store, envelope.Body.CancelShipping)
		case envelope.Body.CreateOrder != nil:
			handleCreate(w, s.Store, envelope.Body.CreateOrder)
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

func handleCreate(w http.ResponseWriter, store *storage.Store, req *models.CreateShippingOrderRequest) {
	id := uuid.New().String()

	order := models.ShippingOrder{
		ID:              id,
		RecipientName:   req.RecipientName,
		DeliveryAddress: req.DeliveryAddress,
		ItemDescription: req.ItemDescription,
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

func handleGetStatus(w http.ResponseWriter, store *storage.Store, req *models.GetShippingStatusRequest) {
	order, exists := store.GetOrder(req.TrackingNumber)
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

func handleUpdateAddress(w http.ResponseWriter, store *storage.Store, req *models.UpdateDeliveryAddressRequest) {
	success := store.UpdateAddress(req.TrackingNumber, req.NewAddress)
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

func handleCancelOrder(w http.ResponseWriter, store *storage.Store, req *models.CancelShippingOrderRequest) {
	success := store.CancelOrder(req.TrackingNumber)
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

