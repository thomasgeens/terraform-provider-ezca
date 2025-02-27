// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/models/shared"
	"net/http"
)

type GetAPICAGetAvailableRootsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TwoHundredTextPlainRes *string
	// OK
	TwoHundredApplicationJSONCAInformationModels []shared.CAInformationModel
	// OK
	TwoHundredTextJSONCAInformationModels []shared.CAInformationModel
}

func (o *GetAPICAGetAvailableRootsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPICAGetAvailableRootsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPICAGetAvailableRootsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAPICAGetAvailableRootsResponse) GetTwoHundredTextPlainRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoHundredTextPlainRes
}

func (o *GetAPICAGetAvailableRootsResponse) GetTwoHundredApplicationJSONCAInformationModels() []shared.CAInformationModel {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONCAInformationModels
}

func (o *GetAPICAGetAvailableRootsResponse) GetTwoHundredTextJSONCAInformationModels() []shared.CAInformationModel {
	if o == nil {
		return nil
	}
	return o.TwoHundredTextJSONCAInformationModels
}
