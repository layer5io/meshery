package models

import (
	"net/http"

	"github.com/gofrs/uuid"
	SMP "github.com/layer5io/service-mesh-performance/spec"
)

// ProviderType - for representing provider types
type ProviderType string

// ProviderProperties represents the structure of properties that a provider has
type ProviderProperties struct {
	ProviderType        ProviderType `json:"provider_type,omitempty"`
	PackageVersion      string       `json:"package_version,omitempty"`
	PackageURL          string       `json:"package_url,omitempty"`
	ProviderName        string       `json:"provider_name,omitempty"`
	ProviderDescription []string     `json:"provider_description,omitempty"`
	Extensions          Extensions   `json:"extensions,omitempty"`
	Capabilities        Capabilities `json:"capabilities,omitempty"`
}

// Extensions defines the UI extension points
type Extensions struct {
	Navigator NavigatorExtensions `json:"navigator,omitempty"`
	UserPrefs UserPrefsExtensions `json:"user_prefs,omitempty"`
}

// NavigatorExtensions is a collection of NavigatorExtension
type NavigatorExtensions []NavigatorExtension

// UserPrefsExtensions is a collection of UserPrefsExtension
type UserPrefsExtensions []UserPrefsExtension

// NavigatorExtension describes the Navigator extension point in the UI
type NavigatorExtension struct {
	Title     string              `json:"title,omitempty"`
	Href      Href                `json:"href,omitempty"`
	Component string              `json:"component,omitempty"`
	Icon      string              `json:"icon,omitempty"`
	Link      *bool               `json:"link,omitempty"`
	Show      *bool               `json:"show,omitempty"`
	Children  NavigatorExtensions `json:"children,omitempty"`
}

// UserPrefsExtension describes the user preference extension point in the UI
type UserPrefsExtension struct {
	Component string `json:"component,omitempty"`
}

// Href describes a link along with its type
type Href struct {
	URI      string `json:"uri,omitempty"`
	External *bool  `json:"external,omitempty"`
}

// Capabilities is the collection of capability
type Capabilities []Capability

// Capability is a capability of Provider indicating whether a feature is present
type Capability struct {
	Feature  Feature `json:"feature,omitempty"`
	Endpoint string  `json:"endpoint,omitempty"`
}

// Feature is a type to store the features of the provider
type Feature string

const (
	// SyncPrefs indicates the Preference Synchronization feature
	SyncPrefs Feature = "sync-prefs" // /user/preferences

	PersistResults Feature = "persist-results" // /results

	PersistResult Feature = "persist-result" // /result

	PersistSMIResult Feature = "persist-smi-result" // /smi/results

	PersistMetrics Feature = "persist-metrics" // /result/metrics

	PersistSMPTestProfile Feature = "persist-smp-test-profile" // /user/test-config
)

const (
	// LocalProviderType - represents local providers
	LocalProviderType ProviderType = "local"

	// RemoteProviderType - represents cloud providers
	RemoteProviderType ProviderType = "remote"

	// ProviderCtxKey is the context key for persisting provider to context
	ProviderCtxKey = "provider"
)

// IsSupported returns true if the given feature is listed as one of
// the capabilities of the provider
func (caps Capabilities) IsSupported(feature Feature) bool {
	for _, cap := range caps {
		if feature == cap.Feature {
			return true
		}
	}

	return false
}

// GetEndpointForFeature returns the endpoint for the given feature
//
// Existence of a feature DOES NOT guarantee that the endpoint would be a not empty
// string as some of the features may not require an endpoint
func (caps Capabilities) GetEndpointForFeature(feature Feature) (string, bool) {
	for _, cap := range caps {
		if feature == cap.Feature {
			return cap.Endpoint, true
		}
	}

	return "", false
}

// Provider - interface for providers
type Provider interface {
	PreferencePersister

	// Initialize will initialize a provider instance
	// by loading its capabilities and other metadata in the memory
	Initialize()

	Name() string

	// Returns ProviderType
	GetProviderType() ProviderType

	PackageLocation() string

	GetProviderCapabilities(http.ResponseWriter, *http.Request)

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
	FetchSmiResults(req *http.Request, page, pageSize, search, order string) ([]byte, error)
	PublishSmiResults(result *SmiResult) (string, error)
	PublishMetrics(tokenVal string, data *MesheryResult) error
	GetResult(*http.Request, uuid.UUID) (*MesheryResult, error)
	RecordPreferences(req *http.Request, userID string, data *Preference) error

	SMPTestConfigStore(req *http.Request, perfConfig *SMP.PerformanceTestConfig) (string, error)
	SMPTestConfigGet(req *http.Request, testUUID string) (*SMP.PerformanceTestConfig, error)
	SMPTestConfigFetch(req *http.Request, page, pageSize, search, order string) ([]byte, error)
	SMPTestConfigDelete(req *http.Request, testUUID string) error
}
