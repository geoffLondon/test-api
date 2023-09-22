//go:build unit

package handlers_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	mock_store "test-api/mocks/storage"

	"test-api/fixtures"
	"test-api/handlers"
)

var _ = Describe("TestReadAllHandler unit fixtures", func() {
	var (
		controller *gomock.Controller
		mockStore  *mock_store.MockStore

		rah *handlers.ReadHandler
	)

	JustBeforeEach(func() {
		controller = gomock.NewController(GinkgoT())
		mockStore = mock_store.NewMockStore(controller)

		rah = handlers.NewReadHandler(mockStore)
	})
	AfterEach(func() {
		controller.Finish()
	})

	Context("when a get all query is made", func() {
		It("it returns all customers", func() {
			mockStore.EXPECT().ReadAll().Return(fixtures.ValidCustomers(), nil)

			req := httptest.NewRequest("GET", "/customer", nil)
			req.Header.Add("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			readAllHandler := http.HandlerFunc(rah.Read)
			readAllHandler.ServeHTTP(rr, req)
			status := rr.Code

			Expect(status).To(Equal(http.StatusOK))
		})
	})
})
