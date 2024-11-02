package handlers

import (
	"dealls-test/dto"
	"dealls-test/pkg"
	"encoding/json"
	"net/http"
)

// Register handles the user registration request
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags User
// @Accept json
// @Produce json
// @Param registerReq body dto.RegisterReq true "User registration request"
// @Success 201 {object} dto.RegisterRes "User registered successfully"
// @Failure 400 {object} pkg.Response "Validation failed"
// @Failure 500 {object} pkg.Response "Internal Server Error"
// @Router /register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.SendErrResp(w, pkg.ErrBadRequest)
		return
	}

	res, err := h.svc.User.Register(r.Context(), req)
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	pkg.SendSuccessCreatedResp(w, res)
}

// Login handles the user login request
// @Summary User login
// @Description Log in a user with their email and password
// @Tags User
// @Accept json
// @Produce json
// @Param loginReq body dto.LoginReq true "User login request"
// @Success 200 {object} dto.RegisterRes "Successful login response with access token"
// @Failure 400 {object} pkg.Response "Validation failed"
// @Failure 401 {object} pkg.Response "Unauthorized, invalid credentials"
// @Failure 500 {object} pkg.Response "Internal Server Error"
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.SendErrResp(w, pkg.ErrBadRequest)
		return
	}

	res, err := h.svc.User.Login(r.Context(), req)
	if err != nil {
		pkg.SendErrResp(w, err)
		return
	}

	pkg.SendSuccessOkResp(w, res)
}
