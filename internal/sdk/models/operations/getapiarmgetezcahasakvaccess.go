// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type GetAPIARMGetEZCAHasAKVAccessRequest struct {
	AkvID *string `queryParam:"style=form,explode=true,name=akvID"`
}

func (o *GetAPIARMGetEZCAHasAKVAccessRequest) GetAkvID() *string {
	if o == nil {
		return nil
	}
	return o.AkvID
}

type GetAPIARMGetEZCAHasAKVAccessResponse struct {
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

func (o *GetAPIARMGetEZCAHasAKVAccessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPIARMGetEZCAHasAKVAccessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPIARMGetEZCAHasAKVAccessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPIARMGetEZCAHasAKVAccessResponse) GetAPIResultModel() *string {
	if o == nil {
		return nil
	}
	return o.APIResultModel
}

func (o *GetAPIARMGetEZCAHasAKVAccessResponse) GetAPIResultModel1() *shared.APIResultModel {
	if o == nil {
		return nil
	}
	return o.APIResultModel1
}
