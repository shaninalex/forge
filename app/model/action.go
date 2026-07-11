package model

// Action - a general action that can be called during test.
type Action interface{}

// ActionType is the discriminator value stored in the "type" field.
type ActionType string

const (
	ActionHTTP            ActionType = "http"
	ActionGraphQL         ActionType = "graphql"
	ActionWebhookListener ActionType = "webhook_listener"
	ActionMcpCall         ActionType = "mcp_call"
)

// HttpAction performs a single HTTP request.
type HttpAction struct {
	Type        ActionType        `yaml:"type"`
	Method      HttpMethod        `yaml:"method"`
	Url         string            `yaml:"url"`
	Query       map[string]string `yaml:"query,omitempty"`
	Headers     map[string]string `yaml:"headers,omitempty"`
	Body        string            `yaml:"body,omitempty"`
	ContentType string            `yaml:"content_type,omitempty"`
}

// GraphQLAction performs a single GraphQL query.
type GraphQLAction struct {
	Type      ActionType        `yaml:"type"`
	Url       string            `yaml:"url"`
	Query     string            `yaml:"query"`
	Variables any               `yaml:"variables,omitempty"`
	Headers   map[string]string `yaml:"headers,omitempty"`
}

// WebhookListenerAction starts a listener that waits for an inbound request.
type WebhookListenerAction struct {
	Type           ActionType        `yaml:"type"`
	Port           uint16            `yaml:"port"`
	TimeoutSeconds uint64            `yaml:"timeout_seconds"`
	Headers        map[string]string `yaml:"headers,omitempty"`
}

// McpCallAction issues a prompt against an MCP server.
type McpCallAction struct {
	Type    ActionType        `yaml:"type"`
	Server  string            `yaml:"server"`
	Prompt  string            `yaml:"prompt"`
	Headers map[string]string `yaml:"headers,omitempty"`
}
