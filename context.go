package twtxt

import (
	"fmt"
	"net/http"

	"github.com/prologic/twtxt/session"
	log "github.com/sirupsen/logrus"
)

type Context struct {
	InstanceName string

	Username      string
	User          *User
	Authenticated bool

	Error   bool
	Message string

	Tweeter Tweeter
	Tweets  Tweets

	RegisterDisabled        bool
	RegisterDisabledMessage string
}

func NewContext(conf *Config, db Store, req *http.Request) *Context {
	ctx := &Context{
		InstanceName: conf.Name,

		RegisterDisabled: !conf.Register,
	}

	if sess := req.Context().Value("sesssion"); sess != nil {
		if username, ok := sess.(*session.Session).Get("username"); ok {
			ctx.Authenticated = true
			ctx.Username = username
		}
	}

	if ctx.Authenticated && ctx.Username != "" {
		ctx.Tweeter = Tweeter{
			Nick: ctx.Username,
			URL:  fmt.Sprintf("http://0.0.0.0:8000/u/%s", ctx.Username),
		}

		user, err := db.GetUser(ctx.Username)
		if err != nil {
			log.WithError(err).Warnf("error loading user object for %s", ctx.Username)
		}
		ctx.User = user
	}

	return ctx
}
