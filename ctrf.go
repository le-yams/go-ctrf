package ctrf

type Report struct {
	Results Results `json:"results"`
}

type Results struct {
	Tool        *Tool        `json:"tool"`
	Summary     *Summary     `json:"summary"`
	Tests       []*Test      `json:"tests"`
	Environment *Environment `json:"environment,omitempty"`
	Extra       interface{}  `json:"extra,omitempty"`
}

type Tool struct {
	Name    string      `json:"name"`
	Version string      `json:"version,omitempty"`
	Extra   interface{} `json:"extra,omitempty"`
}

type Summary struct {
	Tests   int         `json:"tests"`
	Passed  int         `json:"passed"`
	Failed  int         `json:"failed"`
	Pending int         `json:"pending"`
	Skipped int         `json:"skipped"`
	Other   int         `json:"other"`
	Suites  int         `json:"suites,omitempty"`
	Start   int64       `json:"start"`
	Stop    int64       `json:"stop"`
	Extra   interface{} `json:"extra,omitempty"`
}

type TestStatus string

const (
	TestPassed  TestStatus = "passed"
	TestFailed  TestStatus = "failed"
	TestSkipped TestStatus = "skipped"
	TestPending TestStatus = "pending"
	TestOther   TestStatus = "other"
)

type Test struct {
	Name       string        `json:"name"`
	Status     TestStatus    `json:"status"`
	Duration   int64         `json:"duration"`
	Start      int64         `json:"start,omitempty"`
	Stop       int64         `json:"stop,omitempty"`
	Suite      string        `json:"suite,omitempty"`
	Message    string        `json:"message,omitempty"`
	Trace      string        `json:"trace,omitempty"`
	RawStatus  string        `json:"rawStatus,omitempty"`
	Tags       []string      `json:"tags,omitempty"`
	Type       string        `json:"type,omitempty"`
	Filepath   string        `json:"filepath,omitempty"`
	Retry      int           `json:"retry,omitempty"`
	Flake      bool          `json:"flake,omitempty"`
	Browser    string        `json:"browser,omitempty"`
	Device     string        `json:"device,omitempty"`
	Screenshot string        `json:"screenshot,omitempty"`
	Parameters interface{}   `json:"parameters,omitempty"`
	Steps      []interface{} `json:"steps,omitempty"`
	Extra      interface{}   `json:"extra,omitempty"`
}

type Environment struct {
	AppName     string      `json:"appName,omitempty"`
	AppVersion  string      `json:"appVersion,omitempty"`
	OsPlatform  string      `json:"osPlatform,omitempty"`
	OsRelease   string      `json:"osRelease,omitempty"`
	OsVersion   string      `json:"osVersion,omitempty"`
	BuildName   string      `json:"buildName,omitempty"`
	BuildNumber string      `json:"buildNumber,omitempty"`
	Extra       interface{} `json:"extra,omitempty"`
}
