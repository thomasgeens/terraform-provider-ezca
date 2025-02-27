// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CAChainInfoModel struct {
	RootCertificate    *BasicCertInfoModel  `json:"RootCertificate,omitempty"`
	SubordinateCAChain []BasicCertInfoModel `json:"SubordinateCAChain,omitempty"`
}

func (o *CAChainInfoModel) GetRootCertificate() *BasicCertInfoModel {
	if o == nil {
		return nil
	}
	return o.RootCertificate
}

func (o *CAChainInfoModel) GetSubordinateCAChain() []BasicCertInfoModel {
	if o == nil {
		return nil
	}
	return o.SubordinateCAChain
}
