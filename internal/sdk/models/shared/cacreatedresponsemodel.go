// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CACreatedResponseModel struct {
	Message  *string                       `json:"Message,omitempty"`
	Caid     *string                       `json:"CAID,omitempty"`
	LocalCAs []LocalCACreatedResponseModel `json:"LocalCAs,omitempty"`
	Success  *bool                         `json:"Success,omitempty"`
}

func (o *CACreatedResponseModel) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CACreatedResponseModel) GetCaid() *string {
	if o == nil {
		return nil
	}
	return o.Caid
}

func (o *CACreatedResponseModel) GetLocalCAs() []LocalCACreatedResponseModel {
	if o == nil {
		return nil
	}
	return o.LocalCAs
}

func (o *CACreatedResponseModel) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}
