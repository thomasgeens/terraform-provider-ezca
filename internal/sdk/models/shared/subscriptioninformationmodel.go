// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SubscriptionInformationModel struct {
	SubscriptionName *string    `json:"SubscriptionName,omitempty"`
	SubscriptionID   *string    `json:"SubscriptionID,omitempty"`
	CAPrice          *string    `json:"CAPrice,omitempty"`
	SupportType      *int       `json:"SupportType,omitempty"`
	Byok             *bool      `json:"BYOK,omitempty"`
	Byoca            *bool      `json:"BYOCA,omitempty"`
	Locations        []AKVModel `json:"Locations,omitempty"`
}

func (o *SubscriptionInformationModel) GetSubscriptionName() *string {
	if o == nil {
		return nil
	}
	return o.SubscriptionName
}

func (o *SubscriptionInformationModel) GetSubscriptionID() *string {
	if o == nil {
		return nil
	}
	return o.SubscriptionID
}

func (o *SubscriptionInformationModel) GetCAPrice() *string {
	if o == nil {
		return nil
	}
	return o.CAPrice
}

func (o *SubscriptionInformationModel) GetSupportType() *int {
	if o == nil {
		return nil
	}
	return o.SupportType
}

func (o *SubscriptionInformationModel) GetByok() *bool {
	if o == nil {
		return nil
	}
	return o.Byok
}

func (o *SubscriptionInformationModel) GetByoca() *bool {
	if o == nil {
		return nil
	}
	return o.Byoca
}

func (o *SubscriptionInformationModel) GetLocations() []AKVModel {
	if o == nil {
		return nil
	}
	return o.Locations
}
