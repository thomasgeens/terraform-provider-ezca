// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type APIResultModel struct {
	Success *bool   `json:"Success,omitempty"`
	Message *string `json:"Message,omitempty"`
}

func (o *APIResultModel) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

func (o *APIResultModel) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
