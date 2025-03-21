// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CompositionState string

// Enum values for CompositionState
const (
	CompositionStateStarting CompositionState = "STARTING"
	CompositionStateActive   CompositionState = "ACTIVE"
	CompositionStateStopping CompositionState = "STOPPING"
	CompositionStateFailed   CompositionState = "FAILED"
	CompositionStateStopped  CompositionState = "STOPPED"
)

// Values returns all known values for CompositionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CompositionState) Values() []CompositionState {
	return []CompositionState{
		"STARTING",
		"ACTIVE",
		"STOPPING",
		"FAILED",
		"STOPPED",
	}
}

type DestinationState string

// Enum values for DestinationState
const (
	DestinationStateStarting     DestinationState = "STARTING"
	DestinationStateActive       DestinationState = "ACTIVE"
	DestinationStateStopping     DestinationState = "STOPPING"
	DestinationStateReconnecting DestinationState = "RECONNECTING"
	DestinationStateFailed       DestinationState = "FAILED"
	DestinationStateStopped      DestinationState = "STOPPED"
)

// Values returns all known values for DestinationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DestinationState) Values() []DestinationState {
	return []DestinationState{
		"STARTING",
		"ACTIVE",
		"STOPPING",
		"RECONNECTING",
		"FAILED",
		"STOPPED",
	}
}

type EventErrorCode string

// Enum values for EventErrorCode
const (
	EventErrorCodeInsufficientCapabilities EventErrorCode = "INSUFFICIENT_CAPABILITIES"
	EventErrorCodeQuotaExceeded            EventErrorCode = "QUOTA_EXCEEDED"
	EventErrorCodePublisherNotFound        EventErrorCode = "PUBLISHER_NOT_FOUND"
)

// Values returns all known values for EventErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventErrorCode) Values() []EventErrorCode {
	return []EventErrorCode{
		"INSUFFICIENT_CAPABILITIES",
		"QUOTA_EXCEEDED",
		"PUBLISHER_NOT_FOUND",
	}
}

type EventName string

// Enum values for EventName
const (
	EventNameJoined           EventName = "JOINED"
	EventNameLeft             EventName = "LEFT"
	EventNamePublishStarted   EventName = "PUBLISH_STARTED"
	EventNamePublishStopped   EventName = "PUBLISH_STOPPED"
	EventNameSubscribeStarted EventName = "SUBSCRIBE_STARTED"
	EventNameSubscribeStopped EventName = "SUBSCRIBE_STOPPED"
	EventNamePublishError     EventName = "PUBLISH_ERROR"
	EventNameSubscribeError   EventName = "SUBSCRIBE_ERROR"
	EventNameJoinError        EventName = "JOIN_ERROR"
)

// Values returns all known values for EventName. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventName) Values() []EventName {
	return []EventName{
		"JOINED",
		"LEFT",
		"PUBLISH_STARTED",
		"PUBLISH_STOPPED",
		"SUBSCRIBE_STARTED",
		"SUBSCRIBE_STOPPED",
		"PUBLISH_ERROR",
		"SUBSCRIBE_ERROR",
		"JOIN_ERROR",
	}
}

type ParticipantState string

// Enum values for ParticipantState
const (
	ParticipantStateConnected    ParticipantState = "CONNECTED"
	ParticipantStateDisconnected ParticipantState = "DISCONNECTED"
)

// Values returns all known values for ParticipantState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ParticipantState) Values() []ParticipantState {
	return []ParticipantState{
		"CONNECTED",
		"DISCONNECTED",
	}
}

type ParticipantTokenCapability string

// Enum values for ParticipantTokenCapability
const (
	ParticipantTokenCapabilityPublish   ParticipantTokenCapability = "PUBLISH"
	ParticipantTokenCapabilitySubscribe ParticipantTokenCapability = "SUBSCRIBE"
)

// Values returns all known values for ParticipantTokenCapability. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ParticipantTokenCapability) Values() []ParticipantTokenCapability {
	return []ParticipantTokenCapability{
		"PUBLISH",
		"SUBSCRIBE",
	}
}

type PipBehavior string

// Enum values for PipBehavior
const (
	PipBehaviorStatic  PipBehavior = "STATIC"
	PipBehaviorDynamic PipBehavior = "DYNAMIC"
)

// Values returns all known values for PipBehavior. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PipBehavior) Values() []PipBehavior {
	return []PipBehavior{
		"STATIC",
		"DYNAMIC",
	}
}

type PipPosition string

// Enum values for PipPosition
const (
	PipPositionTopLeft     PipPosition = "TOP_LEFT"
	PipPositionTopRight    PipPosition = "TOP_RIGHT"
	PipPositionBottomLeft  PipPosition = "BOTTOM_LEFT"
	PipPositionBottomRight PipPosition = "BOTTOM_RIGHT"
)

// Values returns all known values for PipPosition. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PipPosition) Values() []PipPosition {
	return []PipPosition{
		"TOP_LEFT",
		"TOP_RIGHT",
		"BOTTOM_LEFT",
		"BOTTOM_RIGHT",
	}
}

type RecordingConfigurationFormat string

// Enum values for RecordingConfigurationFormat
const (
	RecordingConfigurationFormatHls RecordingConfigurationFormat = "HLS"
)

// Values returns all known values for RecordingConfigurationFormat. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecordingConfigurationFormat) Values() []RecordingConfigurationFormat {
	return []RecordingConfigurationFormat{
		"HLS",
	}
}

type VideoAspectRatio string

// Enum values for VideoAspectRatio
const (
	VideoAspectRatioAuto     VideoAspectRatio = "AUTO"
	VideoAspectRatioVideo    VideoAspectRatio = "VIDEO"
	VideoAspectRatioSquare   VideoAspectRatio = "SQUARE"
	VideoAspectRatioPortrait VideoAspectRatio = "PORTRAIT"
)

// Values returns all known values for VideoAspectRatio. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VideoAspectRatio) Values() []VideoAspectRatio {
	return []VideoAspectRatio{
		"AUTO",
		"VIDEO",
		"SQUARE",
		"PORTRAIT",
	}
}

type VideoFillMode string

// Enum values for VideoFillMode
const (
	VideoFillModeFill    VideoFillMode = "FILL"
	VideoFillModeCover   VideoFillMode = "COVER"
	VideoFillModeContain VideoFillMode = "CONTAIN"
)

// Values returns all known values for VideoFillMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VideoFillMode) Values() []VideoFillMode {
	return []VideoFillMode{
		"FILL",
		"COVER",
		"CONTAIN",
	}
}
