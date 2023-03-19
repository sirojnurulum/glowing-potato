package orders

import (
	"encoding/json"
	"glowing-potato/infra/context/service"
	"glowing-potato/objects"
	"glowing-potato/utils"
	"gopkg.in/validator.v2"
	"net/http"
)

type ordersHandler struct {
	*service.ServiceCtx
}

func (o ordersHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req objects.CreateOrder
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res := utils.SetResponseJSON(http.StatusInternalServerError, nil, utils.StatusInternalErr, err.Error())
		res.JSONResponse(w)
		return
	}

	if err = validator.Validate(req); err != nil {
		res := utils.SetResponseJSON(http.StatusBadRequest, nil, utils.StatusBadRequest, err.Error())
		res.JSONResponse(w)
		return
	}

	err = o.OrdersService.CreateOrder(r.Context(), req)
	if err != nil {
		res := utils.SetResponseJSON(http.StatusInternalServerError, nil, utils.StatusInternalErr, err.Error())
		res.JSONResponse(w)
		return
	}
	res := utils.SetResponseJSON(http.StatusCreated, "", "", "Success Create Order")
	res.JSONResponse(w)
	return
}
