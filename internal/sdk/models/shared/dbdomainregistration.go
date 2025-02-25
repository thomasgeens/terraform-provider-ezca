// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DBDomainRegistration struct {
	TenantID           *string `json:"TenantID,omitempty"`
	Caid               *string `json:"CAID,omitempty"`
	TemplateID         *string `json:"TemplateID,omitempty"`
	Domain             *string `json:"Domain,omitempty"`
	DomainID           *string `json:"DomainID,omitempty"`
	CAFriendlyName     *string `json:"CAFriendlyName,omitempty"`
	NotificationEmails *string `json:"NotificationEmails,omitempty"`
}

func (o *DBDomainRegistration) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}

func (o *DBDomainRegistration) GetCaid() *string {
	if o == nil {
		return nil
	}
	return o.Caid
}

func (o *DBDomainRegistration) GetTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

func (o *DBDomainRegistration) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *DBDomainRegistration) GetDomainID() *string {
	if o == nil {
		return nil
	}
	return o.DomainID
}

func (o *DBDomainRegistration) GetCAFriendlyName() *string {
	if o == nil {
		return nil
	}
	return o.CAFriendlyName
}

func (o *DBDomainRegistration) GetNotificationEmails() *string {
	if o == nil {
		return nil
	}
	return o.NotificationEmails
}
