package client

type Identity struct {
    Token string
    UserId string
}

type Client struct {
    engineURL                   string
    basePath                    string
    identity                    *Identity
    // anonymousSessionClient      AnonymousSessionClient
    // applicationInfoClient       ApplicationInfoClient
    // correlationClient           CorrelationClient
    // cronjobClient               CronjobClient
    // dataObjectInstanceClient    DataObjectInstanceClient
    // untypedTaskClient           UntypedTaskClient
    // eventClient                 EventClient
    // externalTaskClient          ExternalTaskClient
    // flowNodeInstanceClient      FlowNodeInstanceClient
    // manualTaskClient            ManualTaskClient
    // notificationClient          NotificationClient
    // processDefinitionClient     ProcessDefinitionClient
    // processInstanceClient       ProcessInstanceClient
    ProcessModels               *ProcessModelClient
    // userTaskClient              UserTaskClient
}

func NewClient(engineURL string, identity ...*Identity) *Client {
    basePath := "/atlas_engine/api/v1"

    var identityToUse *Identity // Declare a pointer for identity
    if len(identity) > 0 {
        identityToUse = identity[0] // Use the first element if provided
    } else {
        // Create a dummy identity if none is provided
        identityToUse = &Identity{Token: "ZHVtbXlfdG9rZW4=", UserId: "dummy_token"}
    }

    return &Client{
        engineURL:     engineURL,
        basePath:      basePath,
        identity:      identityToUse,
        ProcessModels: NewProcessModels(engineURL, basePath, identityToUse),
    }
}

func (c *Client) GetEngineURL() string {
    return c.engineURL
}

func (c *Client) GetIdentity() *Identity {
    return c.identity
}

