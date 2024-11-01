package client

type ProcessModelResponse struct {
    ProcessModels []ProcessModel `json:"processModels"`
    TotalCount    int             `json:"totalCount"`
}

type ProcessModel struct {
    ProcessDefinitionID string                  `json:"processDefinitionId"`
    ProcessModelID      string                  `json:"processModelId"`
    ProcessModelName    string                  `json:"processModelName"`
    CustomProperties     map[string]interface{}  `json:"customProperties"`
    IsExecutable        bool                    `json:"isExecutable"`
    LaneSet             LaneSet                 `json:"laneSet"`
    StartEvents         []FlowNode              `json:"startEvents"`
    EndEvents           []FlowNode              `json:"endEvents"`
    FlowNodes           []FlowNode              `json:"flowNodes"`
}

type LaneSet struct {
    Lanes []Lane `json:"lanes"`
}

type Lane struct {
    ID                string              `json:"id"`
    ExtensionElements ExtensionElements   `json:"extensionElements"`
    Name              string              `json:"name"`
    FlowNodeReferences []string           `json:"flowNodeReferences"`
}

type ExtensionElements struct {
    CamundaExtensionProperties []interface{} `json:"camundaExtensionProperties"`
}

type FlowNode struct {
    ID                string              `json:"id"`
    CustomProperties   map[string]interface{} `json:"customProperties"`
    ProcessModelID    string              `json:"processModelId"`
    ProcessModelName  string              `json:"processModelName"`
    Name              string              `json:"name,omitempty"`       // Omit if empty
    FlowNodeType      string              `json:"flowNodeType"`
    TimerType         *string             `json:"timerType,omitempty"`   // Omit if nil
    FormFields        map[string]FormField `json:"formFields,omitempty"`  // Omit if empty
}

type FormField struct {
    ID          string `json:"id"`
    Label       string `json:"label"`
    Type        string `json:"type"`
    DefaultValue *string `json:"defaultValue,omitempty"` // Omit if nil
    Index       int    `json:"index"`
}
