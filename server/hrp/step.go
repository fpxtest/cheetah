package hrp

import (
	"github.com/test-instructor/cheetah/server/hrp/pkg/gidevice"
	"github.com/test-instructor/cheetah/server/hrp/pkg/uixt"
)

type StepType string

const (
	stepTypeRequest     StepType = "request"
	stepTypeAPI         StepType = "api"
	stepTypeTestCase    StepType = "testcase"
	stepTypeTransaction StepType = "transaction"
	stepTypeRendezvous  StepType = "rendezvous"
	stepTypeThinkTime   StepType = "thinktime"
	stepTypeWebSocket   StepType = "websocket"
	stepTypeAndroid     StepType = "android"
	stepTypeIOS         StepType = "ios"
)

var (
	WithIdentifier          = uixt.WithIdentifier
	WithMaxRetryTimes       = uixt.WithMaxRetryTimes
	WithWaitTime            = uixt.WithWaitTime
	WithIndex               = uixt.WithIndex
	WithTimeout             = uixt.WithTimeout
	WithIgnoreNotFoundError = uixt.WithIgnoreNotFoundError
	WithText                = uixt.WithText
	WithID                  = uixt.WithID
	WithDescription         = uixt.WithDescription
	WithDirection           = uixt.WithDirection
	WithCustomDirection     = uixt.WithCustomDirection
	WithScope               = uixt.WithScope
	WithOffset              = uixt.WithOffset
)

var (
	WithPerfSystemCPU         = gidevice.WithPerfSystemCPU
	WithPerfSystemMem         = gidevice.WithPerfSystemMem
	WithPerfSystemDisk        = gidevice.WithPerfSystemDisk
	WithPerfSystemNetwork     = gidevice.WithPerfSystemNetwork
	WithPerfGPU               = gidevice.WithPerfGPU
	WithPerfFPS               = gidevice.WithPerfFPS
	WithPerfNetwork           = gidevice.WithPerfNetwork
	WithPerfBundleID          = gidevice.WithPerfBundleID
	WithPerfPID               = gidevice.WithPerfPID
	WithPerfOutputInterval    = gidevice.WithPerfOutputInterval
	WithPerfProcessAttributes = gidevice.WithPerfProcessAttributes
	WithPerfSystemAttributes  = gidevice.WithPerfSystemAttributes
)

type StepResult struct {
	ID               uint                   `json:"ID"`
	ParntID          uint                   `json:"parntID"`
	Name             string                 `json:"name" yaml:"name"`                                   // step name
	StepType         StepType               `json:"step_type" yaml:"step_type"`                         // step type, testcase/request/transaction/rendezvous
	Success          bool                   `json:"success" yaml:"success"`                             // step execution result
	Elapsed          int64                  `json:"elapsed_ms" yaml:"elapsed_ms"`                       // step execution time in millisecond(ms)
	HttpStat         map[string]int64       `json:"httpstat,omitempty" yaml:"httpstat,omitempty"`       // httpstat in millisecond(ms)
	Data             interface{}            `json:"data,omitempty" yaml:"data,omitempty"`               // session data or slice of step data
	ContentSize      int64                  `json:"content_size" yaml:"content_size"`                   // response body length
	ExportVars       map[string]interface{} `json:"export_vars,omitempty" yaml:"export_vars,omitempty"` // extract variables
	Attachment       string                 `json:"attachment,omitempty" yaml:"attachment,omitempty"`   // step error information
	ValidatorsNumber uint                   `json:"validate_number,omitempty"`

	Attachments interface{} `json:"attachments,omitempty" yaml:"attachments,omitempty"` // store extra step information, such as error message or screenshots
}

// TStep represents teststep data structure.
// Each step maybe three different types: make one request or reference another api/testcase.
type TStep struct {
	Name             string                 `json:"name" yaml:"name"` // required
	Request          *Request               `json:"request,omitempty" yaml:"request,omitempty"`
	API              interface{}            `json:"api,omitempty" yaml:"api,omitempty"`           // *APIPath or *API
	TestCase         interface{}            `json:"testcase,omitempty" yaml:"testcase,omitempty"` // *TestCasePath or *TestCase
	Transaction      *Transaction           `json:"transaction,omitempty" yaml:"transaction,omitempty"`
	Rendezvous       *Rendezvous            `json:"rendezvous,omitempty" yaml:"rendezvous,omitempty"`
	ThinkTime        *ThinkTime             `json:"think_time,omitempty" yaml:"think_time,omitempty"`
	WebSocket        *WebSocketAction       `json:"websocket,omitempty" yaml:"websocket,omitempty"`
	Android          *MobileStep            `json:"android,omitempty" yaml:"android,omitempty"`
	IOS              *MobileStep            `json:"ios,omitempty" yaml:"ios,omitempty"`
	Variables        map[string]interface{} `json:"variables,omitempty" yaml:"variables,omitempty"`
	SetupHooks       []string               `json:"setup_hooks,omitempty" yaml:"setup_hooks,omitempty"`
	TeardownHooks    []string               `json:"teardown_hooks,omitempty" yaml:"teardown_hooks,omitempty"`
	Extract          map[string]string      `json:"extract,omitempty" yaml:"extract,omitempty"`
	Validators       []interface{}          `json:"validate,omitempty" yaml:"validate,omitempty"`
	Export           []string               `json:"export,omitempty" yaml:"export,omitempty"`
	ValidatorsNumber uint                   `json:"validate_number,omitempty"`
	ID               uint                   `json:"ID"`
	ParntID          uint                   `json:"parntID"`
	ExportHeader     []string               `json:"export_header"`
	ExportParameter  []string               `json:"export_parameter"`
}

// IStep represents interface for all types for teststeps, includes:
// StepRequest, StepRequestWithOptionalArgs, StepRequestValidation, StepRequestExtraction,
// StepTestCaseWithOptionalArgs,
// StepTransaction, StepRendezvous, StepWebSocket.
type IStep interface {
	Name() string
	Type() StepType
	Struct() *TStep
	Run(*SessionRunner) (*StepResult, error)
}
