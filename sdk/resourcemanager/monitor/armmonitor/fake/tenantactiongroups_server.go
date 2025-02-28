//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"regexp"
)

// TenantActionGroupsServer is a fake server for instances of the armmonitor.TenantActionGroupsClient type.
type TenantActionGroupsServer struct {
	// CreateOrUpdate is the fake for method TenantActionGroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, actionGroup armmonitor.TenantActionGroupResource, options *armmonitor.TenantActionGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armmonitor.TenantActionGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TenantActionGroupsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, options *armmonitor.TenantActionGroupsClientDeleteOptions) (resp azfake.Responder[armmonitor.TenantActionGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TenantActionGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, options *armmonitor.TenantActionGroupsClientGetOptions) (resp azfake.Responder[armmonitor.TenantActionGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByManagementGroupIDPager is the fake for method TenantActionGroupsClient.NewListByManagementGroupIDPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagementGroupIDPager func(managementGroupID string, xmsClientTenantID string, options *armmonitor.TenantActionGroupsClientListByManagementGroupIDOptions) (resp azfake.PagerResponder[armmonitor.TenantActionGroupsClientListByManagementGroupIDResponse])

	// Update is the fake for method TenantActionGroupsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, managementGroupID string, tenantActionGroupName string, xmsClientTenantID string, tenantActionGroupPatch armmonitor.ActionGroupPatchBodyAutoGenerated, options *armmonitor.TenantActionGroupsClientUpdateOptions) (resp azfake.Responder[armmonitor.TenantActionGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTenantActionGroupsServerTransport creates a new instance of TenantActionGroupsServerTransport with the provided implementation.
// The returned TenantActionGroupsServerTransport instance is connected to an instance of armmonitor.TenantActionGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTenantActionGroupsServerTransport(srv *TenantActionGroupsServer) *TenantActionGroupsServerTransport {
	return &TenantActionGroupsServerTransport{
		srv:                             srv,
		newListByManagementGroupIDPager: newTracker[azfake.PagerResponder[armmonitor.TenantActionGroupsClientListByManagementGroupIDResponse]](),
	}
}

// TenantActionGroupsServerTransport connects instances of armmonitor.TenantActionGroupsClient to instances of TenantActionGroupsServer.
// Don't use this type directly, use NewTenantActionGroupsServerTransport instead.
type TenantActionGroupsServerTransport struct {
	srv                             *TenantActionGroupsServer
	newListByManagementGroupIDPager *tracker[azfake.PagerResponder[armmonitor.TenantActionGroupsClientListByManagementGroupIDResponse]]
}

// Do implements the policy.Transporter interface for TenantActionGroupsServerTransport.
func (t *TenantActionGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TenantActionGroupsClient.CreateOrUpdate":
		resp, err = t.dispatchCreateOrUpdate(req)
	case "TenantActionGroupsClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "TenantActionGroupsClient.Get":
		resp, err = t.dispatchGet(req)
	case "TenantActionGroupsClient.NewListByManagementGroupIDPager":
		resp, err = t.dispatchNewListByManagementGroupIDPager(req)
	case "TenantActionGroupsClient.Update":
		resp, err = t.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TenantActionGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/tenantActionGroups/(?P<tenantActionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.TenantActionGroupResource](req)
	if err != nil {
		return nil, err
	}
	managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	tenantActionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tenantActionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdate(req.Context(), managementGroupIDUnescaped, tenantActionGroupNameUnescaped, getHeaderValue(req.Header, "x-ms-client-tenant-id"), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TenantActionGroupResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TenantActionGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/tenantActionGroups/(?P<tenantActionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	tenantActionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tenantActionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), managementGroupIDUnescaped, tenantActionGroupNameUnescaped, getHeaderValue(req.Header, "x-ms-client-tenant-id"), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TenantActionGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/tenantActionGroups/(?P<tenantActionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	tenantActionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tenantActionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), managementGroupIDUnescaped, tenantActionGroupNameUnescaped, getHeaderValue(req.Header, "x-ms-client-tenant-id"), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TenantActionGroupResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TenantActionGroupsServerTransport) dispatchNewListByManagementGroupIDPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByManagementGroupIDPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagementGroupIDPager not implemented")}
	}
	newListByManagementGroupIDPager := t.newListByManagementGroupIDPager.get(req)
	if newListByManagementGroupIDPager == nil {
		const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/tenantActionGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListByManagementGroupIDPager(managementGroupIDUnescaped, getHeaderValue(req.Header, "x-ms-client-tenant-id"), nil)
		newListByManagementGroupIDPager = &resp
		t.newListByManagementGroupIDPager.add(req, newListByManagementGroupIDPager)
	}
	resp, err := server.PagerResponderNext(newListByManagementGroupIDPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByManagementGroupIDPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByManagementGroupIDPager) {
		t.newListByManagementGroupIDPager.remove(req)
	}
	return resp, nil
}

func (t *TenantActionGroupsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/providers/Microsoft.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Insights/tenantActionGroups/(?P<tenantActionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.ActionGroupPatchBodyAutoGenerated](req)
	if err != nil {
		return nil, err
	}
	managementGroupIDUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	tenantActionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tenantActionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Update(req.Context(), managementGroupIDUnescaped, tenantActionGroupNameUnescaped, getHeaderValue(req.Header, "x-ms-client-tenant-id"), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TenantActionGroupResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
