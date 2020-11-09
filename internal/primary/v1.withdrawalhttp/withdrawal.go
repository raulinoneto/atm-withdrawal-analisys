package withdrawalhttp

import (
	"context"
	"errors"
	"github.com/raulinoneto/atm-withdrawal-analisys/internal/httpserver"
	"github.com/raulinoneto/atm-withdrawal-analisys/pkg/domain/v1.withdrawal"
	"github.com/raulinoneto/atm-withdrawal-analisys/tools/logger"
	"math"
	"net/http"
	"strconv"
)

type WithdrawalService interface {
	ProcessAmount(ctx context.Context, amount int) *withdrawal.ServiceResponse
}

type Adapter struct {
	svc WithdrawalService
}

func New(svc WithdrawalService) *Adapter {
	return &Adapter{svc: svc}
}

func (a *Adapter) WithdrawalHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	log := logger.New(ctx)
	log.Info("Started WithdrawalHttpAdapter::WithdrawalHandler")
	amount, err := validateAmount(r, log)
	if err != nil {
		httpErr := httpserver.Error{
			HttpStatus: http.StatusBadRequest,
			Message:    "Invalid amount param " + err.Error(),
			Code:       "0001",
		}
		return httpErr.ToHttpResponse(w)
	}
	return httpserver.BuildOkResponse(w, a.svc.ProcessAmount(ctx, amount))
}

func validateAmount(r *http.Request, log *logger.Logger) (int, error) {
	amtParam := r.FormValue("amount")
	amount, err := strconv.ParseFloat(amtParam, 64)
	if err != nil {
		err = errors.New("Invalid amount: " + err.Error())
		return 0, err
	}
	log.WithField("amount", amount)
	log.Info("Amount is a valid float")
	if amount <= 0 {
		log.Warn("User sent an invalid amount")
		return 0, errors.New(" Amount must be higher then zero ")
	}
	amtInt := math.Floor(amount)
	if amount-amtInt > 0 {
		return 0, errors.New(" Amount must be integer")
	}
	return int(amtInt), nil
}
