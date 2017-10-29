package errors

import (
	"encoding/json"
	"internauth-go/shared/types"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Error type encapsulates details about error
type Error struct {
	InternalCode int        `json:"code"`
	CodeGRPC     codes.Code `json:"-"`
	CodeHTTP     int        `json:"-"`
	Message      string     `json:"message"`
}

// Errors type encapsultates Err
type Errors struct {
	Errors []*Error `json:"errors"`
}

/* Main code used
 *
 * NotFound:          5   - http: 400
 * InvalidArgument:   3   - http: 400
 * Unauthenticated:   16  - http: 401
 * PermissionDenied:  7   - http: 403
 * AlreadyExists:     6   - http: 404
 * Internal:          13  - http: 500
 * Unimplemented:     12  - http: 501
 * Unavailable:       14  - http: 503
 */

// Default error xx
var (
	ErrInternalError      = &Error{1, codes.Internal, 500, "Internal error."}
	ErrNotFound           = &Error{2, codes.NotFound, 404, "Resource not found."}
	ErrServiceUnreachable = &Error{3, codes.Unavailable, 502, "Service unreachable."}
	ErrAccessDenied       = &Error{4, codes.Unauthenticated, 401, "Unauthenticated."}
	ErrAccessForbidden    = &Error{5, codes.PermissionDenied, 403, "Access forbidden."}
	ErrBadRequest         = &Error{6, codes.InvalidArgument, 400, "Bad request."}
	ErrNotImplemented     = &Error{7, codes.Unimplemented, 501, "Not implemented."}
	ErrDatabaseError      = &Error{8, codes.Internal, 500, "Database internal error."}
	ErrJWTOutDated        = &Error{9, codes.PermissionDenied, 403, "JWT out dated."}
	ErrJWTInvalid         = &Error{10, codes.PermissionDenied, 403, "JWT invalid."}
	ErrJWTMissing         = &Error{11, codes.PermissionDenied, 403, "JWT is missing."}
	ErrNotAuthenticated   = &Error{12, codes.Unauthenticated, 401, "You need to be authenticated."}
)

// WriteGRPC write the given error as a grpc error
func WriteGRPC(err *Error) error {
	return grpc.Errorf(err.CodeGRPC, err.Message, err.InternalCode)
}

// WriteHTTP write the given error as an HTTP response
func WriteHTTP(rw http.ResponseWriter, err *Error, stats *types.StatContext) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(err.CodeHTTP)
	stats.StatusCode = err.CodeHTTP
	json.NewEncoder(rw).Encode(Errors{[]*Error{err}})
}
