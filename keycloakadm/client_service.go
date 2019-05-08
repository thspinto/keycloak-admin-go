package keycloakadm

import (
	"context"
	"net/url"
	"strings"
)

// ClientService interacts with all user resources
type ClientService service

// Clients returns a new client service for working with client resources
// in a realm.
func (c *Client) Clients() *ClientService {
	return &ClientService{
		client: c,
	}
}

// Create creates a new client and returns the ID
// Response is a 201 with a location redirect
func (s *ClientService) Create(ctx context.Context, client *ClientRepresentation) (string, error) {
	path := "/realms/{realm}/clients"

	response, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetBody(client).
		Post(path)

	if err != nil {
		return "", err
	}

	location, err := url.Parse(response.Header().Get("Location"))

	if err != nil {
		return "", err
	}

	components := strings.Split(location.Path, "/")

	return components[len(components)-1], nil
}

// Get returns a client in a realm
func (s *ClientService) Get(ctx context.Context, ID string) (*ClientRepresentation, error) {

	path := "/realms/{realm}/clients/{id}"

	client := &ClientRepresentation{}

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    ID,
		}).
		SetResult(client).
		Get(path)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// Update updates the given client
func (s *ClientService) Update(ctx context.Context, client *ClientRepresentation) error {

	path := "/realms/{realm}/clients/{id}"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    client.ID,
		}).
		SetBody(client).
		Put(path)

	return err
}

// Find returns clients based on query params
// Params:
// - clientId
func (s *ClientService) Find(ctx context.Context, params map[string]string) ([]ClientRepresentation, error) {

	path := "/realms/{realm}/clients"

	var clients []ClientRepresentation

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
		}).
		SetQueryParams(params).
		SetResult(&clients).
		Get(path)

	if err != nil {
		return nil, err
	}

	return clients, nil
}

// GetRole gets a client role by name
func (s *ClientService) GetRole(ctx context.Context, client *ClientRepresentation, roleName string) (*RoleRepresentation, error) {

	path := "/realms/{realm}/clients/{id}/roles/{role-name}"

	role := &RoleRepresentation{}

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"id":        client.ID,
			"role-name": roleName,
		}).
		SetResult(role).
		Get(path)

	if err != nil {
		return nil, err
	}

	return role, nil
}

// ListRoles returns all the client roles
func (s *ClientService) ListRoles(ctx context.Context, client *ClientRepresentation) ([]RoleRepresentation, error) {

	path := "/realms/{realm}/clients/{id}/roles"

	var roles []RoleRepresentation

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    client.ID,
		}).
		SetResult(&roles).
		Get(path)

	if err != nil {
		return nil, err
	}

	return roles, nil
}
