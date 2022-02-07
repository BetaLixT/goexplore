package graphHelper

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	azure "github.com/microsoft/kiota/authentication/go/azure"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)


type MSGraphOptions struct {
	TenantId string
	ClientId string
	Secret string
}
type GraphHelper struct {
	TenantId string
	ClientId string
	secret   string
	client *msgraphsdk.GraphServiceClient
}

func (graphHelper *GraphHelper) GetUser(userId string) (*graph.User, error) {
	user, err := (*graphHelper).client.UsersById(userId).Get(nil)
	return user, err
} 

func NewGraphHelper(msGraphOptions MSGraphOptions) (graphHelper *GraphHelper, e error) {
	graphHelper = &GraphHelper{
		TenantId: msGraphOptions.TenantId,
		ClientId: msGraphOptions.ClientId,
		secret: msGraphOptions.Secret,
	}
	
	err := graphHelper.generateApplicationCredentialClient()
	if err != nil {
		e = fmt.Errorf("[ERR]Failed to generate graph client %s", err)
		return
	}
	return
}

func (graphHelper *GraphHelper) generateApplicationCredentialClient() (e error) {

	options := new(azidentity.ClientSecretCredentialOptions)

	// - Generating credentials
	cred, err := azidentity.NewClientSecretCredential(
		(*graphHelper).TenantId,
		(*graphHelper).ClientId,
		(*graphHelper).secret,
		options,
	)
	if err != nil {
		e = fmt.Errorf("[ERR]Failed to generate client credential %s", err)
		return
	}

	// - Generating authentication provider
	auth, err := azure.NewAzureIdentityAuthenticationProvider(cred)
	if err != nil {
		e = fmt.Errorf("[ERR]Failed to generate auth provider %v", err)
		return
	}

	// - Generating adapter
	adapter, err := msgraphsdk.NewGraphRequestAdapter(auth)
	if err != nil {
		e = fmt.Errorf("[ERR]Failed to generate adapter %s", err)
		return
	}

	(*graphHelper).client = msgraphsdk.NewGraphServiceClient(adapter)

	return
}
