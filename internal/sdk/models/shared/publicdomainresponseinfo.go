// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/thomasgeens/terraform-provider-ezca/internal/sdk/internal/utils"
	"time"
)

type PublicDomainResponseInfo struct {
	TenantID             *string    `json:"TenantID,omitempty"`
	Caid                 *string    `json:"CAID,omitempty"`
	DomainID             *string    `json:"DomainID,omitempty"`
	DomainName           *string    `json:"DomainName,omitempty"`
	Validated            *bool      `json:"Validated,omitempty"`
	ValidationMethod     *string    `json:"ValidationMethod,omitempty"`
	ValidationLevel      *string    `json:"ValidationLevel,omitempty"`
	ValidationDateExpiry *time.Time `json:"ValidationDateExpiry,omitempty"`
	AllowWildCards       *bool      `json:"AllowWildCards,omitempty"`
	IsRegistered         *bool      `json:"IsRegistered,omitempty"`
	NewRegistration      *bool      `json:"NewRegistration,omitempty"`
	AllowACME            *bool      `json:"AllowACME,omitempty"`
}

func (p PublicDomainResponseInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PublicDomainResponseInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PublicDomainResponseInfo) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}

func (o *PublicDomainResponseInfo) GetCaid() *string {
	if o == nil {
		return nil
	}
	return o.Caid
}

func (o *PublicDomainResponseInfo) GetDomainID() *string {
	if o == nil {
		return nil
	}
	return o.DomainID
}

func (o *PublicDomainResponseInfo) GetDomainName() *string {
	if o == nil {
		return nil
	}
	return o.DomainName
}

func (o *PublicDomainResponseInfo) GetValidated() *bool {
	if o == nil {
		return nil
	}
	return o.Validated
}

func (o *PublicDomainResponseInfo) GetValidationMethod() *string {
	if o == nil {
		return nil
	}
	return o.ValidationMethod
}

func (o *PublicDomainResponseInfo) GetValidationLevel() *string {
	if o == nil {
		return nil
	}
	return o.ValidationLevel
}

func (o *PublicDomainResponseInfo) GetValidationDateExpiry() *time.Time {
	if o == nil {
		return nil
	}
	return o.ValidationDateExpiry
}

func (o *PublicDomainResponseInfo) GetAllowWildCards() *bool {
	if o == nil {
		return nil
	}
	return o.AllowWildCards
}

func (o *PublicDomainResponseInfo) GetIsRegistered() *bool {
	if o == nil {
		return nil
	}
	return o.IsRegistered
}

func (o *PublicDomainResponseInfo) GetNewRegistration() *bool {
	if o == nil {
		return nil
	}
	return o.NewRegistration
}

func (o *PublicDomainResponseInfo) GetAllowACME() *bool {
	if o == nil {
		return nil
	}
	return o.AllowACME
}
