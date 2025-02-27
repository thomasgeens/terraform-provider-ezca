// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type PostAPICARenewCAResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CACreatedResponseModel *string
	// OK
	CACreatedResponseModel1 *shared.CACreatedResponseModel
}

func (o *PostAPICARenewCAResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPICARenewCAResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPICARenewCAResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostAPICARenewCAResponse) GetCACreatedResponseModel() *string {
	if o == nil {
		return nil
	}
	return o.CACreatedResponseModel
}

func (o *PostAPICARenewCAResponse) GetCACreatedResponseModel1() *shared.CACreatedResponseModel {
	if o == nil {
		return nil
	}
	return o.CACreatedResponseModel1
}
