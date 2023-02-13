package dnsimple

import (
	"context"
)

// EnableDomainAutoRenewal enables auto-renewal for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/auto-renewal/#enable
func (s *RegistrarService) EnableDomainAutoRenewal(ctx context.Context, accountID string, domainName string) (*DomainResponse, error) {
	path, err := registrarDomainPath(accountID, domainName)
	if err != nil {
		return nil, err
	}

	path = versioned(path + "/auto_renewal")

	domainResponse := &DomainResponse{}

	resp, err := s.client.put(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	domainResponse.HTTPResponse = resp
	return domainResponse, nil
}

// DisableDomainAutoRenewal disables auto-renewal for the domain.
//
// See https://developer.dnsimple.com/v2/registrar/auto-renewal/#enable
func (s *RegistrarService) DisableDomainAutoRenewal(ctx context.Context, accountID string, domainName string) (*DomainResponse, error) {
	path, err := registrarDomainPath(accountID, domainName)
	if err != nil {
		return nil, err
	}

	path = versioned(path + "/auto_renewal")

	domainResponse := &DomainResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	domainResponse.HTTPResponse = resp
	return domainResponse, nil
}
