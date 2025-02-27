// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CertificateCreateRequestModel struct {
	SubjectName         *string  `json:"SubjectName,omitempty"`
	SubjectAltNames     []string `json:"SubjectAltNames,omitempty"`
	Caid                *string  `json:"CAID,omitempty"`
	TemplateID          *string  `json:"TemplateID,omitempty"`
	Csr                 *string  `json:"CSR,omitempty"`
	ValidityInDays      *int     `json:"ValidityInDays,omitempty"`
	EKUs                []string `json:"EKUs,omitempty"`
	KeyUsages           []string `json:"KeyUsages,omitempty"`
	SelectedLocation    *string  `json:"SelectedLocation,omitempty"`
	ResourceID          *string  `json:"ResourceID,omitempty"`
	SecretName          *string  `json:"SecretName,omitempty"`
	AKVName             *string  `json:"AKVName,omitempty"`
	AutoRenew           *bool    `json:"AutoRenew,omitempty"`
	AutoRenewPercentage *int     `json:"AutoRenewPercentage,omitempty"`
	CertAppID           *string  `json:"CertAppID,omitempty"`
	CertificateTags     *string  `json:"CertificateTags,omitempty"`
	Sid                 *string  `json:"SID,omitempty"`
}

func (o *CertificateCreateRequestModel) GetSubjectName() *string {
	if o == nil {
		return nil
	}
	return o.SubjectName
}

func (o *CertificateCreateRequestModel) GetSubjectAltNames() []string {
	if o == nil {
		return nil
	}
	return o.SubjectAltNames
}

func (o *CertificateCreateRequestModel) GetCaid() *string {
	if o == nil {
		return nil
	}
	return o.Caid
}

func (o *CertificateCreateRequestModel) GetTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

func (o *CertificateCreateRequestModel) GetCsr() *string {
	if o == nil {
		return nil
	}
	return o.Csr
}

func (o *CertificateCreateRequestModel) GetValidityInDays() *int {
	if o == nil {
		return nil
	}
	return o.ValidityInDays
}

func (o *CertificateCreateRequestModel) GetEKUs() []string {
	if o == nil {
		return nil
	}
	return o.EKUs
}

func (o *CertificateCreateRequestModel) GetKeyUsages() []string {
	if o == nil {
		return nil
	}
	return o.KeyUsages
}

func (o *CertificateCreateRequestModel) GetSelectedLocation() *string {
	if o == nil {
		return nil
	}
	return o.SelectedLocation
}

func (o *CertificateCreateRequestModel) GetResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceID
}

func (o *CertificateCreateRequestModel) GetSecretName() *string {
	if o == nil {
		return nil
	}
	return o.SecretName
}

func (o *CertificateCreateRequestModel) GetAKVName() *string {
	if o == nil {
		return nil
	}
	return o.AKVName
}

func (o *CertificateCreateRequestModel) GetAutoRenew() *bool {
	if o == nil {
		return nil
	}
	return o.AutoRenew
}

func (o *CertificateCreateRequestModel) GetAutoRenewPercentage() *int {
	if o == nil {
		return nil
	}
	return o.AutoRenewPercentage
}

func (o *CertificateCreateRequestModel) GetCertAppID() *string {
	if o == nil {
		return nil
	}
	return o.CertAppID
}

func (o *CertificateCreateRequestModel) GetCertificateTags() *string {
	if o == nil {
		return nil
	}
	return o.CertificateTags
}

func (o *CertificateCreateRequestModel) GetSid() *string {
	if o == nil {
		return nil
	}
	return o.Sid
}
