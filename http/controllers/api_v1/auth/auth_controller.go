package api_v1_auth

import (
	"github.com/RafikFarhad/hoax/app"
	"github.com/RafikFarhad/hoax/database/model"
	"github.com/RafikFarhad/hoax/http/response"
	"github.com/RafikFarhad/hoax/http/validator/api"
	"github.com/RafikFarhad/hoax/types"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Me(ctx *fiber.Ctx) error {
	userId := ctx.Locals("authData").(types.AuthData).UserId
	user, err := model.GetUserById(userId, "UserInfo")
	if err != nil {
		return response.ErrorMessageWithStatus(ctx, err.Error(), 401)
	}
	return response.SuccessMessageWithData(ctx, user.WithUserInfo())
}

func Login(ctx *fiber.Ctx) error {
	loginRequest := &api_validator.LoginRequest{}
	var err error
	// parse request
	if err = ctx.BodyParser(loginRequest); err != nil {
		return response.ErrorMessageWithStatus(ctx, err.Error(), 422)
	}
	// validate
	if validationErr := api_validator.ValidateLoginRequest(loginRequest); validationErr != nil {
		return response.ValidationError(ctx, validationErr)
	}
	// fetch user from db
	var user *model.User
	if user, err = model.GetUserByUsername(loginRequest.Username); err != nil {
		return response.ErrorMessage(ctx, "Incorrect username and/or password")
	}
	// match password
	if matchPassword(user.Password, loginRequest.Password) {
		// token generate
		tokenString, expiry, err := generateToken(user)
		if err != nil {
			return response.ErrorMessage(ctx, "Token generation error")
		}
		return response.SuccessMessageWithData(ctx, map[string]interface{}{
			"token":  tokenString,
			"expiry": expiry,
		})
	}
	return response.ErrorMessage(ctx, "Incorrect username and/or password")
}

func generateToken(user *model.User) (string, int64, error) {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)
	// session expiry time
	expiry := getExpiryTime()
	// set claims
	claims := token.Claims.(jwt.MapClaims)
	//claims["user_id"] = user.Id
	claims["userId"] = user.Id
	//claims["user"], _ = user.MarshalJSON()
	claims["exp"] = expiry
	// generate token string
	tokenString, err := token.SignedString([]byte(app.App.Config.JwtSecret))
	return tokenString, expiry, err
}

func getExpiryTime() int64 {
	return time.Now().Add(24 * time.Hour).Unix()
}

func matchPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
