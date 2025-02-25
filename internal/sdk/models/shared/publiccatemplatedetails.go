// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PublicCATemplateDetails struct {
	ErrorMessage            *string                    `json:"ErrorMessage,omitempty"`
	TenantID                *string                    `json:"TenantID,omitempty"`
	Caid                    *string                    `json:"CAID,omitempty"`
	CAFriendlyName          *string                    `json:"CAFriendlyName,omitempty"`
	CATemplateType          *string                    `json:"CATemplateType,omitempty"`
	TemplateID              *string                    `json:"TemplateID,omitempty"`
	TemplateName            *string                    `json:"TemplateName,omitempty"`
	MaxCertLifeDays         *int                       `json:"MaxCertLifeDays,omitempty"`
	DualKeyRequired         *bool                      `json:"DualKeyRequired,omitempty"`
	AllowWildCardDomains    *bool                      `json:"AllowWildCardDomains,omitempty"`
	AllowAllOrgAsRequesters *bool                      `json:"AllowAllOrgAsRequesters,omitempty"`
	ApprovedAADRequesters   []AADObjectModel           `json:"ApprovedAADRequesters,omitempty"`
	Approvers               []AADObjectModel           `json:"Approvers,omitempty"`
	AllowedDomains          []PublicDomainResponseInfo `json:"AllowedDomains,omitempty"`
	AvailableDomains        []PublicDomainResponse     `json:"AvailableDomains,omitempty"`
}

func (o *PublicCATemplateDetails) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *PublicCATemplateDetails) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}

func (o *PublicCATemplateDetails) GetCaid() *string {
	if o == nil {
		return nil
	}
	return o.Caid
}

func (o *PublicCATemplateDetails) GetCAFriendlyName() *string {
	if o == nil {
		return nil
	}
	return o.CAFriendlyName
}

func (o *PublicCATemplateDetails) GetCATemplateType() *string {
	if o == nil {
		return nil
	}
	return o.CATemplateType
}

func (o *PublicCATemplateDetails) GetTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

func (o *PublicCATemplateDetails) GetTemplateName() *string {
	if o == nil {
		return nil
	}
	return o.TemplateName
}

func (o *PublicCATemplateDetails) GetMaxCertLifeDays() *int {
	if o == nil {
		return nil
	}
	return o.MaxCertLifeDays
}

func (o *PublicCATemplateDetails) GetDualKeyRequired() *bool {
	if o == nil {
		return nil
	}
	return o.DualKeyRequired
}

func (o *PublicCATemplateDetails) GetAllowWildCardDomains() *bool {
	if o == nil {
		return nil
	}
	return o.AllowWildCardDomains
}

func (o *PublicCATemplateDetails) GetAllowAllOrgAsRequesters() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAllOrgAsRequesters
}

func (o *PublicCATemplateDetails) GetApprovedAADRequesters() []AADObjectModel {
	if o == nil {
		return nil
	}
	return o.ApprovedAADRequesters
}

func (o *PublicCATemplateDetails) GetApprovers() []AADObjectModel {
	if o == nil {
		return nil
	}
	return o.Approvers
}

func (o *PublicCATemplateDetails) GetAllowedDomains() []PublicDomainResponseInfo {
	if o == nil {
		return nil
	}
	return o.AllowedDomains
}

func (o *PublicCATemplateDetails) GetAvailableDomains() []PublicDomainResponse {
	if o == nil {
		return nil
	}
	return o.AvailableDomains
}
