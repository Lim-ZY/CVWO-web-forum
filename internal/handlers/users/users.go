package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lim-zy/CVWO-web-forum/internal/api"
	pg "github.com/lim-zy/CVWO-web-forum/internal/database"
	model "github.com/lim-zy/CVWO-web-forum/internal/models"
)

type UserHandler struct {
  DB            *pg.Database
  CookieSecret  string
  CookieName    string
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request, username string) error {
  sg, err := time.LoadLocation("Asia/Singapore")
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return fmt.Errorf("Failed to load timezone: %v\n", err)
  }
  creationTime := time.Now().In(sg)

  user := model.User{
    Username:     username,
    CreationTime: creationTime,
    LastActive:   creationTime,
  }

  err = h.DB.CreateUser(&user)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return fmt.Errorf("Failed to create user: %v\n", err)
  }
  
  return nil
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
  var body struct {
    Username  string  `json:"username"`
  }

  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  var user model.User = model.User{Username: body.Username}
  if err := h.DB.GetUser(&user); err != nil {
    if err2 := h.Create(w, r, body.Username); err2 != nil {
      return
    }
  }

  sg, err := time.LoadLocation("Asia/Singapore")
  if err != nil {
    fmt.Printf("Failed to load timezone: %v\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  creationTime := time.Now().In(sg)

  claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
    Issuer:     user.Username,
    IssuedAt:   jwt.NewNumericDate(creationTime),
    ExpiresAt:  jwt.NewNumericDate(creationTime.Add(time.Hour * 24)),
  })

  if h.CookieSecret == "" {
    fmt.Println("JWT_SECRET env var not set")
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  token, err := claims.SignedString([]byte(h.CookieSecret))
  if err != nil {
    fmt.Println("Failed to sign token")
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  if h.CookieName == "" {
    fmt.Println("COOKIE_NAME env var not set")
    return
  }
  cookie := http.Cookie{
    Name: h.CookieName,
    Value: token,
    Expires: creationTime.Add(time.Hour * 24),
    HttpOnly: true,
    SameSite: http.SameSiteLaxMode,
  }
  http.SetCookie(w, &cookie)

  response := &api.Response{
		Messages: []string{fmt.Sprintf("User %s's token issued", user.Username)},
	}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
  if h.CookieName == "" || h.CookieSecret == "" {
    fmt.Println("COOKIE_NAME and JWT_SECRET env var not set")
    return
  }

  cookie, err := r.Cookie(h.CookieName); 
  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    response := &api.Response{
      Messages: []string{"Unauthorised: User not signed in"},
	  }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    return
  }

  token, err := jwt.ParseWithClaims(cookie.Value, &jwt.RegisteredClaims{}, 
                                    func(token *jwt.Token) (any, error) {
    return []byte(h.CookieSecret), nil
  })

  if err != nil || !token.Valid {
    w.WriteHeader(http.StatusUnauthorized)
    response := &api.Response{
      Messages: []string{"Unauthorised: Invalid token"},
	  }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    return
  }

  claims := token.Claims.(*jwt.RegisteredClaims)
  user := model.User{Username: claims.Issuer}
  if err := h.DB.GetUser(&user); err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    response := &api.Response{
      Messages: []string{"Unauthorised: Invalid user"},
	  }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    return
  }

  res, err := json.Marshal(user)
  if err != nil {
    fmt.Println("Failed to marshal: %w\n", err)
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  response := &api.Response{
    Payload: api.Payload{
      Data: res,
    },
    Messages: []string{fmt.Sprintf("Get user %s success", user.Username)},
	}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
  cookie := http.Cookie{
    Name: h.CookieName,
    Value: "",
    Expires: time.Now().Add(-time.Hour),
    HttpOnly: true,
  }
  http.SetCookie(w, &cookie)
  response := &api.Response{
    Messages: []string{"User logout (delete cookie) success"},
	}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}
