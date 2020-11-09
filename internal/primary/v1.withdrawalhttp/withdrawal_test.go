package withdrawalhttp

import (
	"context"
	"encoding/json"
	withdrawal "github.com/raulinoneto/atm-withdrawal-analisys/pkg/domain/v1.withdrawal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type withdrawalServiceMock struct {
	mock.Mock
}

func (w *withdrawalServiceMock) ProcessAmount(ctx context.Context, amount int) *withdrawal.ServiceResponse {
	args := w.Called(ctx, amount)
	if res := args.Get(0); res != nil {
		return res.(*withdrawal.ServiceResponse)
	}
	return nil
}

func TestAdapter_WithdrawalHandler(t *testing.T) {
	svcmock := new(withdrawalServiceMock)
	svcmock.On("ProcessAmount", mock.Anything, 10).Once().Return(&withdrawal.ServiceResponse{})
	adapter := New(svcmock)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://test?amount=10.00", nil)
	err := adapter.WithdrawalHandler(w, r)
	assert.NoError(t, err)
	result, err := ioutil.ReadAll(w.Body)
	resSvc := new(withdrawal.ServiceResponse)
	err = json.Unmarshal(result,resSvc)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, *resSvc, withdrawal.ServiceResponse{})
}

func TestAdapter_WithdrawalHandlerErrorNoFloat(t *testing.T) {
	svcmock := new(withdrawalServiceMock)
	adapter := New(svcmock)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://test?amount=1a0", nil)
	err := adapter.WithdrawalHandler(w, r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAdapter_WithdrawalHandlerErrorEmpty(t *testing.T) {
	svcmock := new(withdrawalServiceMock)
	adapter := New(svcmock)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://test", nil)
	err := adapter.WithdrawalHandler(w, r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAdapter_WithdrawalHandlerErrorDecimal(t *testing.T) {
	svcmock := new(withdrawalServiceMock)
	adapter := New(svcmock)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://test?amount=2.20", nil)
	err := adapter.WithdrawalHandler(w, r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAdapter_WithdrawalHandlerErrorNegative(t *testing.T) {
	svcmock := new(withdrawalServiceMock)
	adapter := New(svcmock)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://test?amount=-2", nil)
	err := adapter.WithdrawalHandler(w, r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
