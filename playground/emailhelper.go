package graphhelper

import (
	"context"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	auth "github.com/microsoft/kiota-authentication-azure-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

func SendEmail() {

	clientId := os.Getenv("CLIENT_ID")
	tenantId := os.Getenv("TENANT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	scopes := os.Getenv("GRAPH_USER_SCOPES")
	graphUserScopes := strings.Split(scopes, ",")

	// Create the device code credential
	credential, err := azidentity.NewClientSecretCredential(tenantId,
		clientId,
		clientSecret,
		nil)

	if err != nil {
	}

	// Create an auth provider using the credential
	authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(credential, graphUserScopes)
	if err != nil {
	}

	// Create a request adapter using the auth provider
	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
	if err != nil {
	}

	// Create a Graph client using request adapter
	client := msgraphsdk.NewGraphServiceClient(adapter)
	subject := "Email Subject"
	body := "Email Content"
	senderEmail := "sender_email@eexample.com"
	recipient := "recipient_email@eexample.com"
	message := models.NewMessage()
	message.SetSubject(&subject)

	messageBody := models.NewItemBody()
	messageBody.SetContent(&body)
	contentType := models.TEXT_BODYTYPE
	messageBody.SetContentType(&contentType)
	message.SetBody(messageBody)

	toRecipient := models.NewRecipient()
	emailAddress := models.NewEmailAddress()
	emailAddress.SetAddress(&recipient)
	toRecipient.SetEmailAddress(emailAddress)
	message.SetToRecipients([]models.Recipientable{
		toRecipient,
	})

	sendMailBody := users.NewItemSendMailPostRequestBody()
	sendMailBody.SetMessage(message)

	// Send the message
	err = client.Users().ByUserId(senderEmail).SendMail().Post(context.Background(), sendMailBody, nil)
	if err != nil {

	}
}
