package handlers

import (
	"dealls-test/dto"
	"dealls-test/pkg"
	"encoding/json"
	"net/http"
)

// GetAllPackages handles the request to get all premium packages
// @Summary Get all premium packages
// @Description Retrieve a list of all available premium packages
// @Tags Package
// @Accept json
// @Produce json
// @Success 200 {array} []entities.PremiumPackage "Successful response with list of premium packages"
// @Failure 500 {object} pkg.Response "Internal Server Error"
// @Router /packages [get]
func (h *Handler) GetAllPackage(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.Purchase.GetPremiumPackages(r.Context())
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	pkg.SendSuccessOkResp(w, res)
}

// PurchasePackage handles the request to purchase a premium package
// @Summary Purchase a premium package
// @Description Allows a user to purchase a specified premium package
// @Tags Package
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param purchaseReq body dto.PurchasePackageReq true "Purchase package request"
// @Success 201 {object} entities.Purchase "Successful response with purchase details"
// @Failure 400 {object} pkg.Response "Validation failed"
// @Failure 404 {object} pkg.Response "Package not found"
// @Failure 500 {object} pkg.Response "Internal Server Error"
// @Router /packages/purchase [post]
func (h *Handler) PurchasePackage(w http.ResponseWriter, r *http.Request) {
	var req dto.PurchasePackageReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.SendErrResp(w, pkg.ErrBadRequest)
		return
	}

	jwtPayload, err := pkg.GetJwtPayload(r.Context())
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	err = h.svc.Purchase.PurchasePackage(r.Context(), req, jwtPayload)
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	pkg.SendSuccessOkResp(w, nil)
}
