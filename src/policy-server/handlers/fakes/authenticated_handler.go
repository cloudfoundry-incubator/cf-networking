// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"policy-server/handlers"
	"policy-server/uaa_client"
	"sync"

	"code.cloudfoundry.org/lager"
)

type AuthenticatedHandler struct {
	ServeHTTPStub        func(logger lager.Logger, response http.ResponseWriter, request *http.Request, tokenData uaa_client.CheckTokenResponse)
	serveHTTPMutex       sync.RWMutex
	serveHTTPArgsForCall []struct {
		logger    lager.Logger
		response  http.ResponseWriter
		request   *http.Request
		tokenData uaa_client.CheckTokenResponse
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AuthenticatedHandler) ServeHTTP(logger lager.Logger, response http.ResponseWriter, request *http.Request, tokenData uaa_client.CheckTokenResponse) {
	fake.serveHTTPMutex.Lock()
	fake.serveHTTPArgsForCall = append(fake.serveHTTPArgsForCall, struct {
		logger    lager.Logger
		response  http.ResponseWriter
		request   *http.Request
		tokenData uaa_client.CheckTokenResponse
	}{logger, response, request, tokenData})
	fake.recordInvocation("ServeHTTP", []interface{}{logger, response, request, tokenData})
	fake.serveHTTPMutex.Unlock()
	if fake.ServeHTTPStub != nil {
		fake.ServeHTTPStub(logger, response, request, tokenData)
	}
}

func (fake *AuthenticatedHandler) ServeHTTPCallCount() int {
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	return len(fake.serveHTTPArgsForCall)
}

func (fake *AuthenticatedHandler) ServeHTTPArgsForCall(i int) (lager.Logger, http.ResponseWriter, *http.Request, uaa_client.CheckTokenResponse) {
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	return fake.serveHTTPArgsForCall[i].logger, fake.serveHTTPArgsForCall[i].response, fake.serveHTTPArgsForCall[i].request, fake.serveHTTPArgsForCall[i].tokenData
}

func (fake *AuthenticatedHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AuthenticatedHandler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ handlers.AuthenticatedHandler = new(AuthenticatedHandler)