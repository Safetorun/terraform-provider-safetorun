// ampli.go
//
// Ampli - A strong typed wrapper for your Analytics
//
// This file is generated by Amplitude.
// To update run 'ampli pull admin-cli'
//
// Required dependencies: github.com/amplitude/analytics-go@latest
// Tracking Plan Version: 3
// Build: 1.0.0
// Runtime: go-ampli
//
// View Tracking Plan: https://data.amplitude.com/safetorun/Safetorun/events/main/latest
//
// Full Setup Instructions: https://data.amplitude.com/safetorun/Safetorun/implementation/main/latest/getting-started/admin-cli
//

package ampli

import (
	"log"
	"sync"

	"github.com/amplitude/analytics-go/amplitude"
)

type (
	EventOptions  = amplitude.EventOptions
	ExecuteResult = amplitude.ExecuteResult
)

const (
	IdentifyEventType      = amplitude.IdentifyEventType
	GroupIdentifyEventType = amplitude.GroupIdentifyEventType

	ServerZoneUS = amplitude.ServerZoneUS
	ServerZoneEU = amplitude.ServerZoneEU
)

var (
	NewClientConfig = amplitude.NewConfig
	NewClient       = amplitude.NewClient
)

var Instance = Ampli{}

type Environment string

const (
	EnvironmentProd Environment = `prod`
)

var APIKey = map[Environment]string{
	EnvironmentProd: `99aad7c317e1398a2cc1f5799d6b90ba`,
}

// LoadClientOptions is Client options setting to initialize Ampli client.
//
// Params:
//   - APIKey: the API key of Amplitude project
//   - Instance: the core SDK instance used by Ampli client
//   - Configuration: the core SDK client configuration instance
type LoadClientOptions struct {
	APIKey        string
	Instance      amplitude.Client
	Configuration amplitude.Config
}

// LoadOptions is options setting to initialize Ampli client.
//
// Params:
//   - Environment: the environment of Amplitude Data project
//   - Disabled: the flag of disabled Ampli client
//   - Client: the LoadClientOptions struct
type LoadOptions struct {
	Environment Environment
	Disabled    bool
	Client      LoadClientOptions
}

type baseEvent struct {
	eventType  string
	properties map[string]interface{}
}

type Event interface {
	ToAmplitudeEvent() amplitude.Event
}

func newBaseEvent(eventType string, properties map[string]interface{}) baseEvent {
	return baseEvent{
		eventType:  eventType,
		properties: properties,
	}
}

func (event baseEvent) ToAmplitudeEvent() amplitude.Event {
	return amplitude.Event{
		EventType:       event.eventType,
		EventProperties: event.properties,
	}
}

var CreateApplication = struct {
	Builder func() interface {
		OrganisationId(organisationId string) CreateApplicationBuilder
	}
}{
	Builder: func() interface {
		OrganisationId(organisationId string) CreateApplicationBuilder
	} {
		return &createApplicationBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type CreateApplicationEvent interface {
	Event
	createApplication()
}

type createApplicationEvent struct {
	baseEvent
}

func (e createApplicationEvent) createApplication() {
}

type CreateApplicationBuilder interface {
	Build() CreateApplicationEvent
}

type createApplicationBuilder struct {
	properties map[string]interface{}
}

func (b *createApplicationBuilder) OrganisationId(organisationId string) CreateApplicationBuilder {
	b.properties[`organisation_id`] = organisationId

	return b
}

func (b *createApplicationBuilder) Build() CreateApplicationEvent {
	return &createApplicationEvent{
		newBaseEvent(`CreateApplication`, b.properties),
	}
}

var CreateOrganisation = struct {
	Builder func() interface {
		OrganisationId(organisationId string) CreateOrganisationBuilder
	}
}{
	Builder: func() interface {
		OrganisationId(organisationId string) CreateOrganisationBuilder
	} {
		return &createOrganisationBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type CreateOrganisationEvent interface {
	Event
	createOrganisation()
}

type createOrganisationEvent struct {
	baseEvent
}

func (e createOrganisationEvent) createOrganisation() {
}

type CreateOrganisationBuilder interface {
	Build() CreateOrganisationEvent
}

type createOrganisationBuilder struct {
	properties map[string]interface{}
}

func (b *createOrganisationBuilder) OrganisationId(organisationId string) CreateOrganisationBuilder {
	b.properties[`organisation_id`] = organisationId

	return b
}

func (b *createOrganisationBuilder) Build() CreateOrganisationEvent {
	return &createOrganisationEvent{
		newBaseEvent(`CreateOrganisation`, b.properties),
	}
}

var DeleteApplication = struct {
	Builder func() interface {
		ApplicationId(applicationId string) interface {
			OrganisationId(organisationId string) DeleteApplicationBuilder
		}
	}
}{
	Builder: func() interface {
		ApplicationId(applicationId string) interface {
			OrganisationId(organisationId string) DeleteApplicationBuilder
		}
	} {
		return &deleteApplicationBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type DeleteApplicationEvent interface {
	Event
	deleteApplication()
}

type deleteApplicationEvent struct {
	baseEvent
}

func (e deleteApplicationEvent) deleteApplication() {
}

type DeleteApplicationBuilder interface {
	Build() DeleteApplicationEvent
}

type deleteApplicationBuilder struct {
	properties map[string]interface{}
}

func (b *deleteApplicationBuilder) ApplicationId(applicationId string) interface {
	OrganisationId(organisationId string) DeleteApplicationBuilder
} {
	b.properties[`application_id`] = applicationId

	return b
}

func (b *deleteApplicationBuilder) OrganisationId(organisationId string) DeleteApplicationBuilder {
	b.properties[`organisation_id`] = organisationId

	return b
}

func (b *deleteApplicationBuilder) Build() DeleteApplicationEvent {
	return &deleteApplicationEvent{
		newBaseEvent(`DeleteApplication`, b.properties),
	}
}

var DeleteOrganisation = struct {
	Builder func() interface {
		OrganisationId(organisationId string) DeleteOrganisationBuilder
	}
}{
	Builder: func() interface {
		OrganisationId(organisationId string) DeleteOrganisationBuilder
	} {
		return &deleteOrganisationBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type DeleteOrganisationEvent interface {
	Event
	deleteOrganisation()
}

type deleteOrganisationEvent struct {
	baseEvent
}

func (e deleteOrganisationEvent) deleteOrganisation() {
}

type DeleteOrganisationBuilder interface {
	Build() DeleteOrganisationEvent
}

type deleteOrganisationBuilder struct {
	properties map[string]interface{}
}

func (b *deleteOrganisationBuilder) OrganisationId(organisationId string) DeleteOrganisationBuilder {
	b.properties[`organisation_id`] = organisationId

	return b
}

func (b *deleteOrganisationBuilder) Build() DeleteOrganisationEvent {
	return &deleteOrganisationEvent{
		newBaseEvent(`DeleteOrganisation`, b.properties),
	}
}

var ListApplications = struct {
	Builder func() interface {
		OrganisationId(organisationId string) ListApplicationsBuilder
	}
}{
	Builder: func() interface {
		OrganisationId(organisationId string) ListApplicationsBuilder
	} {
		return &listApplicationsBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type ListApplicationsEvent interface {
	Event
	listApplications()
}

type listApplicationsEvent struct {
	baseEvent
}

func (e listApplicationsEvent) listApplications() {
}

type ListApplicationsBuilder interface {
	Build() ListApplicationsEvent
}

type listApplicationsBuilder struct {
	properties map[string]interface{}
}

func (b *listApplicationsBuilder) OrganisationId(organisationId string) ListApplicationsBuilder {
	b.properties[`organisation_id`] = organisationId

	return b
}

func (b *listApplicationsBuilder) Build() ListApplicationsEvent {
	return &listApplicationsEvent{
		newBaseEvent(`ListApplications`, b.properties),
	}
}

var ListOrgs = struct {
	Builder func() ListOrgsBuilder
}{
	Builder: func() ListOrgsBuilder {
		return &listOrgsBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type ListOrgsEvent interface {
	Event
	listOrgs()
}

type listOrgsEvent struct {
	baseEvent
}

func (e listOrgsEvent) listOrgs() {
}

type ListOrgsBuilder interface {
	Build() ListOrgsEvent
}

type listOrgsBuilder struct {
	properties map[string]interface{}
}

func (b *listOrgsBuilder) Build() ListOrgsEvent {
	return &listOrgsEvent{
		newBaseEvent(`ListOrgs`, b.properties),
	}
}

var UpdateApplication = struct {
	Builder func() interface {
		ApplicationId(applicationId string) interface {
			OrganisationId(organisationId string) UpdateApplicationBuilder
		}
	}
}{
	Builder: func() interface {
		ApplicationId(applicationId string) interface {
			OrganisationId(organisationId string) UpdateApplicationBuilder
		}
	} {
		return &updateApplicationBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type UpdateApplicationEvent interface {
	Event
	updateApplication()
}

type updateApplicationEvent struct {
	baseEvent
}

func (e updateApplicationEvent) updateApplication() {
}

type UpdateApplicationBuilder interface {
	Build() UpdateApplicationEvent
}

type updateApplicationBuilder struct {
	properties map[string]interface{}
}

func (b *updateApplicationBuilder) ApplicationId(applicationId string) interface {
	OrganisationId(organisationId string) UpdateApplicationBuilder
} {
	b.properties[`application_id`] = applicationId

	return b
}

func (b *updateApplicationBuilder) OrganisationId(organisationId string) UpdateApplicationBuilder {
	b.properties[`organisation_id`] = organisationId

	return b
}

func (b *updateApplicationBuilder) Build() UpdateApplicationEvent {
	return &updateApplicationEvent{
		newBaseEvent(`UpdateApplication`, b.properties),
	}
}

type Ampli struct {
	Disabled bool
	Client   amplitude.Client
	mutex    sync.RWMutex
}

// Load initializes the Ampli wrapper.
// Call once when your application starts.
func (a *Ampli) Load(options LoadOptions) {
	if a.Client != nil {
		log.Print("Warn: Ampli is already initialized. Ampli.Load() should be called once at application start up.")

		return
	}

	var apiKey string
	switch {
	case options.Client.APIKey != "":
		apiKey = options.Client.APIKey
	case options.Environment != "":
		apiKey = APIKey[options.Environment]
	default:
		apiKey = options.Client.Configuration.APIKey
	}

	if apiKey == "" && options.Client.Instance == nil {
		log.Print("Error: Ampli.Load() requires option.Environment, " +
			"and apiKey from either options.Instance.APIKey or APIKey[options.Environment], " +
			"or options.Instance.Instance")
	}

	clientConfig := options.Client.Configuration

	if clientConfig.Plan == nil {
		clientConfig.Plan = &amplitude.Plan{
			Branch:    `main`,
			Source:    `admin-cli`,
			Version:   `3`,
			VersionID: `429db6bb-e413-46b1-b931-2aab6c879039`,
		}
	}

	if clientConfig.IngestionMetadata == nil {
		clientConfig.IngestionMetadata = &amplitude.IngestionMetadata{
			SourceName:    `go-go-ampli`,
			SourceVersion: `2.0.0`,
		}
	}

	if options.Client.Instance != nil {
		a.Client = options.Client.Instance
	} else {
		clientConfig.APIKey = apiKey
		a.Client = amplitude.NewClient(clientConfig)
	}

	a.mutex.Lock()
	a.Disabled = options.Disabled
	a.mutex.Unlock()
}

// InitializedAndEnabled checks if Ampli is initialized and enabled.
func (a *Ampli) InitializedAndEnabled() bool {
	if a.Client == nil {
		log.Print("Error: Ampli is not yet initialized. Have you called Ampli.Load() on app start?")

		return false
	}

	a.mutex.RLock()
	defer a.mutex.RUnlock()

	return !a.Disabled
}

func (a *Ampli) setUserID(userID string, eventOptions *EventOptions) {
	if userID != "" {
		eventOptions.UserID = userID
	}
}

// Track tracks an event.
func (a *Ampli) Track(userID string, event Event, eventOptions ...EventOptions) {
	if !a.InitializedAndEnabled() {
		return
	}

	var options EventOptions
	if len(eventOptions) > 0 {
		options = eventOptions[0]
	}

	a.setUserID(userID, &options)

	baseEvent := event.ToAmplitudeEvent()
	baseEvent.EventOptions = options

	a.Client.Track(baseEvent)
}

// Identify identifies a user and set user properties.
func (a *Ampli) Identify(userID string, eventOptions ...EventOptions) {
	identify := newBaseEvent(IdentifyEventType, nil)
	a.Track(userID, identify, eventOptions...)
}

// Flush flushes events waiting in buffer.
func (a *Ampli) Flush() {
	if !a.InitializedAndEnabled() {
		return
	}

	a.Client.Flush()
}

// Shutdown disables and shutdowns Ampli Instance.
func (a *Ampli) Shutdown() {
	if !a.InitializedAndEnabled() {
		return
	}

	a.mutex.Lock()
	a.Disabled = true
	a.mutex.Unlock()

	a.Client.Shutdown()
}

func (a *Ampli) CreateApplication(userID string, event CreateApplicationEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) CreateOrganisation(userID string, event CreateOrganisationEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) DeleteApplication(userID string, event DeleteApplicationEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) DeleteOrganisation(userID string, event DeleteOrganisationEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) ListApplications(userID string, event ListApplicationsEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) ListOrgs(userID string, eventOptions ...EventOptions) {
	event := ListOrgs.Builder().Build()
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) UpdateApplication(userID string, event UpdateApplicationEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}
