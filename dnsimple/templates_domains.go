package dnsimple

import (
	"context"
)

// ApplyTemplate applies a template to the given domain.
//
// See https://developer.dnsimple.com/v2/templates/domains/#applyTemplateToDomain
func (s *TemplatesService) ApplyTemplate(ctx context.Context, accountID string, templateIdentifier string, domainIdentifier string) (*TemplateResponse, error) {
	path, err := domainPath(accountID, domainIdentifier)
	if err != nil {
		return nil, err
	}

	path = versioned(path + "/templates")

	templateResponse := &TemplateResponse{}

	resp, err := s.client.post(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	templateResponse.HTTPResponse = resp
	return templateResponse, nil
}
