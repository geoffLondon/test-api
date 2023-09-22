//go:build unit

package handlers_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	mock_store "test-api/mocks/storage"

	"test-api/fixtures"
	"test-api/handlers"
	"test-api/model"
)

var _ = Describe("TestReadAllHandler unit fixtures", func() {
	var (
		controller *gomock.Controller
		mockStore  *mock_store.MockStore

		rbid *handlers.ReadByIDHandler
	)

	JustBeforeEach(func() {
		controller = gomock.NewController(GinkgoT())
		mockStore = mock_store.NewMockStore(controller)

		rbid = handlers.NewReadByIDHandler(mockStore)
	})
	AfterEach(func() {
		controller.Finish()
	})

	var req *http.Request

	BeforeEach(func() {
		req = httptest.NewRequest("GET", "/customer/foo", nil)
		req.Header.Add("Content-Type", "application/json")
	})
	Context("when a get customer query is made", func() {
		It("it returns a customer", func() {
			mockStore.EXPECT().ReadByID("foo").Return(fixtures.ValidCustomer(), nil)

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/customer/{id}", rbid.ReadByID)
			router.ServeHTTP(rr, req)

			status := rr.Code
			Expect(status).To(Equal(http.StatusOK))

			responseBody, err := io.ReadAll(rr.Body)
			Expect(err).NotTo(HaveOccurred())

			var returnedCustomer model.Customer
			err = json.Unmarshal(responseBody, &returnedCustomer)
			Expect(err).NotTo(HaveOccurred())

			expectedCustomer := fixtures.ValidCustomer()
			Expect(returnedCustomer).To(Equal(expectedCustomer))
		})
	})

	Context("when the ReadByID() fails", func() {
		It("it returns an error", func() {
			mockStore.EXPECT().ReadByID(gomock.Any()).Return(model.Customer{}, errors.New("foo"))

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/customer/{id}", rbid.ReadByID)
			router.ServeHTTP(rr, req)
			status := rr.Code

			Expect(status).To(Equal(http.StatusInternalServerError))
		})
	})
})
