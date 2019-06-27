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

// Realm returns the reaml the service is operating in
func (s *ClientService) Realm() string {
	return s.client.Realm
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

// Delete deletes a client
func (s *ClientService) Delete(ctx context.Context, client *ClientRepresentation) error {
	path := "/realms/{realm}/clients/{id}"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    client.ID,
		}).
		Delete(path)

	return err
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

// CreateRole creates a role
func (s *ClientService) CreateRole(ctx context.Context, client *ClientRepresentation, role *RoleRepresentation) error {

	path := "/realms/{realm}/clients/{id}/roles"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    client.ID,
		}).
		SetBody(role).
		Post(path)

	return err
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

// DeleteRole deletes a role
func (s *ClientService) DeleteRole(ctx context.Context, client *ClientRepresentation, role *RoleRepresentation) error {

	path := "/realms/{realm}/clients/{id}/roles/{role-name}"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"id":        client.ID,
			"role-name": role.Name,
		}).
		Delete(path)

	return err
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

// AddProtocolMappers adds protocol mappers
func (s *ClientService) AddProtocolMappers(ctx context.Context, client *ClientRepresentation, mappers []ProtocolMapperRepresentation) error {
	path := "/realms/{realm}/clients/{id}/protocol-mappers/add-models"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    client.ID,
		}).
		SetBody(mappers).
		Post(path)

	return err
}

// GetProtocolMappers gets protocol mappers
func (s *ClientService) GetProtocolMappers(ctx context.Context, client *ClientRepresentation) ([]ProtocolMapperRepresentation, error) {
	path := "realms/{realm}/clients/{id}/protocol-mappers/models"

	var mappers []ProtocolMapperRepresentation

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": s.client.Realm,
			"id":    client.ID,
		}).
		SetResult(&mappers).
		Get(path)

	if err != nil {
		return nil, err
	}
	return mappers, nil
}

// UpdateProtocolMapper updates a protocol mapper
func (s *ClientService) UpdateProtocolMapper(ctx context.Context, client *ClientRepresentation, mapper *ProtocolMapperRepresentation) error {
	path := "realms/{realm}/clients/{id}/protocol-mappers/models/{mapper_id}"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"id":        client.ID,
			"mapper_id": mapper.ID,
		}).
		SetBody(mapper).
		Put(path)

	return err
}

// DeleteProtocolMapper deletes a protocol mapper
func (s *ClientService) DeleteProtocolMapper(ctx context.Context, client *ClientRepresentation, mapper *ProtocolMapperRepresentation) error {
	path := "realms/{realm}/clients/{id}/protocol-mappers/models/{mapper_id}"

	_, err := s.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":     s.client.Realm,
			"id":        client.ID,
			"mapper_id": mapper.ID,
		}).
		Delete(path)

	return err
}
