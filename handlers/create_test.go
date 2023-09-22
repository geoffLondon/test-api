//go:build unit

package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	mock_store "api-with-interfaces/mocks/storage"
	mock_utils "api-with-interfaces/mocks/utils"

	"api-with-interfaces/fixtures"
	"api-with-interfaces/handlers"
)

var _ = Describe("TestCreateHandler unit fixtures", func() {
	var (
		controller *gomock.Controller
		mockStore  *mock_store.MockStore
		mockUUID   *mock_utils.MockUUIDGenerator

		ch *handlers.CreateHandler
	)

	JustBeforeEach(func() {
		controller = gomock.NewController(GinkgoT())
		mockStore = mock_store.NewMockStore(controller)
		mockUUID = mock_utils.NewMockUUIDGenerator(controller)

		ch = handlers.NewCreateHandler(mockStore, mockUUID)
	})
	AfterEach(func() {
		controller.Finish()
	})

	Context("when a customer fills in Create ISA form", func() {
		BeforeEach(func() {
			handlers.RootPath = "../"
		})
		It("it returns a created customer", func() {
			form := fixtures.ValidFormValues()

			mockUUID.EXPECT().New().Return("mocked-uuid")
			mockStore.EXPECT().Save(fixtures.ValidCustomer()).Return(nil)

			req := httptest.NewRequest("POST", "/customer", strings.NewReader(form.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			rr := httptest.NewRecorder()
			customerHandler := http.HandlerFunc(ch.CreateCustomer)
			customerHandler.ServeHTTP(rr, req)
			status := rr.Code

			Expect(status).To(Equal(http.StatusOK))
		})
	})

	Context("when the Save() fails", func() {
		BeforeEach(func() {
			handlers.RootPath = "../"
		})
		It("it returns an error", func() {
			form := fixtures.ValidFormValues()

			mockUUID.EXPECT().New().Return("mocked-uuid")

			mockStore.EXPECT().Save(gomock.Any()).Return(errors.New("foo"))

			req := httptest.NewRequest("POST", "/customer", strings.NewReader(form.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			rr := httptest.NewRecorder()
			customerHandler := http.HandlerFunc(ch.CreateCustomer)
			customerHandler.ServeHTTP(rr, req)
			status := rr.Code

			Expect(status).To(Equal(http.StatusInternalServerError))
		})
	})
})
