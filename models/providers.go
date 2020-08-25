package models

import (
	"net/http"

	"github.com/gofrs/uuid"
	SMPS "github.com/layer5io/service-mesh-performance-specification/spec"
)

// ProviderType - for representing provider types
type ProviderType string

// ProviderProperties represents the structure of properties that a provider has
type ProviderProperties struct {
	ProviderType
	DisplayName  string
	Description  string
	Capabilities []Capability
}

// Capability is a capability of Provider indicating whether a feature is present
type Capability struct {
	FeatureName string
	IsPresent   bool
}

const (
	// LocalProviderType - represents local providers
	LocalProviderType ProviderType = "local"

	// RemoteProviderType - represents cloud providers
	RemoteProviderType ProviderType = "remote"

	// ProviderCtxKey is the context key for persisting provider to context
	ProviderCtxKey = "provider"
)

// Provider - interface for providers
type Provider interface {
	PreferencePersister

	Name() string

	// Returns ProviderType
	GetProviderType() ProviderType

	GetProviderProperties() ProviderProperties
	// InitiateLogin - does the needed check, returns a true to indicate "return" or false to continue
	InitiateLogin(http.ResponseWriter, *http.Request, bool)
	TokenHandler(http.ResponseWriter, *http.Request, bool)
	ExtractToken(http.ResponseWriter, *http.Request)
	GetSession(req *http.Request) error
	GetUserDetails(*http.Request) (*User, error)
	GetProviderToken(req *http.Request) (string, error)
	UpdateToken(http.ResponseWriter, *http.Request)
	Logout(http.ResponseWriter, *http.Request)
	FetchResults(req *http.Request, page, pageSize, search, order string) ([]byte, error)
	PublishResults(req *http.Request, result *MesheryResult) (string, error)
	PublishMetrics(tokenVal string, data *MesheryResult) error
	GetResult(*http.Request, uuid.UUID) (*MesheryResult, error)
	RecordPreferences(req *http.Request, userID string, data *Preference) error

	SMPSTestConfigStore(req *http.Request, perfConfig *SMPS.PerformanceTestConfig) (string, error)
	SMPSTestConfigGet(req *http.Request, testUUID string) (*SMPS.PerformanceTestConfig, error)
	SMPSTestConfigFetch(req *http.Request, page, pageSize, search, order string) ([]byte, error)
	SMPSTestConfigDelete(req *http.Request, testUUID string) error
}
