// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CAIssuerInformationModel struct {
	Organization     *string `json:"Organization,omitempty"`
	OrganizationUnit *string `json:"OrganizationUnit,omitempty"`
	CommonName       *string `json:"CommonName,omitempty"`
	CountryCode      *string `json:"CountryCode,omitempty"`
}

func (o *CAIssuerInformationModel) GetOrganization() *string {
	if o == nil {
		return nil
	}
	return o.Organization
}

func (o *CAIssuerInformationModel) GetOrganizationUnit() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationUnit
}

func (o *CAIssuerInformationModel) GetCommonName() *string {
	if o == nil {
		return nil
	}
	return o.CommonName
}

func (o *CAIssuerInformationModel) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}
