package integration_test

import (
	"github.com/thspinto/keycloak-admin-go/keycloakadm"
)

func (suite *integrationTester) TestUserFetch() {
	users, err := suite.client.Users().Find(suite.ctx, map[string]string{
		"username": keycloakAdmin,
	})
	suite.NotNil(users, suite.version)
	suite.NoError(err, suite.version)
	suite.Len(users, 1, suite.version)
	suite.Equal(keycloakAdmin, users[0].Username, suite.version)
	suite.True(*users[0].Enabled, suite.version)

	user := users[0]
	t := true
	user.EmailVerified = &t

	err = suite.client.Users().Update(suite.ctx, &user)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestUserCreateDelete() {

	user := &keycloakadm.UserRepresentation{
		Username: pseudoRandString(),
		Email:    pseudoRandString() + "@example.com",
	}

	id, err := suite.client.Users().Create(suite.ctx, user)
	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)
	user.ID = id

	err = suite.client.Users().Delete(suite.ctx, user)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestReamlRolesAddListDelete() {

	user := &keycloakadm.UserRepresentation{
		Username: pseudoRandString(),
		Email:    pseudoRandString() + "@example.com",
	}

	id, err := suite.client.Users().Create(suite.ctx, user)
	user.ID = id

	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)

	role, err := suite.client.Roles().Get(suite.ctx, "create-realm")
	suite.NotEmpty(role, suite.version)
	suite.NoError(err, suite.version)

	err = suite.client.Users().AddRole(suite.ctx, user, *role)
	suite.NotEmpty(role, suite.version)
	suite.NoError(err, suite.version)

	realmRoles, err := suite.client.Users().ListRealmRoles(suite.ctx, user)
	suite.NotEmpty(realmRoles, suite.version)
	suite.NoError(err, suite.version)

	suite.Equal(3, len(realmRoles), suite.version)
	suite.NoError(err, suite.version)

	err = suite.client.Users().Delete(suite.ctx, user)
	suite.NoError(err, suite.version)
}

func (suite *integrationTester) TestClientRolesAddListDelete() {

	user := &keycloakadm.UserRepresentation{
		Username: pseudoRandString(),
		Email:    pseudoRandString() + "@example.com",
	}

	id, err := suite.client.Users().Create(suite.ctx, user)
	user.ID = id

	suite.NotEmpty(id, suite.version)
	suite.NoError(err, suite.version)

	clients, err := suite.client.Clients().Find(suite.ctx, map[string]string{
		"clientId": "broker",
	})
	suite.NotEmpty(clients, suite.version)
	suite.Equal(1, len(clients), suite.version)
	suite.NoError(err, suite.version)

	client := clients[0]

	roles, err := suite.client.Clients().ListRoles(suite.ctx, &client)
	suite.NotEmpty(roles, suite.version)
	suite.NoError(err, suite.version)

	err = suite.client.Users().AddRole(suite.ctx, user, roles[0])
	suite.NoError(err, suite.version)

	clientRoles, err := suite.client.Users().ListClientRoles(suite.ctx, user, &client)
	suite.NotEmpty(clientRoles, suite.version)
	suite.NoError(err, suite.version)

	suite.Equal(1, len(clientRoles), suite.version)
	suite.NoError(err, suite.version)

	err = suite.client.Users().Delete(suite.ctx, user)
	suite.NoError(err, suite.version)
}
