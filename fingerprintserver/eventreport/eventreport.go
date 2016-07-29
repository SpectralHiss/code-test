package eventreport

type CopyEventReport struct {
	EventType  string `json:"eventType"`
	WebsiteURL string `json:"websiteUrl"`
	SessionID  string `json:"sessionId"`
	Pasted     bool
	FormID     string `json:"formId"`
}

type Data struct {
	WebsiteUrl         string          `json:"website_url"`
	SessionId          string          `json:"session_id"`
	ResizeFrom         Dimension       `json:"resize_from"`
	ResizeTo           Dimension       `json:"resize_to"`
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int             `json:"form_completion_time"`
}

type DelayEventReport struct {
	EventType  string `json:"eventType"`
	WebsiteURL string `json:"websiteUrl"`
	SessionID  string `json:"sessionId"`
	Time       int
}

type Dimension struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

type ResizeEventReport struct {
	EventType  string    `json:"eventType"`
	WebsiteURL string    `json:"websiteUrl"`
	SessionID  string    `json:"sessionId"`
	ResizeFrom Dimension `json:"resizeFrom"`
	ResizeTo   Dimension `json:"resizeTo"`
}
