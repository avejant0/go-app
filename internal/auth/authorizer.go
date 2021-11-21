package auth

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/casbin/casbin"
)

func New(model, policy string) *Authorizer {
	enforcer := casbin.NewEnforcer(model, policy)

	return &Authorizer{
		enforcer: enforcer,
	}
}

type Authorizer struct {
	enforcer *casbin.Enforcer
}

func (a *Authorizer) Authorize(subject, object, action string) error {
	if !a.enforcer.Enforce(subject, object, action) {
		msg := fmt.Sprintf(
			"%s not permitted to %s to %s",
			subject,
			action,
			object,
		)

		st := status.New(codes.PermissionDenied, msg)
		return st.Err()
	}

	return nil
}
