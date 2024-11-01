package client

type Identity struct {
    Token string
    UserId string
}

type Client struct {
    engineURL                   string
    basePath                    string
    identity                    Identity
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

func NewClient(engineURL string, identity Identity) *Client {
    basePath := "/atlas_engine/api/v1"
    
	return &Client{
		engineURL:    engineURL,
        basePath: basePath,
		ProcessModels: NewProcessModels(engineURL, basePath, identity),
	}
}

func (c *Client) GetEngineURL() string {
    return c.engineURL
}

func (c *Client) GetIdentity() Identity {
    return c.identity
}

