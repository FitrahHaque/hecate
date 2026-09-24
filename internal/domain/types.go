package domain

import "time"

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)

// Message is one entry in a conversation transcript
type Message struct {
	ID        string
	SessionID string
	Role      Role
	Content   string
	Model     string // provider/model that produced it; empty for user messages
	Sequence  int    // order within the session
	CreatedAt time.Time
}

// ChatRequest is the normalized request every provider receives
type ChatRequest struct {
	Mode        string    // alias, e.g., "local" or "remote"
	Messages    []Message // effective conversation history for the request
	Stream      bool      // whether the output should be streamed to the user
	Temperature *float64  // determines predictibility; Pointer distinguishes between zero-value and default-value cases
	MaxToken    *int      // maximum number of tokens in the output; Pointer distinguishes between zero-value and default-value cases
}

// ChatResponse is the normalized completion every provider returns
type ChatResponse struct {
	ID           string
	Model        string
	Content      string
	FinishReason string // informs the kind of reason for stopping the response
	Usage        Usage
}

// Usage is the per-request token and cost accounting; EstimatedCost (USD) is computed by the gateway, not returned by providers
type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
	EstimatedCost    float64
}

// Session is a persistent conversation; the transcript lives in Messages.
type Session struct {
	ID            string
	Title         string
	Summary       string // short summary of deleted (old) messages, re-injected into every future request's context
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
