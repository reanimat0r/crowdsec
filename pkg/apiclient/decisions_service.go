package apiclient

import (
	"context"
	"fmt"

	"github.com/crowdsecurity/crowdsec/pkg/models"
	qs "github.com/google/go-querystring/query"
)

type DecisionsService service

type DecisionsListOpts struct {
	Scope_equals *string `url:"scope,omitempty"`
	Value_equals *string `url:"value,omitempty"`
	Type_equals  *string `url:"type,omitempty"`
	IP_equals    *string `url:"ip,omitempty"`
	Range_equals *string `url:"range,omitempty"`
	ListOpts
}

type DecisionsDeleteOpts struct {
	Scope_equals *string `url:"scope,omitempty"`
	Value_equals *string `url:"value,omitempty"`
	Type_equals  *string `url:"type,omitempty"`
	IP_equals    *string `url:"ip,omitempty"`
	Range_equals *string `url:"range,omitempty"`
	ListOpts
}

//to demo query arguments
func (s *DecisionsService) List(ctx context.Context, opts DecisionsListOpts) (*models.GetDecisionsResponse, *Response, error) {
	var decisions models.GetDecisionsResponse
	params, err := qs.Values(opts)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("v1/decisions/?%s", params.Encode())

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, &decisions)
	if err != nil {
		return nil, resp, err
	}
	return &decisions, resp, nil
}

func (s *DecisionsService) GetStream(ctx context.Context, startup bool) (*models.DecisionsStreamResponse, *Response, error) {
	var decisions models.DecisionsStreamResponse

	u := fmt.Sprintf("v1/decisions/stream?startup=%t", startup)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, &decisions)
	if err != nil {
		return nil, resp, err
	}

	return &decisions, resp, nil
}

func (s *DecisionsService) StopStream(ctx context.Context) (*Response, error) {

	u := "v1/decisions"
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *DecisionsService) Delete(ctx context.Context, opts DecisionsDeleteOpts) (*models.DeleteDecisionResponse, *Response, error) {
	var deleteDecisionResponse models.DeleteDecisionResponse
	params, err := qs.Values(opts)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("v1/decisions?%s", params.Encode())

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, &deleteDecisionResponse)
	if err != nil {
		return nil, resp, err
	}
	return &deleteDecisionResponse, resp, nil
}

func (s *DecisionsService) DeleteOne(ctx context.Context, decision_id string) (*models.DeleteDecisionResponse, *Response, error) {
	var deleteDecisionResponse models.DeleteDecisionResponse
	u := fmt.Sprintf("v1/decisions/%s", decision_id)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, &deleteDecisionResponse)
	if err != nil {
		return nil, resp, err
	}
	return &deleteDecisionResponse, resp, nil
}