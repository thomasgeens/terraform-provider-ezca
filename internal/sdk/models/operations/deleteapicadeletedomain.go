// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type DeleteAPICADeleteDomainRequest struct {
	CaID       *string `queryParam:"style=form,explode=true,name=caID"`
	Domain     *string `queryParam:"style=form,explode=true,name=domain"`
	TemplateID *string `queryParam:"style=form,explode=true,name=templateID"`
}

func (o *DeleteAPICADeleteDomainRequest) GetCaID() *string {
	if o == nil {
		return nil
	}
	return o.CaID
}

func (o *DeleteAPICADeleteDomainRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *DeleteAPICADeleteDomainRequest) GetTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

type DeleteAPICADeleteDomainResponse struct {
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

func (o *DeleteAPICADeleteDomainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAPICADeleteDomainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAPICADeleteDomainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteAPICADeleteDomainResponse) GetAPIResultModel() *string {
	if o == nil {
		return nil
	}
	return o.APIResultModel
}

func (o *DeleteAPICADeleteDomainResponse) GetAPIResultModel1() *shared.APIResultModel {
	if o == nil {
		return nil
	}
	return o.APIResultModel1
}
