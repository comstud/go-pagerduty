package pagerduty

import "net/http"

// IncidentsService type
type IncidentsService struct {
	client *Client
}

// Incident type
type Incident struct {
	ID                    string            `json:"id,omitempty"`
	IncidentNumber        int               `json:"incident_number,omitempty"`
	Status                string            `json:"status,omitempty"`
	CreatedOn             string            `json:"created_on,omitempty"`
	Summary               *IncidentSummary  `json:"trigger_summary_data,omitempty"`
	User                  *User             `json:"assigned_to_user,omitempty"`
	Service               *Service          `json:"service,omitempty"`
	EscalationPolicy      *EscalationPolicy `json:"escalation_policy,omitempty"`
	HTMLURL               string            `json:"html_url,omitempty"`
	IncidentKey           string            `json:"incident_key,omitempty"`
	TriggerDetailsHTMLURL string            `json:"trigger_details_html_url,omitempty"`
	TriggerType           string            `json:"trigger_type,omitempty"`
	LastStatusChangeOn    string            `json:"last_status_change_on,omitempty"`
	LastStatusChangeBy    *User             `json:"last_status_change_by,omitempty"`
	NumberOfEscalations   int               `json:"number_of_escalations,omitempty"`
	ResolvedByUser        *User             `json:"resolved_by_user,omitempty"`
	AssignedToUser        *User             `json:"assigned_to_user,omitempty"`
	AssignedTo            []*UserObject     `json:"assigned_to,omitempty"`
}

// Incidents is a list of incidents
type Incidents struct {
	Incidents []Incident
}

// IncidentSummary type
type IncidentSummary struct {
	Subject     string `json:"subject,omitempty"`
	Hostname    string `json:"HOSTNAME"`
	Description string `json:"pd_description,omitempty"`
}

// Get returns a single incident by id if found
func (s *IncidentsService) Get(id string) (*Incident, *http.Response, error) {
	incident := new(Incident)

	res, err := s.client.Get("incidents/"+id, incident)
	if err != nil {
		return nil, res, err
	}

	return incident, res, nil
}

// IncidentsOptions provides optional parameters to list requests
type IncidentsOptions struct {
	Status string `url:"status,omitempty"`
	SortBy string `url:"sort_by,omitempty"`
	Since  string `url:"since,omitempty"`
	Until  string `url:"until,omitempty"`
}

// List returns a list of incidents
func (s *IncidentsService) List(opt *IncidentsOptions) ([]Incident, *http.Response, error) {
	u, err := addOptions("incidents", opt)
	if err != nil {
		return nil, nil, err
	}

	incidents := new(Incidents)

	res, err := s.client.Get(u, incidents)
	if err != nil {
		return nil, res, err
	}

	return incidents.Incidents, res, err
}

type IncidentResolveBody struct {
	// RequesterID is the ID of a user. Required in token auth mode.
	RequesterID string `json:"requester_id,omitempty"`
}

// Resolve an incident
func (s *IncidentsService) Resolve(id string, body *IncidentResolveBody) (*Incident, *http.Response, error) {
	incident := new(Incident)

	res, err := s.client.Put("incidents/"+id+"/resolve", body, incident)
	if err != nil {
		return nil, res, err
	}

	return incident, res, err
}

type IncidentAcknowledgeBody struct {
	// RequesterID is the ID of a user. Required in token auth mode.
	RequesterID string `json:"requester_id,omitempty"`
}

// Acknowledge an incident
func (s *IncidentsService) Acknowledge(id string, body *IncidentAcknowledgeBody) (*Incident, *http.Response, error) {
	incident := new(Incident)

	res, err := s.client.Put("incidents/"+id+"/acknowledge", body, incident)
	if err != nil {
		return nil, res, err
	}

	return incident, res, err
}

// IncidentSnoozeBody provides the body for the Snooze request
type IncidentSnoozeBody struct {
	// RequesterID is the ID of a user. Required in token auth mode.
	RequesterID string `json:"requester_id,omitempty"`
	// Duration is the number of seconds to snooze. It is required.
	Duration int `json:"duration"`
}

// Snooze an incident
func (s *IncidentsService) Snooze(id string, body *IncidentSnoozeBody) (*Incident, *http.Response, error) {
	incident := new(Incident)

	res, err := s.client.Put("incidents/"+id+"/snooze", body, incident)
	if err != nil {
		return nil, res, err
	}

	return incident, res, err
}
