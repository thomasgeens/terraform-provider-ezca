// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type OrderStatusModel struct {
	Success       *bool   `json:"Success,omitempty"`
	OrderStatus   *string `json:"OrderStatus,omitempty"`
	Certificate   *string `json:"Certificate,omitempty"`
	CACertificate *string `json:"CACertificate,omitempty"`
	ExtraInfo     *string `json:"ExtraInfo,omitempty"`
}

func (o *OrderStatusModel) GetSuccess() *bool {
	if o == nil {
		return nil
	}
	return o.Success
}

func (o *OrderStatusModel) GetOrderStatus() *string {
	if o == nil {
		return nil
	}
	return o.OrderStatus
}

func (o *OrderStatusModel) GetCertificate() *string {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *OrderStatusModel) GetCACertificate() *string {
	if o == nil {
		return nil
	}
	return o.CACertificate
}

func (o *OrderStatusModel) GetExtraInfo() *string {
	if o == nil {
		return nil
	}
	return o.ExtraInfo
}
