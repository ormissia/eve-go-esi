package evegoesi

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"net/http"
	"sync"
)

const (
	authURL  = "https://login.eveonline.com/v2/oauth/authorize"
	tokenURL = "https://login.eveonline.com/v2/oauth/token"
)

var (
	errParam = errors.New("missing params")
)

type SSOAuthenticator struct {
	scopeLock   sync.Mutex
	HttpClient  *http.Client
	OauthConfig *oauth2.Config
	Token       *oauth2.Token
	Code        string

	Error error
}

func NewSSOAuthenticatorV2(clientID string, clientSecret string, redirectURL string, scopes []string) (sa *SSOAuthenticator, err error) {
	if clientID == "" || clientSecret == "" || redirectURL == "" {
		return nil, errParam
	}

	sa = &SSOAuthenticator{}
	sa.OauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     oauth2.Endpoint{AuthURL: authURL, TokenURL: tokenURL},
	}

	return sa, nil
}

func (sa *SSOAuthenticator) AuthorizeURL(state string, onlineAccess bool) (url string) {
	// Generate the URL
	if onlineAccess == true {
		url = sa.OauthConfig.AuthCodeURL(state, oauth2.AccessTypeOnline)
	} else {
		url = sa.OauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	}
	return
}

func (sa *SSOAuthenticator) InitWithRefreshToken(refreshToken string) (res *SSOAuthenticator) {
	res = sa
	token := &oauth2.Token{RefreshToken: refreshToken}

	tokenSource := res.OauthConfig.TokenSource(context.Background(), token)
	oauth2.ReuseTokenSource(token, tokenSource)
	res.HttpClient = oauth2.NewClient(context.Background(), tokenSource)
	res.Token = token

	return
}

func (sa *SSOAuthenticator) InitWithCode(code string) (res *SSOAuthenticator) {
	var err1, err2 error
	res = sa
	res.Code = code
	res.Token, err1 = res.OauthConfig.Exchange(context.Background(), res.Code)
	res.HttpClient = sa.OauthConfig.Client(context.Background(), sa.Token)
	res.Token, err2 = sa.OauthConfig.TokenSource(context.Background(), sa.Token).Token()
	if err1 != nil {
		res.Error = err1
	} else if err2 != nil {
		res.Error = err2
	}
	return
}
