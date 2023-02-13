package dnsimple

import (
	"context"
	"fmt"
)

func domainServicesPath(accountID, domainIdentifier string) (string, error) {
	path, err := domainPath(accountID, domainIdentifier)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/services", path), nil
}

func domainServicePath(accountID, domainIdentifier string, serviceIdentifier string) (string, error) {
	path, err := domainServicesPath(accountID, domainIdentifier)
	if err != nil {
		return "", err
	}

	if err := checkEmptyString("serviceIdentifier", serviceIdentifier); err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/%v", path, serviceIdentifier), nil
}

// DomainServiceSettings represents optional settings when applying a DNSimple one-click service to a domain.
type DomainServiceSettings struct {
	Settings map[string]string `url:"settings,omitempty"`
}

// AppliedServices lists the applied one-click services for a domain.
//
// See https://developer.dnsimple.com/v2/services/domains/#applied
func (s *ServicesService) AppliedServices(ctx context.Context, accountID string, domainIdentifier string, options *ListOptions) (*ServicesResponse, error) {
	path, err := domainServicesPath(accountID, domainIdentifier)
	if err != nil {
		return nil, err
	}

	path = versioned(path)

	servicesResponse := &ServicesResponse{}

	path, err = addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, servicesResponse)
	if err != nil {
		return servicesResponse, err
	}

	servicesResponse.HTTPResponse = resp
	return servicesResponse, nil
}

// ApplyService applies a one-click services to a domain.
//
// See https://developer.dnsimple.com/v2/services/domains/#apply
func (s *ServicesService) ApplyService(ctx context.Context, accountID string, serviceIdentifier string, domainIdentifier string, settings DomainServiceSettings) (*ServiceResponse, error) {
	path, err := domainServicePath(accountID, domainIdentifier, serviceIdentifier)
	if err != nil {
		return nil, err
	}

	path = versioned(path)

	serviceResponse := &ServiceResponse{}

	resp, err := s.client.post(ctx, path, settings, nil)
	if err != nil {
		return nil, err
	}

	serviceResponse.HTTPResponse = resp
	return serviceResponse, nil
}

// UnapplyService unapplies a one-click services from a domain.
//
// See https://developer.dnsimple.com/v2/services/domains/#unapply
func (s *ServicesService) UnapplyService(ctx context.Context, accountID string, serviceIdentifier string, domainIdentifier string) (*ServiceResponse, error) {
	path, err := domainServicePath(accountID, domainIdentifier, serviceIdentifier)
	if err != nil {
		return nil, err
	}

	path = versioned(path)

	serviceResponse := &ServiceResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	serviceResponse.HTTPResponse = resp
	return serviceResponse, nil
}
