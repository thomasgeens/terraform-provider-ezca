// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DBSelfServiceScep struct {
	TenantID              *string `json:"TenantID,omitempty"`
	CaID                  *string `json:"CaID,omitempty"`
	TemplateID            *string `json:"TemplateID,omitempty"`
	ProfileID             *string `json:"ProfileID,omitempty"`
	AllUsers              *bool   `json:"allUsers,omitempty"`
	MultipleDevices       *bool   `json:"MultipleDevices,omitempty"`
	OtherTenants          *bool   `json:"otherTenants,omitempty"`
	OtherTenantsIds       *string `json:"otherTenantsIds,omitempty"`
	SubjectName           *string `json:"subjectName,omitempty"`
	SubjectAltNames       *string `json:"subjectAltNames,omitempty"`
	DurationInDays        *int    `json:"durationInDays,omitempty"`
	Ekus                  *string `json:"Ekus,omitempty"`
	KeyUsages             *string `json:"KeyUsages,omitempty"`
	PolicyName            *string `json:"PolicyName,omitempty"`
	EncryptionKeyLocation *string `json:"EncryptionKeyLocation,omitempty"`
	BehalfOfAgents        *string `json:"BehalfOfAgents,omitempty"`
}

func (o *DBSelfServiceScep) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}

func (o *DBSelfServiceScep) GetCaID() *string {
	if o == nil {
		return nil
	}
	return o.CaID
}

func (o *DBSelfServiceScep) GetTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.TemplateID
}

func (o *DBSelfServiceScep) GetProfileID() *string {
	if o == nil {
		return nil
	}
	return o.ProfileID
}

func (o *DBSelfServiceScep) GetAllUsers() *bool {
	if o == nil {
		return nil
	}
	return o.AllUsers
}

func (o *DBSelfServiceScep) GetMultipleDevices() *bool {
	if o == nil {
		return nil
	}
	return o.MultipleDevices
}

func (o *DBSelfServiceScep) GetOtherTenants() *bool {
	if o == nil {
		return nil
	}
	return o.OtherTenants
}

func (o *DBSelfServiceScep) GetOtherTenantsIds() *string {
	if o == nil {
		return nil
	}
	return o.OtherTenantsIds
}

func (o *DBSelfServiceScep) GetSubjectName() *string {
	if o == nil {
		return nil
	}
	return o.SubjectName
}

func (o *DBSelfServiceScep) GetSubjectAltNames() *string {
	if o == nil {
		return nil
	}
	return o.SubjectAltNames
}

func (o *DBSelfServiceScep) GetDurationInDays() *int {
	if o == nil {
		return nil
	}
	return o.DurationInDays
}

func (o *DBSelfServiceScep) GetEkus() *string {
	if o == nil {
		return nil
	}
	return o.Ekus
}

func (o *DBSelfServiceScep) GetKeyUsages() *string {
	if o == nil {
		return nil
	}
	return o.KeyUsages
}

func (o *DBSelfServiceScep) GetPolicyName() *string {
	if o == nil {
		return nil
	}
	return o.PolicyName
}

func (o *DBSelfServiceScep) GetEncryptionKeyLocation() *string {
	if o == nil {
		return nil
	}
	return o.EncryptionKeyLocation
}

func (o *DBSelfServiceScep) GetBehalfOfAgents() *string {
	if o == nil {
		return nil
	}
	return o.BehalfOfAgents
}
