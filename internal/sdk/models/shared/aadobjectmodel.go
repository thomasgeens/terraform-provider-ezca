// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type AADObjectModel struct {
	ObjectID     *string `json:"ObjectId,omitempty"`
	FriendlyName *string `json:"FriendlyName,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty"`
	IsValid      *bool   `json:"isValid,omitempty"`
}

func (o *AADObjectModel) GetObjectID() *string {
	if o == nil {
		return nil
	}
	return o.ObjectID
}

func (o *AADObjectModel) GetFriendlyName() *string {
	if o == nil {
		return nil
	}
	return o.FriendlyName
}

func (o *AADObjectModel) GetObjectType() *string {
	if o == nil {
		return nil
	}
	return o.ObjectType
}

func (o *AADObjectModel) GetIsValid() *bool {
	if o == nil {
		return nil
	}
	return o.IsValid
}
