// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type PostAPICAGetTemplateCertsResponse struct {
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

func (o *PostAPICAGetTemplateCertsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostAPICAGetTemplateCertsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostAPICAGetTemplateCertsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostAPICAGetTemplateCertsResponse) GetAPIResultModel() *string {
	if o == nil {
		return nil
	}
	return o.APIResultModel
}

func (o *PostAPICAGetTemplateCertsResponse) GetAPIResultModel1() *shared.APIResultModel {
	if o == nil {
		return nil
	}
	return o.APIResultModel1
}
