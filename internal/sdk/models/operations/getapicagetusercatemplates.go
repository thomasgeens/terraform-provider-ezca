// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type GetAPICAGetUserCATemplatesRequest struct {
	CaID *string `queryParam:"style=form,explode=true,name=caID"`
}

func (o *GetAPICAGetUserCATemplatesRequest) GetCaID() *string {
	if o == nil {
		return nil
	}
	return o.CaID
}

type GetAPICAGetUserCATemplatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TwoHundredTextPlainRes *string
	// OK
	TwoHundredApplicationJSONTemplateDetailsModels []shared.TemplateDetailsModel
	// OK
	TwoHundredTextJSONTemplateDetailsModels []shared.TemplateDetailsModel
}

func (o *GetAPICAGetUserCATemplatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPICAGetUserCATemplatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPICAGetUserCATemplatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPICAGetUserCATemplatesResponse) GetTwoHundredTextPlainRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoHundredTextPlainRes
}

func (o *GetAPICAGetUserCATemplatesResponse) GetTwoHundredApplicationJSONTemplateDetailsModels() []shared.TemplateDetailsModel {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONTemplateDetailsModels
}

func (o *GetAPICAGetUserCATemplatesResponse) GetTwoHundredTextJSONTemplateDetailsModels() []shared.TemplateDetailsModel {
	if o == nil {
		return nil
	}
	return o.TwoHundredTextJSONTemplateDetailsModels
}
