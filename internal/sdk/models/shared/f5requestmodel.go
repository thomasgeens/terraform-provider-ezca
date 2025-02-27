// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type F5RequestModel struct {
	F5DeviceName *string `json:"F5DeviceName,omitempty"`
	Endpoint     *string `json:"Endpoint,omitempty"`
	Username     *string `json:"Username,omitempty"`
	Password     *string `json:"Password,omitempty"`
}

func (o *F5RequestModel) GetF5DeviceName() *string {
	if o == nil {
		return nil
	}
	return o.F5DeviceName
}

func (o *F5RequestModel) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *F5RequestModel) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *F5RequestModel) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}
