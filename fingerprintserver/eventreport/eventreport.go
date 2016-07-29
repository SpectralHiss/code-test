package eventreport

type CopyEventReport struct {
	Common
	Pasted bool
	FormID string
}

type Common struct {
	EventType  string
	WebsiteURL string
	SessionID  string
}

type Data struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int             // Seconds
}

type Dimension struct {
	Width  string
	Height string
}

type DelayEventReport struct {
	Common
	Time int
}

type Dimensions struct {
	Width  string
	Height string
}

type ResizeEventReport struct {
	Common
	BeforeDimensions Dimensions
	AfterDimensions  Dimensions
}
