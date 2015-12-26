package openid_test

import (
	"fmt"
	"net/http"

	"github.com/emanoelxavier/openid2go/openid"
)

func AuthenticatedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The user was authenticated!")
}

func UnauthenticatedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Reached without authentication!")
}

func AuthenticatedHandlerWithUser(u *openid.User, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The user was autthenticated! The token was issued by %v and the user is %+v.", u.Issuer, u)
}

func Example() {
	configuration, _ := openid.NewConfiguration(openid.ProvidersGetter(getProviders_googlePlayground))

	http.Handle("/user", openid.AuthenticateUser(configuration, openid.UserHandlerFunc(AuthenticatedHandlerWithUser)))
	http.Handle("/authn", openid.Authenticate(configuration, http.HandlerFunc(AuthenticatedHandler)))
	http.HandleFunc("/unauth", UnauthenticatedHandler)

	http.ListenAndServe(":5100", nil)
}

// getProviders returns the identity providers that will authenticate the users of the underlying service.
// A Provider is composed by its unique issuer and the collection of client IDs registered with the provider that
// are allowed to call this service.
// On this example Google OP is the provider of choice and the client ID used refers
// to the Google OAUTH Playground https://developers.google.com/oauthplayground
func getProviders_googlePlayground() ([]openid.Provider, error) {
	provider, err := openid.NewProvider("https://accounts.google.com", []string{"407408718192.apps.googleusercontent.com"})
	if err != nil {
		return nil, err
	}

	return []openid.Provider{provider}, nil
}