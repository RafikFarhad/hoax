package api_v1_auth

import (
	"github.com/RafikFarhad/hoax/config"
	"github.com/RafikFarhad/hoax/database/model"
	_ "github.com/RafikFarhad/hoax/docs"
	"github.com/RafikFarhad/hoax/http/request"
	"github.com/RafikFarhad/hoax/http/response"
	"github.com/RafikFarhad/hoax/types"
	"github.com/dchest/uniuri"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const (
	TokenValidity = 12 * time.Hour
)

// Me godoc
// @Summary Me API
// @Description Get auth user info
// @Tags Auth
// @Accept json
// @Produce json
// @Security JWT
// @Success 200 {object} response.ApiResponse
// @Failure 401 {object} response.ApiResponse
// @Router /api/v1/me [get]
func Me(ctx *fiber.Ctx) error {
	userId := ctx.Locals("auth_data").(types.AuthData).UserId
	user, err := model.GetUserById(userId, "UserInfo")
	if err != nil {
		return response.WithError(ctx, err.Error(), 1, 401)
	}
	return response.WithSuccess(ctx, "", response.MakeUserResponse(user))
}

// Login godoc
// @Summary Login API
// @Description Login via username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body request.LoginRequest true "raw json body"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 422 {object} response.ApiResponse
// @Router /api/v1/login [post]
func Login(ctx *fiber.Ctx) error {
	loginRequest := &request.LoginRequest{}
	var err error
	// parse request
	if err = ctx.BodyParser(loginRequest); err != nil {
		return response.WithError(ctx, err.Error(), 1, 422)
	}
	// validate
	if validationErr := request.ValidateLoginRequest(loginRequest); validationErr != nil {
		return response.WithValidationError(ctx, validationErr)
	}
	// fetch user from db
	var user *model.User
	if user, err = model.GetUserByUsername(loginRequest.Username); err != nil {
		// TODO :: error may be db error
		return response.WithError(ctx, "Incorrect username and/or password")
	}
	// match password
	if matchPassword(user.Password, loginRequest.Password) {
		// token generate
		tokenString, expiry, err := generateToken(user)
		if err != nil {
			return response.WithError(ctx, "Token generation error")
		}
		return response.WithSuccess(ctx, "", response.MakeLoginResponse(tokenString, expiry))
	}
	return response.WithError(ctx, "Incorrect username and/or password")
}

func generateToken(user *model.User) (string, int64, error) {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)
	// session expiry time
	expiry := time.Now().Add(TokenValidity).Unix()
	// salt
	salt := uniuri.NewLen(8)
	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["id"] = strconv.Itoa(int(user.Id)) + "::" + user.Username + "::" + salt
	claims["exp"] = expiry
	// generate token string
	tokenString, err := token.SignedString([]byte(config.AppConfig.JwtSecret))
	return tokenString, expiry, err
}

func matchPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
