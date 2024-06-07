// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package public

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Broadcasts a raw transaction or transaction group to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "txid", runtime.ParamLocationPath, ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.POST(baseURL+"/v2/transactions", wrapper.RawTransaction, m...)
	router.GET(baseURL+"/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET(baseURL+"/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"8aePjk8nH/5/AAAA//9iOfQe9x8BAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
