// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type GetAPICATestTemplateConnectionRequest struct {
	CaID         *string `queryParam:"style=form,explode=true,name=caID"`
	TemplateName *string `queryParam:"style=form,explode=true,name=templateName"`
}

func (o *GetAPICATestTemplateConnectionRequest) GetCaID() *string {
	if o == nil {
		return nil
	}
	return o.CaID
}

func (o *GetAPICATestTemplateConnectionRequest) GetTemplateName() *string {
	if o == nil {
		return nil
	}
	return o.TemplateName
}

type GetAPICATestTemplateConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	APIResultModel *string
	// OK
	APIResultModel1 *shared.APIResultModel
}

func (o *GetAPICATestTemplateConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPICATestTemplateConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPICATestTemplateConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPICATestTemplateConnectionResponse) GetAPIResultModel() *string {
	if o == nil {
		return nil
	}
	return o.APIResultModel
}

func (o *GetAPICATestTemplateConnectionResponse) GetAPIResultModel1() *shared.APIResultModel {
	if o == nil {
		return nil
	}
	return o.APIResultModel1
}
