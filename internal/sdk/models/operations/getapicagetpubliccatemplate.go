// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type GetAPICAGetPublicCATemplateRequest struct {
	CaID *string `queryParam:"style=form,explode=true,name=caID"`
}

func (o *GetAPICAGetPublicCATemplateRequest) GetCaID() *string {
	if o == nil {
		return nil
	}
	return o.CaID
}

type GetAPICAGetPublicCATemplateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	PublicCATemplateDetails *string
	// OK
	PublicCATemplateDetails1 *shared.PublicCATemplateDetails
}

func (o *GetAPICAGetPublicCATemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPICAGetPublicCATemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPICAGetPublicCATemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPICAGetPublicCATemplateResponse) GetPublicCATemplateDetails() *string {
	if o == nil {
		return nil
	}
	return o.PublicCATemplateDetails
}

func (o *GetAPICAGetPublicCATemplateResponse) GetPublicCATemplateDetails1() *shared.PublicCATemplateDetails {
	if o == nil {
		return nil
	}
	return o.PublicCATemplateDetails1
}
