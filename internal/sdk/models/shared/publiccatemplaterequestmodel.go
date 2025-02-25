// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type PublicCATemplateRequestModel struct {
	DualKeyRequired         *bool                  `json:"DualKeyRequired,omitempty"`
	Approvers               []AADObjectModel       `json:"Approvers,omitempty"`
	AllowWildCardDomains    *bool                  `json:"AllowWildCardDomains,omitempty"`
	AllowAllOrgAsRequesters *bool                  `json:"AllowAllOrgAsRequesters,omitempty"`
	ApprovedAADRequesters   []AADObjectModel       `json:"ApprovedAADRequesters,omitempty"`
	MaxCertLifeDays         *int                   `json:"MaxCertLifeDays,omitempty"`
	AllowedDomains          []PublicDomainResponse `json:"AllowedDomains,omitempty"`
}

func (o *PublicCATemplateRequestModel) GetDualKeyRequired() *bool {
	if o == nil {
		return nil
	}
	return o.DualKeyRequired
}

func (o *PublicCATemplateRequestModel) GetApprovers() []AADObjectModel {
	if o == nil {
		return nil
	}
	return o.Approvers
}

func (o *PublicCATemplateRequestModel) GetAllowWildCardDomains() *bool {
	if o == nil {
		return nil
	}
	return o.AllowWildCardDomains
}

func (o *PublicCATemplateRequestModel) GetAllowAllOrgAsRequesters() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAllOrgAsRequesters
}

func (o *PublicCATemplateRequestModel) GetApprovedAADRequesters() []AADObjectModel {
	if o == nil {
		return nil
	}
	return o.ApprovedAADRequesters
}

func (o *PublicCATemplateRequestModel) GetMaxCertLifeDays() *int {
	if o == nil {
		return nil
	}
	return o.MaxCertLifeDays
}

func (o *PublicCATemplateRequestModel) GetAllowedDomains() []PublicDomainResponse {
	if o == nil {
		return nil
	}
	return o.AllowedDomains
}
