package integration_test

import (
	"github.com/thspinto/keycloak-admin-go/keycloakadm"
)

func (suite *integrationTester) TestClientFetch() {
	clientName := "admin-cli"
	clients, err := suite.client.Clients().Find(suite.ctx, map[string]string{
		"clientId": clientName,
	})
	suite.NotNil(clients, suite.version)
	suite.NoError(err, suite.version)
	suite.Len(clients, 1, suite.version)

	client, err := suite.client.Clients().Get(suite.ctx, clients[0].ID)
	suite.NotNil(clients, suite.version)
	suite.NoError(err, suite.version)
	suite.Equal(clientName, client.ClientID, suite.version)
}

func (suite *integrationTester) TestClientCreateUpdateDelete() {
	client := &keycloakadm.ClientRepresentation{
		ClientID: pseudoRandString(),
	}

	id, err := suite.client.Clients().Create(suite.ctx, client)
	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)
	client, err = suite.client.Clients().Get(suite.ctx, id)
	suite.NotNil(client, suite.version)
	suite.NoError(err, suite.version)

	redirectURI := "test"
	client.RedirectURIs = []string{redirectURI}
	err = suite.client.Clients().Update(suite.ctx, client)
	suite.NoError(err, suite.version)

	client, err = suite.client.Clients().Get(suite.ctx, id)
	suite.NotNil(client, suite.version)
	suite.NoError(err, suite.version)
	suite.Equal(redirectURI, client.RedirectURIs[0])

	err = suite.client.Clients().Delete(suite.ctx, client)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestClientRolesFetch() {
	clientName := "account"
	clients, err := suite.client.Clients().Find(suite.ctx, map[string]string{
		"clientId": clientName,
	})
	suite.NotNil(clients, suite.version)
	suite.NoError(err, suite.version)
	suite.Len(clients, 1, suite.version)

	roles, err := suite.client.Clients().ListRoles(suite.ctx, &clients[0])
	suite.Len(roles, 3, suite.version)
	suite.NoError(err, suite.version)
	role, err := suite.client.Clients().GetRole(suite.ctx, &clients[0], "manage-account")
	suite.NotNil(role, suite.version)
	suite.NoError(err, suite.version)
	suite.Equal("manage-account", role.Name, suite.version)
}

func (suite *integrationTester) TestProtocolMapper() {
	client := &keycloakadm.ClientRepresentation{
		ClientID: pseudoRandString(),
	}

	id, err := suite.client.Clients().Create(suite.ctx, client)
	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)

	client, err = suite.client.Clients().Get(suite.ctx, id)
	suite.NotNil(client, suite.version)
	suite.NoError(err, suite.version)

	mapper := keycloakadm.ProtocolMapperRepresentation{
		Name:           "TestMapper",
		Protocol:       "openid-connect",
		ProtocolMapper: "oidc-usermodel-realm-role-mapper",
		Config: keycloakadm.AttributeMap{
			"claim.name":           "test",
			"access.token.claim":   "true",
			"id.token.claim":       "true",
			"userinfo.token.claim": "true",
		},
	}

	err = suite.client.Clients().AddProtocolMappers(suite.ctx, client, []keycloakadm.ProtocolMapperRepresentation{mapper})
	suite.NoError(err, suite.version)

	mappers, err := suite.client.Clients().GetProtocolMappers(suite.ctx, client)
	suite.NotNil(client, suite.version)
	suite.NoError(err, suite.version)
	suite.Equal(1, len(mappers))
	suite.Equal("TestMapper", mappers[0].Name)

	// Update Mapper
	mappers[0].Config["claim.name"] = "TestClaim"
	err = suite.client.Clients().UpdateProtocolMapper(suite.ctx, client, &mappers[0])
	suite.NoError(err, suite.version)

	// Check if mapper updated
	mappers, err = suite.client.Clients().GetProtocolMappers(suite.ctx, client)
	suite.NotNil(client, suite.version)
	suite.NoError(err, suite.version)
	suite.Equal(1, len(mappers))
	suite.Equal("TestClaim", mappers[0].Config["claim.name"])

	// Delete mapper
	err = suite.client.Clients().DeleteProtocolMapper(suite.ctx, client, &mappers[0])
	suite.NoError(err, suite.version)

	// Check mapper deleted
	mappers, err = suite.client.Clients().GetProtocolMappers(suite.ctx, client)
	suite.NoError(err, suite.version)
	suite.Equal(0, len(mappers))

	err = suite.client.Clients().Delete(suite.ctx, client)
	suite.NoError(err, suite.version)
}
