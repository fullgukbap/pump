package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

const kakaoFetchUserInfo = "https://kapi.kakao.com/v2/user/me"

var conf *oauth2.Config

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Invalid config read %v", err)
	}

	conf = &oauth2.Config{
		ClientID:     viper.GetString("OAUTH_KAKAO_CLIENT_ID"),
		ClientSecret: viper.GetString("OAUTH_KAKAO_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://kauth.kakao.com/oauth/authorize",
			TokenURL: "https://kauth.kakao.com/oauth/token",
		},

		RedirectURL: "http://localhost:8080/auth/kakao/callback",
	}
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(1 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, cookie)
	return state
}

func kakaoLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := generateStateOauthCookie(w)
	url := conf.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func kakaoCallbackHandler(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthstate.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// get user info
	token, err := conf.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	_ = token

	resp := ""
	err = requests.
		URL(kakaoFetchUserInfo).
		Headers(map[string][]string{
			"Content-type":  {"application/x-www-form-urlencoded;charset=utf-8"},
			"Authorization": {fmt.Sprintf("Bearer %s", token.AccessToken)},
		}).
		ToString(&resp).
		Fetch(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("resp: ", resp)

}

func main() {
	http.HandleFunc("/login", kakaoLoginHandler)
	http.HandleFunc("/auth/kakao/callback", kakaoCallbackHandler)
	http.ListenAndServe(":8080", nil)
}
