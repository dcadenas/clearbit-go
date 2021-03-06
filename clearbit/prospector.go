package clearbit

import (
	"github.com/dghubble/sling"
	"net/http"
)

const (
	prospectorBase = "https://prospector.clearbit.com"
)

// ProspectorItem represents each of the items returned by a call to Search
type ProspectorItem struct {
	ID   string `json:"id"`
	Name struct {
		FullName   string `json:"fullName"`
		GivenName  string `json:"givenName"`
		FamilyName string `json:"familyName"`
	} `json:"name"`
	Title string `json:"title"`
	Email string `json:"email"`
}

// ProspectorSearchParams wraps the parameters needed to interact with the
// Prospector API
type ProspectorSearchParams struct {
	Domain    string `url:"domain,omitempty"`
	Role      string `url:"role,omitempty"`
	Seniority string `url:"seniority,omitempty"`
	Title     string `url:"title,omitempty"`
	Name      string `url:"name,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

// ProspectorService gives access to the Prospector API.
//
// The Prospector API lets you fetch contacts and emails associated with a
// company, employment role, seniority, and job title.
type ProspectorService struct {
	baseSling *sling.Sling
	sling     *sling.Sling
}

func newProspectorService(sling *sling.Sling) *ProspectorService {
	return &ProspectorService{
		baseSling: sling.New(),
		sling:     sling.Base(prospectorBase).Path("/v1/people/"),
	}
}

// Search lets you fetch contacts and emails associated with a company,
// employment role, seniority, and job title.
func (s *ProspectorService) Search(params ProspectorSearchParams) ([]ProspectorItem, *http.Response, error) {
	items := new([]ProspectorItem)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("search").QueryStruct(params).Receive(items, apiError)
	return *items, resp, relevantError(err, *apiError)
}
