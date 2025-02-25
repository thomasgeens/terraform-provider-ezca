// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ACMECAsModel struct {
	IsEligible *bool              `json:"IsEligible,omitempty"`
	Agents     []DBACMEAgentModel `json:"Agents,omitempty"`
}

func (o *ACMECAsModel) GetIsEligible() *bool {
	if o == nil {
		return nil
	}
	return o.IsEligible
}

func (o *ACMECAsModel) GetAgents() []DBACMEAgentModel {
	if o == nil {
		return nil
	}
	return o.Agents
}
