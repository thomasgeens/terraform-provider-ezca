// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetAPICAGetCertificatePrivateKeyRequest struct {
	Thumbprint *string `queryParam:"style=form,explode=true,name=thumbprint"`
}

func (o *GetAPICAGetCertificatePrivateKeyRequest) GetThumbprint() *string {
	if o == nil {
		return nil
	}
	return o.Thumbprint
}

type GetAPICAGetCertificatePrivateKeyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAPICAGetCertificatePrivateKeyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAPICAGetCertificatePrivateKeyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAPICAGetCertificatePrivateKeyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
