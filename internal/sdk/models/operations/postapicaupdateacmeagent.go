// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type PostAPICAUpdateACMEAgentResponse struct {
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

func (o *PostAPICAUpdateACMEAgentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPICAUpdateACMEAgentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPICAUpdateACMEAgentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostAPICAUpdateACMEAgentResponse) GetAPIResultModel() *string {
	if o == nil {
		return nil
	}
	return o.APIResultModel
}

func (o *PostAPICAUpdateACMEAgentResponse) GetAPIResultModel1() *shared.APIResultModel {
	if o == nil {
		return nil
	}
	return o.APIResultModel1
}
