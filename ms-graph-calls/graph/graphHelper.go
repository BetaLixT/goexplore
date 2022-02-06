package graphHelper

import "github.com/Azure/azure-sdk-for-go/sdk/azidentity"

type GraphHelper struct {
	TenantId string
	ClientId string
	secret   string
}

func NewGraphHelper(tenantId string, clientId string, secret string) (graphHelper *GraphHelper) {
	graphHelper = &GraphHelper{}
	return
}

func (graphHelper *GraphHelper) getClient() {

	// options := new azidentity.Option
	cred, err := azidentity.NewClientSecretCredential(
		(*graphHelper).TenantId,
		(*graphHelper).ClientId,
		(*graphHelper).secret,
	)
	return
}
