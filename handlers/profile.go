package handlers

import (
	"dealls-test/dto"
	"dealls-test/pkg"
	"encoding/json"
	"net/http"
)

// GetRandomProfile handles the request to get a random profile
// @Summary Get a random user profile
// @Description Retrieve a random user profile that has not been swiped by the current user
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} entities.Profile "Successful response with random profile"
// @Failure 404 {object} pkg.Response "No profile found"
// @Failure 500 {object} pkg.Response "Internal Server Error"
// @Router /profiles/random [get]
func (h *Handler) GetRandomProfile(w http.ResponseWriter, r *http.Request) {
	jwtPayload, err := pkg.GetJwtPayload(r.Context())
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	res, err := h.svc.Profile.View(r.Context(), jwtPayload)
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	pkg.SendSuccessOkResp(w, res)
}

// Swipe handles the request to swipe on a user profile
// @Summary Swipe on a user profile
// @Description Allows a user to swipe left or right on a specified profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param swipeReq body dto.SwipeReq true "Swipe request"
// @Success 200 {object} pkg.Response "Swipe recorded successfully"
// @Failure 400 {object} pkg.Response "Validation failed"
// @Failure 404 {object} pkg.Response "Profile not found"
// @Failure 500 {object} pkg.Response "Internal Server Error"
// @Router /profiles/swipe [post]
func (h *Handler) Swipe(w http.ResponseWriter, r *http.Request) {
	var req dto.SwipeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.SendErrResp(w, pkg.ErrBadRequest)
		return
	}

	jwtPayload, err := pkg.GetJwtPayload(r.Context())
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	err = h.svc.Profile.Swipe(r.Context(), req, jwtPayload)
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	pkg.SendSuccessOkResp(w, nil)
}
