package models

import "time"

type ScheduleBody struct {
	Code               string                      `json:"code"`
	Owner              string                      `json:"owner"`
	Priority           *int                        `json:"priority,omitempty"`
	TimeoutMinutes     *int                        `json:"timeoutMinutes,omitempty"`
	RunVersion         *string                     `json:"runVersion,omitempty"` // 'v1' | 'v2' | 'v3'
	ExecPath           string                      `json:"execPath"`
	Body               *ScheduleBotBody            `json:"body,omitempty"`
	Time               *time.Time                  `json:"time,omitempty"`
	ExponentialBackoff *ScheduleExponentialBackoff `json:"exponentialBackoff,omitempty"`
	Hooks              []ScheduleHook              `json:"hooks"`
}

type Schedule struct {
	ID string `json:"_id"`
	ScheduleBody
}

type BotNext map[string]string

type BotStage struct {
	Name    string                 `json:"name"`
	Key     string                 `json:"key"`
	Version *string                `json:"version,omitempty"`
	Data    map[string]interface{} `json:"data"`
	Next    BotNext                `json:"next"`
}

type SchedulePostResponse struct {
	ID string `json:"_id"`
}

type ScheduleBotBody struct {
	Start         string                 `json:"start"`
	Version       int                    `json:"version"` // 1 | 2 | 3
	MaxIterations int                    `json:"maxIterations"`
	Stages        []BotStage             `json:"stages"`
	Data          map[string]interface{} `json:"data"`
}

type ScheduleHookEvent string

const (
	ScheduleHookEventOnError   ScheduleHookEvent = "on-error"
	ScheduleHookEventOnSuccess ScheduleHookEvent = "on-success"
	ScheduleHookEventOnFinish  ScheduleHookEvent = "on-finish"
	ScheduleHookEventOnStart   ScheduleHookEvent = "on-start"
)

type ScheduleHook struct {
	Event ScheduleHookEvent `json:"event"`
	API   string            `json:"api"`
	URL   string            `json:"url"`
}

type ScheduleExponentialBackoff struct {
	MaxRetries    *int `json:"maxRetries,omitempty"`
	PeriodMinutes *int `json:"periodMinutes,omitempty"`
	RetryCount    *int `json:"retryCount,omitempty"`
}

type SchedulePBBodyFactoryOptions struct {
	TimestampCode      *bool                       `json:"timestampCode,omitempty"`
	StartKey           *string                     `json:"startKey,omitempty"`
	TimeoutMinutes     *int                        `json:"timeoutMinutes,omitempty"`
	RunVersion         *string                     `json:"runVersion,omitempty"` // 'v1'|'v2'
	MaxIterations      *int                        `json:"maxIterations,omitempty"`
	Priority           *int                        `json:"priority,omitempty"`
	ExponentialBackoff *ScheduleExponentialBackoff `json:"exponentialBackoff,omitempty"`
	Hooks              []ScheduleHook              `json:"hooks,omitempty"`
}

type ScheduleLog struct {
	ID         string  `json:"_id"`
	ScheduleID string  `json:"scheduleId"`
	Output     *string `json:"output,omitempty"`
}
