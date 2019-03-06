package wmwebservice

import (
	"context"
	"io/ioutil"
	// "log"
	"net/http"
	// "time"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/auth0"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/events/uoevents"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/uo"
	// wmrepo "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/wm/repository"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/pkg/herror"
)

// func loadUser(ctx context.Context) (*uoevents.User, error) {
// 	// Load requester's user info from Auth0
// 	token, _ := auth0.LoadToken(ctx)
// 	userInfoRef, err := token.LoadUserInfo(ctx)
// 	if err != nil {
// 		return nil, err
// 	} else if userInfoRef == nil {
// 		err = herror.New("unable to load user info from Auth0")
// 		return nil, err
// 	}

// 	// Get the external ID
// 	externalID, _ := auth0.LoadSubject(ctx)

// 	// Read the user
// 	repo := wmrepo.Build(ctx)
// 	uoUserRef, err := repo.ReadUserByExternalID(ctx, externalID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return uoUserRef, nil
// }

func readBody(ctx context.Context, r *http.Request) ([]byte, error) {
	defer r.Body.Close()
	return ioutil.ReadAll(r.Body)
}

// func sessionExpiresAt(ctx context.Context) time.Time {
// 	cfg := uo.BuildConfig(ctx)

// 	// Calculate expiresAt time from JWT token
// 	expiresAt := time.Now().Add(cfg.DefaultSessionDuration).UTC()
// 	if token, ok := auth0.LoadToken(ctx); ok {
// 		if token.ExpiresAt() != nil {
// 			expiresAt = token.ExpiresAt().UTC()
// 			log.Printf(
// 				"using expiration time from authentication token: %s",
// 				expiresAt.String(),
// 			)
// 		}
// 	}

// 	return expiresAt
// }
