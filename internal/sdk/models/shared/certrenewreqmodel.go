// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CertRenewReqModel struct {
	Csr            *string `json:"CSR,omitempty"`
	ValidityInDays *int    `json:"ValidityInDays,omitempty"`
}

func (o *CertRenewReqModel) GetCsr() *string {
	if o == nil {
		return nil
	}
	return o.Csr
}

func (o *CertRenewReqModel) GetValidityInDays() *int {
	if o == nil {
		return nil
	}
	return o.ValidityInDays
}
