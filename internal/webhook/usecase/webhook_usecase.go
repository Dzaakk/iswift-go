package webhook

import (
	"errors"
	"fmt"
	classRoomDto "iswift-go-project/internal/class_room/dto"
	classRoomUsecase "iswift-go-project/internal/class_room/usecase"
	orderDto "iswift-go-project/internal/order/dto"
	orderUsecase "iswift-go-project/internal/order/usecase"
	"os"
	"strings"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type WebhookUseCase interface {
	UpdatePayment(id string) error
}

type WebhookUseCaseImpl struct {
	orderUseCase     orderUsecase.OrderUseCase
	classRoomUseCase classRoomUsecase.ClassRoomUseCase
}

// UpdatePayment implements WebhookUseCase.
func (usecase *WebhookUseCaseImpl) UpdatePayment(id string) error {
	// kita akan melakukan pemeriksaan kembali ke xendit
	params := invoice.GetParams{
		ID: id,
	}

	dataXendit, err := invoice.Get(&params)

	if err != nil {
		return err
	}

	dataOrder, errorOrderUseCase := usecase.orderUseCase.FindByExternalId(dataXendit.ExternalID)

	if errorOrderUseCase != nil {
		return errorOrderUseCase
	}

	if dataOrder == nil {
		return errors.New("order is not found")
	}

	if dataOrder.Status == "settled" {
		return errors.New("payment has already processed")
	}

	if dataOrder.Status != "paid" {
		if dataXendit.Status == "PAID" || dataXendit.Status == "SETTLED" {
			// Add to class room
			for _, orderDetail := range dataOrder.OrderDetails {
				dataClassRoom := classRoomDto.ClassRoom{
					UserID:    dataOrder.UserID,
					ProductID: orderDetail.ProductID,
				}

				_, err := usecase.classRoomUseCase.Create(dataClassRoom)

				if err != nil {
					fmt.Println(err)
				}
			}

			// Mengirimkan Notifikasi melalui wa / email
		}
	}

	// Update data order
	orderDto := orderDto.OrderRequestBody{
		Status: strings.ToLower(dataXendit.Status),
	}

	usecase.orderUseCase.Update(int(dataOrder.ID), orderDto)

	return nil
}

func NewWebhookUseCase(
	orderUseCase orderUsecase.OrderUseCase,
	classRoomUseCase classRoomUsecase.ClassRoomUseCase,
) WebhookUseCase {
	// Setup Xendit
	xendit.Opt.SecretKey = os.Getenv("XENDIT_KEY")
	return &WebhookUseCaseImpl{
		orderUseCase,
		classRoomUseCase,
	}
}
